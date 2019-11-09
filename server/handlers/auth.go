package handlers

import (
	"encoding/base64"
	"encoding/json"
	"leva-api/models"
	"leva-api/models/entities"
	"leva-api/utils"
	"leva-api/utils/jwt"
	"log"
	"net/http"
	"strings"
)

//Auth check the credentials and authorizes an user
func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "" {
		token := (strings.Split(r.Header.Get("Authorization"), "Basic"))[1]

		credentials, err := base64.StdEncoding.DecodeString(strings.TrimSpace(token))
		if err != nil {
			log.Println(err.Error())
			return
		}

		email := strings.Split(string(credentials), ":")[0]
		password := strings.Split(string(credentials), ":")[1]

		results, err := entities.User.Where("email", "=", email).And().Where("senha", "=", password).Limit(1).Get("id", "nome", "email", "rg", "cpf", "genero", "telefone", "status", "foto", "tipo")
		if err != nil {
			log.Println(err.Error())
			utils.HTTPResponse(w, http.StatusInternalServerError, "Any user was found with this credentials. Check them and try again.", false)
			return
		}
		defer results.Close()

		var user models.User
		for results.Next() {
			results.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.RG,
				&user.CPF,
				&user.Gender,
				&user.Phone,
				&user.Status,
				&user.Picture,
				&user.Type,
			)
		}

		payload, _ := json.Marshal(user)
		key := []byte(password)

		token, err = jwt.GenerateToken(string(payload), key)

		utils.HTTPResponse(w, http.StatusOK, utils.BuildString(`{"token": "`, token, `", "user": `, string(payload), `}`), true)
	} else {
		utils.HTTPResponse(w, http.StatusUnauthorized, "Invalid Authorization", false)
		return
	}
}
