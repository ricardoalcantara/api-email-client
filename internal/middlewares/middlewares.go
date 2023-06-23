package middlewares

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/token"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenType, authToken := getToken(c)

		if len(authToken) == 0 || (tokenType != "Basic" && tokenType != "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not authorized"})
			return
		}

		if tokenType == "Basic" {
			decoded, err := base64.StdEncoding.DecodeString(authToken)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
				return
			}

			cred := strings.Split(string(decoded), ":")
			if len(cred) != 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not authorized"})
				return
			}

			client, err := models.LoginCheck(cred[0], cred[1])
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not authorized"})
				return
			}

			c.Set("x-user-id", strconv.Itoa(int(client.ID)))
			c.Next()
		} else {
			authorized, err := token.IsAuthorized(authToken, secret)
			if !authorized {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
				return
			}

			accessToken, err := token.ExtractToken(authToken, secret)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
				return
			}
			claims, err := token.ExtractClaims(accessToken)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
				return
			}

			c.Set("x-user-id", claims.RegisteredClaims.Subject)
			c.Next()
		}
	}
}

func getToken(c *gin.Context) (string, string) {
	authHeader := c.Request.Header.Get("Authorization")

	t := strings.Split(authHeader, " ")
	if len(t) == 2 {
		return t[0], t[1]
	}

	return "", ""
}
