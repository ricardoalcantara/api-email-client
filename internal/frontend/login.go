package frontend

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

func getLogin(c *gin.Context) {
	redirectTo := c.Query("redirectTo")
	c.HTML(http.StatusOK, "pages/login.html", gin.H{
		"redirectTo": redirectTo,
	})
}

func postLogin(c *gin.Context) {
	clientId := c.PostForm("clientId")
	clientSecret := c.PostForm("clientSecret")

	client, err := models.LoginCheck(clientId, clientSecret)
	if err != nil {
		c.HTML(http.StatusOK, "pages/login.html", gin.H{
			"clientId": clientId,
			"error":    "Invalid credentials",
		})
		return
	}

	redirectTo := c.PostForm("redirectTo")

	session := sessions.Default(c)

	session.Set("clientId", client.ClientId)
	session.Set("id", client.ID)
	session.Save()

	if len(redirectTo) > 0 {
		c.Redirect(http.StatusSeeOther, "/"+redirectTo[1:])
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func getLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}
