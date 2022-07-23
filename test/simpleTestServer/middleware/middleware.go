package middleware

import (
	"crypto/subtle"
	"net/http"
)

const (
	username string = "username"
	password string = "password"
	realm    string = "TestRealm"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unauthorised.\n"))

			if err != nil {
				http.Error(w, "Unexpected error", http.StatusInternalServerError)
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}

func MethodValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err := w.Write([]byte("Method not supported"))

			if err != nil {
				http.Error(w, "Unexpected error", http.StatusInternalServerError)
			}

			return
		}
		next.ServeHTTP(w, r)
	})
}
