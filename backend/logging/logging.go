package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var gLogger zerolog.Logger

func GetLogger() *zerolog.Logger {
	return &gLogger
}

func Init() {
	gLogger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}
