package db

import (
	"time"
	"context"
	"database/sql"
)

type Wallet struct {
	Id int		`json:"id"`
	Balance int `json:"balance"`
	AccountId int `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`
}


func (w *Wallet) Create(ctx context.Context) (err error) {
	stmt := "INSERT INTO wallet_sandbox.wallets (account_id) VALUES ($1) returning id, created_at"
	err = DAO.conn.QueryRowContext(ctx, stmt, w.AccountId).Scan(&w.Id, &w.CreatedAt)
	return
}

func (w *Wallet) Get(ctx context.Context) (err error) {
	stmt := "SELECT id, balance, created_at FROM wallet_sandbox.wallets WHERE account_id = $1"
	err = DAO.conn.QueryRowContext(ctx, stmt, w.AccountId).Scan(&w.Id, &w.Balance, &w.CreatedAt)
	return
}

func (w *Wallet) Fund(ctx context.Context, accountId, amount int ) (err error) {
	stmt := "UPDATE wallet_sandbox.wallets SET balance = balance + $2 WHERE account_id = $1"
	err = DAO.conn.QueryRowContext(ctx, stmt, accountId, amount).Err()
	return err
}

func (w *Wallet) Withdraw(ctx context.Context, accountId, amount int) (err error) {
	stmt := "UPDATE wallet_sandbox.wallets SET balance = balance - $2 WHERE account_id = $1"
	err = DAO.conn.QueryRowContext(ctx, stmt, accountId, amount).Err()
	return
}

func (w *Wallet) Transfer(ctx context.Context, fromAccId, toAccId int, amount int) (err error) {
	trxOpts := &sql.TxOptions{}
	trx, err := DAO.conn.BeginTx(ctx, trxOpts)
	if err != nil {
		return
	}
	// update account wallet to transfer from
	_, err = trx.Exec("UPDATE wallet_sandbox.wallets SET balance = balance - $1 WHERE account_id = $2", amount, fromAccId)
	if err != nil {
		_ = trx.Rollback()
		return err 
	}

	// update account wallet to transfer to
	_, err = trx.Exec("UPDATE wallet_sandbox.wallets SET balance = balance + $1 WHERE account_id = $2", amount, toAccId)
	if err != nil {
		_ = trx.Rollback()
		return err 
	}
	
	if err = trx.Commit(); err != nil {
		_  = trx.Rollback()
		return err
	}
	return nil
}