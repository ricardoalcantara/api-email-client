package middlewares

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/token"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

func AuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")
	return func(c *gin.Context) {
		tokenType, authToken := getToken(c)

		if len(authToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Not authorized"})
			return
		}

		switch tokenType {
		case "ApiKey":
			authApiKey(c, authToken)
		case "Bearer":
			authBearer(c, authToken, secret)
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Not authorized"})
		}
	}
}

func authBearer(c *gin.Context, authToken string, secret string) {
	authorized, err := token.IsAuthorized(authToken, secret)
	if !authorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	accessToken, err := token.ExtractToken(authToken, secret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	claims, err := token.ExtractClaims(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.Set("x-id", claims.RegisteredClaims.Subject)
	c.Next()
}

func authApiKey(c *gin.Context, authToken string) {
	dbApiKey, err := models.ApiKeyGetByHash(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	if dbApiKey.ExpiresAt != nil && dbApiKey.ExpiresAt.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, types.ErrorResponse{Error: "Expired"})
		return
	}

	c.Set("x-id", strconv.Itoa(int(dbApiKey.UserId)))
	c.Next()
}

func getToken(c *gin.Context) (string, string) {
	authHeader := c.Request.Header.Get("Authorization")

	t := strings.Split(authHeader, " ")
	if len(t) == 2 {
		return t[0], t[1]
	}

	return "", ""
}
