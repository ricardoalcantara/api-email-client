package setup

import (
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func Env() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
