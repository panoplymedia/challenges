package main

import (
	"net/http"
	"strings"
)

type authMiddleware struct {
	token string
}

func (amw *authMiddleware) SetToken(token string) error {
	amw.token = token
	return nil
}

func (amw *authMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		token = strings.TrimPrefix(token, "Bearer")
		token = strings.TrimSpace(token)
		w.Header().Set("Content-Type", "application/json")

		if token != amw.token {
			w.WriteHeader(http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
