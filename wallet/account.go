package wallet

import (
	"context"
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"wallet-sandbox/db"
)

var globConfig = context.Background()

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	acc := db.Account{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	err = json.Unmarshal(body, &acc)
	if err != nil {
		res, _ := json.Marshal(BadRequest.Error(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return 
	}
	// create account 
	acc, err = db.DAO.CreateAccount(globConfig, acc.Name)

	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}
	// create wallet with account id 
	wallet, err := db.DAO.CreateWallet(globConfig, acc.Id)

	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(wallet)
	w.Write(res)
}

func ListAccountHandler(w http.ResponseWriter, r *http.Request) {
	accList, err := db.DAO.GetAccountList(globConfig)
	if err != nil {
		res, _ := json.Marshal(ServerError.Error(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(accList)
	w.Write(res)
}
