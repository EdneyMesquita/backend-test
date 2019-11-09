package handlers

import (
	"leva-api/utils"
	"math/rand"
	"net/http"
)

//GetCode generates a verification code for the phone number verification
func GetCode(w http.ResponseWriter, r *http.Request) {
	code := rand.Intn(9999)

	for code <= 999 {
		code = rand.Intn(9999)
	}

	utils.HTTPResponse(w, http.StatusOK, utils.BuildString(`{"code": `, code, "}"), true)
}
