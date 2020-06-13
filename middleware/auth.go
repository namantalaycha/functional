package middleware

import (
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("name")
		if val == "" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}
