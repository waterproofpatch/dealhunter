package decorators

import (
	"context"
	"net/http"

	"deals/logging"
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
		// fmt.Printf("Received a %s request at %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		logging.GetLogger().Info().
			Str("method", r.Method).
			Str("url", r.URL.Path).
			Str("remoteAddr", r.RemoteAddr).
			Msg("Received a request")

		// logging.GetLogger().Info().Msg("Hello from Zerolog logger")
		h(w, r)
	}
}
