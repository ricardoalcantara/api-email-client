package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/domain/email"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/samber/lo"
)

func getEmail(c *gin.Context) {
	p := models.NewPagination(c)
	emails, err := models.EmailList(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(emails, func(e models.Email, index int) email.EmailView {
		emailView := email.EmailView{
			ID:      e.ID,
			To:      e.To,
			Subject: e.Subject,
			SentAt:  e.SentAt,
		}
		if e.Smtp != nil {
			emailView.SmtpName = e.Smtp.Name
			emailView.From = e.Smtp.Email
		}

		return emailView
	})

	c.HTML(http.StatusOK, "pages/email/index.html", gin.H{
		"listEmail": result,
	})
}
