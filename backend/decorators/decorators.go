package decorators

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"deals/environment"
	"deals/logging"

	"github.com/dgrijalva/jwt-go"
)

func TokenDecorator(allowUnauthenticated bool, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if allowUnauthenticated {
			h(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			accessToken := bearerToken[1]
			token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(environment.GetEnvironment().JWT_SIGNING_KEY), nil
			})

			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// JWT is expired
					http.Error(w, "Token expired", 419)
					return
				} else {
					logging.GetLogger().Error().Msg(err.Error())
					http.Error(w, "Invalid token", http.StatusUnauthorized)
					return
				}
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "token", claims)
				h(w, r.WithContext(ctx))
			} else {
				logging.GetLogger().Error().Msg(err.Error())
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
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
