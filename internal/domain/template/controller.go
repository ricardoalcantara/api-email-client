package template

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/template", List)
	routes.POST("/template/", Post)
	routes.GET("/template/:id", Get)
	routes.DELETE("/template/:id", Delete)
}

/*
/api/user
/api/user/me
*/

/*
/api/auth/token
/api/auth/refresh_token
/api/auth/logout
/api/auth/password_recovery
/api/auth/claims
*/
