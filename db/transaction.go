package db

import (
	"time"
	"context"
)

type Transaction struct {
	Id int  				`json:"id"`
	Type string  			`json:"type"`
	FromAccountId string  `json:"from_account_id"`
	ToAccountId string    `json:"to_account_id"`
	Amount string 		   `json:"amount"`
	Status string 		  `json:"status"`
	ProcessedAt time.Time  `json:"processed_at"`
	CreatedAt time.Time    `json:"created_at"`
}

func (trxn *Transaction) CreatedAtDate() string {
	return trxn.CreatedAt.Format("2006-01-02 15:04:05")
}

func (trxn *Transaction) Create(cxt context.Context) (err error) {
	stmt := "INSERT INTO wallet_sandbox.transactions (type, from_account_id, to_account_id, amount) VALUES($1, $2, $3, $4) returning id, created_at"
	err = DAO.conn.QueryRowContext(cxt, stmt, 
		 trxn.Type, trxn.FromAccountId,
		 trxn.ToAccountId, trxn.Amount,
		 ).Scan(&trxn.Id, &trxn.CreatedAt)
	return err
}

func (trxn *Transaction) UpdateSuccess(cxt context.Context) (err error) {
	trxn.Status = "SUCCESS"
	stmt := "UPDATE wallet_sandbox.transactions SET status = $2, processed_at = $3 WHERE id = $1"
	err = DAO.conn.QueryRowContext(cxt, stmt, trxn.Id, trxn.Status, time.Now()).Err()
	return err
}

func (trxn *Transaction) UpdateFailure(cxt context.Context) (err error) {
	trxn.Status = "FAILURE"
	stmt := "UPDATE wallet_sandbox.transactions SET status = $2 WHERE id = $1"
	err = DAO.conn.QueryRowContext(cxt, stmt, trxn.Id, trxn.Status).Err()
	return err
}

func TransactionsByDate(cxt context.Context, date time.Time) (trnxs []Transaction, err error) {
	trnxs = []Transaction{}
	stmt := "SELECT id, type, from_account_id, to_account_id, amount, status, processed_at, created_at FROM wallet_sandbox.transactions WHERE  created_at <= TO_TIMESTAMP($1, 'YYYY-MM-DD HH24:MI:SS') ORDER BY created_at DESC"
	rows, err := DAO.conn.QueryContext(cxt, stmt, date)

	if err != nil {
		return trnxs, err
	}

	for rows.Next() {
		trnx := Transaction{}
		if err = rows.Scan(&trnx.Id, &trnx.Type,
			 &trnx.FromAccountId, 
			 &trnx.ToAccountId, 
			 &trnx.Amount, 
			 &trnx.Status, &trnx.ProcessedAt, &trnx.CreatedAt); err != nil {
				return trnxs, err
			 }
		trnxs = append(trnxs, trnx)
	}

	return trnxs, nil
}