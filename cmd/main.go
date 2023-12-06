package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/setup"

	frontend_about "github.com/ricardoalcantara/api-email-client/internal/frontend/about"
	frontend_email "github.com/ricardoalcantara/api-email-client/internal/frontend/email"
	frontend_home "github.com/ricardoalcantara/api-email-client/internal/frontend/home"
	frontend_login "github.com/ricardoalcantara/api-email-client/internal/frontend/login"
	frontend_smtp "github.com/ricardoalcantara/api-email-client/internal/frontend/smtp"
	frontend_template "github.com/ricardoalcantara/api-email-client/internal/frontend/template"
)

func init() {
	setup.Env()
	models.ConnectDataBase()
	emailengine.Create()
}

func main() {

	console()

	// api := gin.New()
	// api.Use(
	// 	gin.LoggerWithWriter(gin.DefaultWriter, "/healthcheck"),
	// 	gin.Recovery(),
	// )

	// api.GET("/healthcheck", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Ok",
	// 	})
	// })

	// auth.RegisterRoutes(api)
	// email.RegisterRoutes(api)
	// smtp.RegisterRoutes(api)
	// template.RegisterRoutes(api)

	// if host, ok := os.LookupEnv("HOST"); ok {
	// 	port := utils.GetEnv("PORT", "3000")
	// 	api.Run(host + ":" + port)
	// } else {
	// 	api.Run()
	// }
}

func console() {
	console := gin.New()
	console.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthcheck"),
		gin.Recovery(),
	)

	console.Static("/assets", "./assets")

	var files []string

	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	console.LoadHTMLFiles(files...)
	// console.LoadHTMLGlob("templates/**/**")

	var sessionSecret string
	var ok bool
	if sessionSecret, ok = os.LookupEnv("SESSION_SECRET"); !ok {
		panic("SESSION_SECRET must be set")
	}
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{MaxAge: int(time.Duration(time.Minute * 15).Seconds())})
	console.Use(sessions.Sessions("console_session", store))
	console.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})
	console.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "pages/errors/404.html", nil)
	})

	frontend_about.RegisterRoutes(console)
	frontend_email.RegisterRoutes(console)
	frontend_home.RegisterRoutes(console)
	frontend_login.RegisterRoutes(console)
	frontend_smtp.RegisterRoutes(console)
	frontend_template.RegisterRoutes(console)

	console.Run("localhost:5555")
}
