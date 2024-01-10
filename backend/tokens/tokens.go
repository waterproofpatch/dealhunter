package tokens

import (
	"strconv"
	"time"

	"deals/environment"
	"deals/logging"
	"deals/models"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(user models.User) string {
	// Create a new token object, specifying signing method and the claims
	expireMin, err := strconv.Atoi(environment.GetEnvironment().JWT_ACCESS_TOKEN_EXPIRE_MIN)
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		panic("Failed generating token (1)!")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * time.Duration(expireMin)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessTokenSigningKey := environment.GetEnvironment().JWT_SIGNING_KEY
	accessToken, err := token.SignedString([]byte(accessTokenSigningKey))
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		panic("Failed generating token (2)!")
	}
	return accessToken
}

func GenerateRefrehToken(user models.User) string {
	// Create a new token object, specifying signing method and the claims
	expireMin, err := strconv.Atoi(environment.GetEnvironment().JWT_REFRESH_TOKEN_EXPIRE_MIN)
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		panic("Failed generating token (1)!")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * time.Duration(expireMin)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	refreshTokenSigningKey := environment.GetEnvironment().JWT_SIGNING_KEY_REFRESH
	refreshToken, err := token.SignedString([]byte(refreshTokenSigningKey))
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		panic("Failed generating token (2)!")
	}
	return refreshToken
}
