package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var gLogger *zerolog.Logger = nil

func GetLogger() *zerolog.Logger {
	return gLogger
}

func Init() {
	_logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	gLogger = &_logger
}
