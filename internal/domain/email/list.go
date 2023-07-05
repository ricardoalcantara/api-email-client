package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/samber/lo"
)

func List(c *gin.Context) {
	p := models.NewPagination(c)
	emails, err := models.EmailList(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(emails, func(e models.Email, index int) EmailView {
		return EmailView{
			ID:       e.ID,
			SmtpName: e.Smtp.Name,
			From:     e.Smtp.Email,
			To:       e.To,
			Subject:  e.Subject,
			SentAt:   e.SentAt,
		}
	})

	c.JSON(http.StatusOK, domain.ListView[EmailView]{List: result, Page: p.Page})
}
