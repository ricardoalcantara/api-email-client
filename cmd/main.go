package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain/auth"
	"github.com/ricardoalcantara/api-email-client/internal/domain/email"
	"github.com/ricardoalcantara/api-email-client/internal/domain/smtp"
	"github.com/ricardoalcantara/api-email-client/internal/domain/template"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/frontend"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/setup"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func init() {
	setup.Env()
	models.ConnectDataBase()
	emailengine.Create()
}

func main() {

	console := gin.Default()

	console.Static("/assets", "./assets")
	console.LoadHTMLGlob("templates/**/*.html")

	var sessionSecret string
	var ok bool
	if sessionSecret, ok = os.LookupEnv("SESSION_SECRET"); !ok {
		panic("SESSION_SECRET must be set")
	}
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{MaxAge: int(time.Duration(time.Minute * 15).Seconds())})
	console.Use(sessions.Sessions("console_session", store))
gin.Recovery()
	console.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})
	frontend.RegisterRoutes(console)

	go console.Run("localhost:5555")

	api := gin.Default()

	api.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	auth.RegisterRoutes(api)
	email.RegisterRoutes(api)
	smtp.RegisterRoutes(api)
	template.RegisterRoutes(api)

	if host, ok := os.LookupEnv("HOST"); ok {
		port := utils.GetEnv("PORT", "3000")
		api.Run(host + ":" + port)
	} else {
		api.Run()
	}
}
