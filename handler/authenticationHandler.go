package handler

import (
	"encoding/json"
	"github.com/edwardsuwirya/mdw/model"
	"github.com/edwardsuwirya/mdw/utils"
	"log"
	"net/http"
)

type authenticationHandler struct {
}

func NewAuthenticationHandler() IHandler {
	return &authenticationHandler{}
}

func (h *authenticationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var sysUser model.SysUser
	json.NewDecoder(r.Body).Decode(&sysUser)

	userName := sysUser.UserName
	userPassword := sysUser.UserPassword

	if userName == "edo" && userPassword == "ini_mestinya_di_hash" {
		w.WriteHeader(http.StatusOK)
		token, err := utils.JwtEncoder(userName, "Rahasia dong")
		if err != nil {
			log.Print(err)
			http.Error(w, "Failed token generation", http.StatusInternalServerError)
		}
		w.Write([]byte(token))
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}

}
