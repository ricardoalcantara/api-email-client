package smtp

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/smtp", List)
	routes.POST("/smtp/", Post)
	routes.GET("/smtp/:id", Get)
	routes.DELETE("/smtp/:id", Delete)
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
