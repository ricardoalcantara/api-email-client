package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	apikey "github.com/ricardoalcantara/api-email-client/internal/domain/api_key"
	"github.com/ricardoalcantara/api-email-client/internal/domain/auth"
	"github.com/ricardoalcantara/api-email-client/internal/domain/dashboard"
	"github.com/ricardoalcantara/api-email-client/internal/domain/email"
	"github.com/ricardoalcantara/api-email-client/internal/domain/smtp"
	"github.com/ricardoalcantara/api-email-client/internal/domain/template"
	"github.com/ricardoalcantara/api-email-client/internal/domain/user"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	log.Logger = log.
		With().
		Caller().
		Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load()
	if err != nil {
		log.Debug().Msg("Fail loading .env file")
	}
	models.ConnectDataBase()
	emailengine.Create()
}

func main() {
	api := gin.New()
	api.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthcheck"),
		gin.Recovery(),
	)

	api.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	auth.RegisterRoutes(api)
	email.RegisterRoutes(api)
	smtp.RegisterRoutes(api)
	template.RegisterRoutes(api)
	apikey.RegisterRoutes(api)
	dashboard.RegisterRoutes(api)
	user.RegisterRoutes(api)

	host := utils.GetEnv("API_HOST", "")
	port := utils.GetEnv("API_PORT", "5555")
	api.Run(host + ":" + port)
}
