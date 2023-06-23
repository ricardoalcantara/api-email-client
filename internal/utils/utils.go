package utils

import (
	"encoding/json"
	"math/rand"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenString(n uint8) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}

func GetEnvOr(key string, fallback func() string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback()
	}
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func PrintError(err error) string {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		logrus.WithFields(logrus.Fields{
			"filename": filename,
			"line":     line,
		}).Error(err)
		return err.Error()
	}
	return ""
}

func PrintErrorAnd(err error, message string) string {
	PrintError(err)
	return message
}
