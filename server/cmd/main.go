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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	calculationService := services.NewCalculationService()
	calculationHandler := handlers.NewCalculateHandler(calculationService)
	routes.NewRoutes(r, calculationHandler)
	r.Run(":8080")
}
