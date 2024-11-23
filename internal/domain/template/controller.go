package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
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
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *TemplateController) post(c *gin.Context) {
	var input types.CreateTemplateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	templateDto, err := controller.service.Create(&input)
	if err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, templateDto)
}

func (controller *TemplateController) generator(c *gin.Context) {
	var input types.RequestTemplateGeneratorDto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	templateGeneratorDto, err := controller.service.Generator(input)
	if err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, templateGeneratorDto)
}

func (controller *TemplateController) patch(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	var input types.UpdateTemplateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	templateDto, err := controller.service.Patch(slug, &input)
	if err != nil {
		errId := uuid.New()
		log.Error().Str("error_id", errId.String()).Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "Internal Server Error: " + errId.String()})
		return
	}

	c.JSON(http.StatusAccepted, templateDto)
}

func (controller *TemplateController) put(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	var input types.CreateTemplateDto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	templateDto, err := controller.service.Update(slug, &input)
	if err != nil {
		log.Debug().Err(err).Msg("Error")
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, templateDto)
}

func (controller *TemplateController) get(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "id/slug is required"})
		return
	}

	template, err := controller.service.Get(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, template)
}

func (controller *TemplateController) delete(c *gin.Context) {
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

	controller := NewTemplateController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/template", controller.list)
	routes.POST("/template", controller.post)
	routes.POST("/template/generator", controller.generator)
	routes.GET("/template/:slug", controller.get)
	routes.DELETE("/template/:slug", controller.delete)
	routes.PATCH("/template/:slug", controller.patch)
	routes.PUT("/template/:slug", controller.put)
}
