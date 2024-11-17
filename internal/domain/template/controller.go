package template

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/rs/zerolog/log"
)

type TemplateController struct {
	service *TemplateService
}

func NewTemplateController() *TemplateController {
	return &TemplateController{
		service: NewTemplateService(),
	}
}

func (controller *TemplateController) list(c *gin.Context) {
	p := models.NewPagination(c)

	result, err := controller.service.List(p)
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *TemplateController) post(c *gin.Context) {
	var input CreateTemplateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	templateDto, err := controller.service.Create(&input)
	if err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, templateDto)
}

func (controller *TemplateController) patch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	var input UpdateTemplateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	templateDto, err := controller.service.Update(uint(id), &input)
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusAccepted, templateDto)
}

func (controller *TemplateController) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}
	template, err := controller.service.Get(uint(id))
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusOK, template)
}

func (controller *TemplateController) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}
	err = controller.service.Delete(uint(id))
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

func RegisterRoutes(r *gin.Engine) {

	controller := NewTemplateController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/template", controller.list)
	routes.POST("/template", controller.post)
	routes.GET("/template/:id", controller.get)
	routes.DELETE("/template/:id", controller.delete)
	routes.PATCH("/template/:id", controller.patch)
}
