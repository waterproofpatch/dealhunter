package logging

import (
	"os"

	"github.com/rs/zerolog"
)

func Init() *zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	return &logger
}
