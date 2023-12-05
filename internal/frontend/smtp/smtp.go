package smtp

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/domain/smtp"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func getSmtp(c *gin.Context) {
	p := models.NewPagination(c)
	smtps, err := models.SmtpList(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(smtps, func(s models.Smtp, index int) smtp.SmtpView {
		return smtp.SmtpView{
			ID:      s.ID,
			Name:    s.Name,
			Server:  s.Server,
			Port:    s.Port,
			Email:   s.Email,
			User:    s.User,
			Default: s.Default,
		}
	})

	c.HTML(http.StatusOK, "pages/smtp/index.html", gin.H{
		"listSmtp": result,
	})
}

func postSmtp(c *gin.Context) {
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
		c.HTML(http.StatusOK, "pages/smtp/register.html", gin.H{
			"name":         name,
			"email":        email,
			"server":       server,
			"port":         port,
			"user":         user,
			"make_default": make_default,
		})
		return
	}

	smtp := models.Smtp{
		Name:     name,
		Server:   server,
		Port:     uint16(iPort),
		Email:    email,
		User:     user,
		Password: password,
		Default:  make_default,
	}

	smtp.Base64Password()
	err = smtp.Save()
	if err != nil {
		c.HTML(http.StatusOK, "pages/smtp/register.html", gin.H{
			"name":         name,
			"email":        email,
			"server":       server,
			"port":         port,
			"user":         user,
			"make_default": make_default,
		})
		return
	}

	c.Redirect(http.StatusFound, "/smtp")
}

func deleteSmtp(c *gin.Context) {
	smtpId := c.PostForm("smtp_id")
	if len(smtpId) > 0 {
		id, err := strconv.Atoi(smtpId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
			return
		}
		models.SmtpDeleteById(uint(id))
	}

	c.Redirect(http.StatusFound, "/smtp")
}
