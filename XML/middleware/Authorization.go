package middleware

import (
	"net/http"

	myauth "user/module/auth"
)

type KeyProduct struct{}

func ValidateToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		tokenString := myauth.ExtractToken(r)
		err := myauth.TokenIsValid(tokenString)
		if err != nil {
			http.Error(rw, "Error middleware validating token: "+err.Error()+"\n", http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusAccepted)
		rw.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	})
}
