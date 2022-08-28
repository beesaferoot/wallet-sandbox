package wallet

import "fmt"

type ErrorMessage struct {
	Message string `json:"error"`
}

func (e ErrorMessage) Error(extra_msg string) ErrorMessage {
	return ErrorMessage{Message: fmt.Sprintf("%s - %s", e.Message, extra_msg)}
}

var BadRequest ErrorMessage = ErrorMessage{Message: "Bad request."}
var ServerError ErrorMessage = ErrorMessage{Message: "Server error occurred."}
var InvalidAmount ErrorMessage = ErrorMessage{Message: "Invalid amount (negative value)."}
var InsufficientBalance ErrorMessage = ErrorMessage{Message: "Your Balance is insufficient."}