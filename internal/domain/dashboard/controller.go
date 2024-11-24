package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type DashboardController struct {
	service *DashboardService
}

func NewDashboardController() *DashboardController {
	return &DashboardController{
		service: NewDashboardService(),
	}
}

func (controller *DashboardController) get(c *gin.Context) {
	view, err := controller.service.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusOK, view)
}

func RegisterRoutes(r *gin.Engine) {
	controller := NewDashboardController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/dashboard", controller.get)
}
