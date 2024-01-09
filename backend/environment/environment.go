package environment

import (
	"errors"
	"os"

	"deals/logging"
)

type Environment struct {
	JWT_SIGNING_KEY             string
	JWT_ACCESS_TOKEN_EXPIRE_MIN string
	DB_NAME                     string
	DB_PASSWORD                 string
	DB_HOST                     string
	DB_USER                     string
	DB_PORT                     string
	PORT                        string
}

var gEnv *Environment

func GetEnvironment() *Environment {
	return gEnv
}

func Init() error {
	var env Environment
	envVars := map[string]*string{
		"JWT_SIGNING_KEY":             &env.JWT_SIGNING_KEY,
		"JWT_ACCESS_TOKEN_EXPIRE_MIN": &env.JWT_ACCESS_TOKEN_EXPIRE_MIN,
		"DB_NAME":                     &env.DB_NAME,
		"DB_PASSWORD":                 &env.DB_PASSWORD,
		"DB_HOST":                     &env.DB_HOST,
		"DB_USER":                     &env.DB_USER,
		"DB_PORT":                     &env.DB_PORT,
		"PORT":                        &env.PORT,
	}

	for name, ptr := range envVars {
		value := os.Getenv(name)
		if value == "" {
			logging.GetLogger().Debug().Msgf("Environment variable %s is not set", name)
			return errors.New("Not all env set.")
		}
		*ptr = value
	}

	gEnv = &env
	return nil
}
