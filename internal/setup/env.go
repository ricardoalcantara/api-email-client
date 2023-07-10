package setup

import (
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func Env() {
	err := godotenv.Load()

	if err != nil {
		logrus.Info("Fail loading .env file")
	}
}
