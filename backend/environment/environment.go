package environment

import (
	"errors"
	"os"
)

type Environment struct {
	JWT_SIGNING_TOKEN string
	DB_NAME           string
	DB_PASSWORD       string
	PORT              string
}

func Init() (*Environment, error) {
	jwtSigningToken := os.Getenv("JWT_SIGNING_TOKEN")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")

	if jwtSigningToken == "" || dbName == "" || dbPassword == "" || port == "" {
		return nil, errors.New("Environment variables JWT_SIGNING_TOKEN, DB_NAME, or DB_PASSWORD are not set")
	}

	env := &Environment{
		JWT_SIGNING_TOKEN: jwtSigningToken,
		DB_NAME:           dbName,
		DB_PASSWORD:       dbPassword,
		PORT:              port,
	}

	return env, nil
}
