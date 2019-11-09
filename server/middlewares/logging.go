package middlewares

import (
	"leva-api/utils"
	"leva-api/utils/jwt"
	"net/http"
)

//Logging is responsible for controlling the access and the user token validation
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		key := []byte(r.Header.Get("Key"))

		if r.RequestURI == "/auth" || r.RequestURI == "/register" || r.RequestURI == "/register/verification/code" &&
			r.Method == "POST" || r.Method == "PATCH" {
			next.ServeHTTP(w, r)
		} else {
			if token != "" {
				_, validtoken := jwt.DecodeToken(token, key)
				if validtoken != false {
					next.ServeHTTP(w, r)
				} else {
					utils.HTTPResponse(w, http.StatusUnauthorized, "Invalid token!", false)
					return
				}
			} else {
				utils.HTTPResponse(w, http.StatusUnauthorized, "Undefined token!", false)
				return
			}
		}
	})
}
