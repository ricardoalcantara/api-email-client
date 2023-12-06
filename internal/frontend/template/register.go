package template

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/sirupsen/logrus"
)

func getRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/template/register.html", nil)
}

func postTemplate(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	server := c.PostForm("server")
	port := c.PostForm("port")
	user := c.PostForm("user")
	password := c.PostForm("password")
	make_default := c.PostForm("make_default") == "on"

	logrus.Info(c.PostForm("make_default"))

	if make_default {
		if err := models.SmtpDisableDefault(); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
			return
		}
	}

	iPort, err := strconv.Atoi(port)
	if err != nil {
		c.HTML(http.StatusOK, "pages/template/register.html", gin.H{
			"name":         name,
			"email":        email,
			"server":       server,
			"port":         port,
			"user":         user,
			"make_default": make_default,
		})
		return
	}

	template := models.Smtp{
		Name:     name,
		Server:   server,
		Port:     uint16(iPort),
		Email:    email,
		User:     user,
		Password: password,
		Default:  make_default,
	}

	template.Base64Password()
	err = template.Save()
	if err != nil {
		c.HTML(http.StatusOK, "pages/template/register.html", gin.H{
			"name":         name,
			"email":        email,
			"server":       server,
			"port":         port,
			"user":         user,
			"make_default": make_default,
		})
		return
	}

	c.Redirect(http.StatusFound, "/template")
}
