package smtp

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

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
