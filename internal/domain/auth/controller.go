package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	routes.GET("/auth/token", Token)
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
