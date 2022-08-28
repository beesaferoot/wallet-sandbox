package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)


var (
	host = os.Getenv("DATABASE_HOST")
	port = os.Getenv("DATABASE_PORT")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB")
)

type DataAccessObject struct {
	conn *sql.DB
}

var DAO = &DataAccessObject{}

func init(){
	var err error 
	sourceStr :=  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	host, port, user, password, dbname)
	DAO.conn, err = sql.Open("postgres", sourceStr)
	if err != nil {
		log.Fatal(err)
	}
}

// Account db operations
func (dao *DataAccessObject) GetAccountList(cxt context.Context) (accounts []Account, err error) {
	return Accounts(cxt)
}

func (dao *DataAccessObject) CreateAccount(cxt context.Context, accName string) (acc Account, err error) {
	acc.Name = accName
	return acc, acc.Create(cxt)
}

// Wallet db operations
func (dao *DataAccessObject) CreateWallet(cxt context.Context, account_id int) (wallet Wallet, err error) {
	wallet.AccountId = account_id
	return wallet, wallet.Create(cxt)
}

func (dao *DataAccessObject) GetWallet(ctx context.Context, account_id int) (wallet Wallet, err error) {
	wallet.AccountId = account_id
	return wallet, wallet.Get(ctx)
}

func (dao *DataAccessObject) FundWallet(ctx context.Context, account_id, amount int) (error ) {
	w := Wallet{}
	trnx, err := dao.createTransaction(ctx, "Credit", fmt.Sprint(account_id), 
										fmt.Sprint(account_id), fmt.Sprint(amount))
	if err != nil {
		return err
	}
	err = w.Fund(ctx, account_id, amount)
	if err != nil {
		_ = trnx.UpdateFailure(ctx)
		return err
	}

	_ = trnx.UpdateSuccess(ctx)
	return err
}

func (dao *DataAccessObject) WalletWithdraw(ctx context.Context, account_id, amount int) (error) {
	w := Wallet{}
	trnx, err := dao.createTransaction(ctx, "Debit", fmt.Sprint(account_id), 
										fmt.Sprint(account_id), fmt.Sprint(amount))
	if err != nil {
		return err
	}
	err = w.Withdraw(ctx, account_id, amount)
	if err != nil {
		_ = trnx.UpdateFailure(ctx)
		return err
	}

	_ = trnx.UpdateSuccess(ctx)
	return err
}

func (dao *DataAccessObject) Transfer(ctx context.Context, fromAccId, toAccId, amount int) (err error) {
	trnx, err := dao.createTransaction(ctx, "Transfer", fmt.Sprint(fromAccId), 
	fmt.Sprint(toAccId), fmt.Sprint(amount))
	if err != nil {
		return
	}

	w := Wallet{}
	err = w.Transfer(ctx, fromAccId, toAccId, amount)
	if err != nil {
		_ = trnx.UpdateFailure(ctx)
		return err
	}

	_ = trnx.UpdateSuccess(ctx)
	return 
}

// Wallet transanction db operations
func (dao *DataAccessObject) createTransaction(
	cxt context.Context, 
	Type,
	fromAccountId, 
	toAccountId, amount string) (trxn Transaction, err error ){
	trxn.FromAccountId = fromAccountId
	trxn.ToAccountId = toAccountId
	trxn.Type = Type
	trxn.Amount = amount
	return trxn, trxn.Create(cxt)
}

func (dao *DataAccessObject) GetTransanctionListByDate(cxt context.Context, date time.Time) (trnxs []Transaction, err error) {
	return TransactionsByDate(cxt, date)
}