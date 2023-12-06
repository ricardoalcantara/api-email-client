package smtp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/domain/smtp"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/samber/lo"
)

func getIndex(c *gin.Context) {
	p := models.NewPagination(c)
	smtps, err := models.SmtpList(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(smtps, func(s models.Smtp, index int) smtp.SmtpView {
		return smtp.SmtpView{
			ID:      s.ID,
			Name:    s.Name,
			Server:  s.Server,
			Port:    s.Port,
			Email:   s.Email,
			User:    s.User,
			Default: s.Default,
		}
	})

	c.HTML(http.StatusOK, "pages/smtp/index.html", gin.H{
		"listSmtp": result,
	})
}
