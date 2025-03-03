package routes

import (
	"packing_manager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine, calculationHandler handlers.CalculataionHandler) {
	r.GET("/calculate", calculationHandler.CalculatePackages)
}
