package email

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
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
	var input types.SendEmailDto

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
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.service.get(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusOK, view)
}

func (controller *EmailController) send(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	email, err := controller.service.send(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusAccepted, email)
}

func RegisterRoutes(r *gin.Engine) {

	controller := NewEmailController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/email", controller.list)
	routes.POST("/email", controller.post)
	routes.PATCH("/email/:id/send", controller.send)
	routes.GET("/email/:id", controller.get)
}
