package setup

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Env() {
	err := godotenv.Load()

	if err != nil {
		log.Debug().Msg("Fail loading .env file")
	}
}
