package main

import (
	"packing_manager/internal/handlers"
	"packing_manager/internal/routes"
	"packing_manager/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	calculationService := services.NewCalculationService()
	calculationHandler := handlers.NewCalculateHandler(calculationService)
	routes.NewRoutes(r, calculationHandler)
	r.Run(":8080")
}
