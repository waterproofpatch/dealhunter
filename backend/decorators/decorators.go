package decorators

import (
	"context"
	"fmt"
	"net/http"
)

func TokenDecorator(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", token)
		h(w, r.WithContext(ctx))
	}
}

func LogDecorator(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received a %s request at %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		h(w, r)
	}
}
