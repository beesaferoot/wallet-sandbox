package wallet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wallet-sandbox/db"
)

type AccountInfo struct {
	AccountId int `json:"account_id"`
	Amount int `json:"amount"`
}

type TransferInfo struct {
	FromAccountId int `json:"from_account_id"`
	ToAccountId int `json:"to_account_id"`
	Amount int `json:"amount"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func FundHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	accInfo := AccountInfo{}
	err = json.Unmarshal(body, &accInfo)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}

	// validate amount value
	if accInfo.Amount < 0 {
		res, _ := json.Marshal(InvalidAmount)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}

	// fund account wallet 
	err = db.DAO.FundWallet(globConfig, accInfo.AccountId, accInfo.Amount)
	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}
	
	res, _ := json.Marshal(SuccessResponse{Message: 
		fmt.Sprintf("Wallet with account = %d has been successfully funded.", accInfo.AccountId)})
	w.Write(res)
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	transferInfo := TransferInfo{}
	err = json.Unmarshal(body, &transferInfo)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}

	err = db.DAO.Transfer(globConfig, 
		transferInfo.FromAccountId, 
		transferInfo.ToAccountId, 
		transferInfo.Amount,
	)

	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(SuccessResponse{Message: "Transfer successful"})
	w.Write(res)
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	accInfo := AccountInfo{}
	err = json.Unmarshal(body, &accInfo)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}

	// validate amount value
	if accInfo.Amount < 0 {
		res, _ := json.Marshal(InvalidAmount)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}

	wallet:= &db.Wallet{AccountId: accInfo.AccountId}
	err = wallet.Get(globConfig)
	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	// check that wallet balance is sufficient 
	if accInfo.Amount > wallet.Balance {
		res, _ := json.Marshal(InsufficientBalance)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}

	// fund account wallet 
	err = db.DAO.WalletWithdraw(globConfig, accInfo.AccountId, accInfo.Amount)
	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}
	
	res, _ := json.Marshal(SuccessResponse{Message: "Successful withdrawal"})
	w.Write(res)
}