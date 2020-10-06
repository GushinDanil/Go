package middlewares

import (
	"Rest/app/auth"
	"errors"
	"net/http"

	"Rest/app/utils"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			utils.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized5"))

			return
		}
		next(w, r)
	}
}
