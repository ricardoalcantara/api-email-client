package apikey

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

type ApiKeyController struct {
	service *ApiKeyService
}

func NewApiKeyController() *ApiKeyController {
	return &ApiKeyController{
		service: NewApiKeyService(),
	}
}

func (controller *ApiKeyController) post(c *gin.Context) {
	var input CreateApiKeyDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.Atoi(c.GetString("x-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	view, err := controller.service.post(uint(userId), input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, view)
}

func (controller *ApiKeyController) list(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetString("x-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.NewPagination(c)
	result, err := controller.service.list(uint(userId), p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *ApiKeyController) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	userId, err := strconv.Atoi(c.GetString("x-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	view, err := controller.service.get(uint(userId), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	c.JSON(http.StatusOK, view)
}

func (controller *ApiKeyController) regenerate(c *gin.Context) {
}

func (controller *ApiKeyController) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	userId, err := strconv.Atoi(c.GetString("x-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.delete(uint(userId), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.Status(http.StatusAccepted)
}

func RegisterRoutes(r *gin.Engine) {

	controller := NewApiKeyController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/api-key", controller.list)
	routes.POST("/api-key", controller.post)
	routes.GET("/api-key/:id", controller.get)
	routes.PATCH("/api-key/:id", controller.regenerate)
	routes.DELETE("/api-key/:id", controller.delete)
}
