package main

import (
	"net/http"
	"wallet-sandbox/utils"
	"wallet-sandbox/wallet"
)

func main() {
	mux := utils.NewHTTPMultiplexer()
	mux.POST("/account/create", wallet.CreateAccountHandler)
	mux.POST("/wallet/fund", wallet.FundHandler)
	mux.POST("/wallet/transfer", wallet.TransferHandler)
	mux.POST("/wallet/withdraw", wallet.WithdrawHandler)
	
	mux.GET("/account/list", wallet.ListAccountHandler)
	mux.GET("/wallet/transaction_history", wallet.TransactionHistoryHandler)
	
	server := &http.Server{
		Addr:     "0.0.0.0:8088",
		Handler: mux,
	}
	server.ListenAndServe()
}