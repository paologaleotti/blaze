package util

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("ENVIRONMENT") == "prod" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)

	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		out := zerolog.ConsoleWriter{Out: os.Stdout}
		out.FormatMessage = func(i any) string {
			return fmt.Sprintf("%+v", i)
		}
		log.Logger = log.Output(out)
	}
}
