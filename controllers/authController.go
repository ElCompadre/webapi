package controllers

import (
	"encoding/json"
	"github.com/elcompadre/webapi/models"
	"github.com/elcompadre/webapi/utils"
	"net/http"
)

var CreateAccount = func(responseWriter http.ResponseWriter, request *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(responseWriter, utils.Message(false, "Invalid request"))
	}

	resp := account.Create()
	utils.Respond(responseWriter, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
