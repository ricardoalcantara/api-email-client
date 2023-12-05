package login

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/login", getLogin)
	r.POST("/login", postLogin)
	r.GET("/logout", getLogout)
}
