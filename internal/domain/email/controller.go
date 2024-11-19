package email

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

type EmailController struct {
	service *EmailService
}

func NewEmailController() *EmailController {
	return &EmailController{
		service: NewEmailService(),
	}
}

func (controller *EmailController) post(c *gin.Context) {
	var input SendEmailDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.post(input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func (controller *EmailController) list(c *gin.Context) {
	p := models.NewPagination(c)
	result, err := controller.service.list(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *EmailController) get(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.service.get(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusOK, view)
}

func RegisterRoutes(r *gin.Engine) {

	controller := NewEmailController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/email", controller.list)
	routes.POST("/email", controller.post)
	routes.GET("/email/:id", controller.get)
}
