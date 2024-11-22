package smtp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
	"github.com/rs/zerolog/log"
)

type SmtpController struct {
	service *SmtpService
}

func NewSmtpController() *SmtpController {
	return &SmtpController{
		service: NewSmtpService(),
	}
}

func (controller *SmtpController) list(c *gin.Context) {
	p := models.NewPagination(c)

	result, err := controller.service.List(p)
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *SmtpController) post(c *gin.Context) {
	var input types.CreateSmtpDto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	smtpDto, err := controller.service.Create(&input)
	if err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, smtpDto)
}

func (controller *SmtpController) patch(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	var input types.UpdateSmtpDto
	if err := c.ShouldBindJSON(&input); err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	smtpDto, err := controller.service.Patch(slug, &input)
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusAccepted, smtpDto)
}

func (controller *SmtpController) put(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	var input types.UpdateSmtpDto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	smtpDto, err := controller.service.Update(slug, &input)
	if err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, smtpDto)
}

func (controller *SmtpController) get(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	smtp, err := controller.service.Get(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, smtp)
}

func (controller *SmtpController) delete(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	err := controller.service.Delete(slug)
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

func RegisterRoutes(r *gin.Engine) {
	controller := NewSmtpController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/smtp", controller.list)
	routes.POST("/smtp/", controller.post)
	routes.GET("/smtp/:slug", controller.get)
	routes.DELETE("/smtp/:slug", controller.delete)
	routes.PATCH("/smtp/:slug", controller.patch)
	routes.PUT("/smtp/:slug", controller.put)
}
