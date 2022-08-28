package db

import (
	"time"
	"context"
)

type Account struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreateAt  time.Time `json:"created_at"`
}

func (acc *Account) Create(cxt context.Context) (err error) {
	stmt := "INSERT INTO wallet_sandbox.accounts (name) VALUES ($1) returning id, created_at"
	row := DAO.conn.QueryRowContext(cxt, stmt, acc.Name)

	err = row.Scan(&acc.Id, &acc.CreateAt)

	return err
}

func (acc *Account) CreatedAtDate() string {
	return acc.CreateAt.Format("2006-01-02 15:04:05")
}

func Accounts(cxt context.Context) (accounts []Account, err error) {
	accounts = []Account{}
	stmt := "SELECT id, name, created_at FROM wallet_sandbox.accounts ORDER BY created_at DESC"
	rows, err := DAO.conn.QueryContext(cxt, stmt)

	if err != nil {
		return accounts, err
	}

	for rows.Next() {
		acc := Account{}
		if err = rows.Scan(&acc.Id, &acc.Name, &acc.CreateAt); err != nil {
			return accounts, err
		}
		accounts = append(accounts, acc)
	}
	
	rows.Close()
	return accounts, nil
}

