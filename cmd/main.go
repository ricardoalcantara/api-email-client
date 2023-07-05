package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/setup"
	"github.com/ricardoalcantara/api-email-client/internal/utils"

	"github.com/ricardoalcantara/api-email-client/internal/domain/auth"
	"github.com/ricardoalcantara/api-email-client/internal/domain/email"
	"github.com/ricardoalcantara/api-email-client/internal/domain/smtp"
	"github.com/ricardoalcantara/api-email-client/internal/domain/template"
)

func init() {
	setup.Env()
	models.ConnectDataBase()
	emailengine.Create()
}

func main() {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	auth.RegisterRoutes(r)
	email.RegisterRoutes(r)
	smtp.RegisterRoutes(r)
	template.RegisterRoutes(r)

	if host, ok := os.LookupEnv("HOST"); ok {
		port := utils.GetEnv("PORT", "3000")
		r.Run(host + ":" + port)
	} else {
		r.Run()
	}
}
