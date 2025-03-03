package handlers

import "github.com/gin-gonic/gin"

type CalculataionHandler interface {
	CalculatePackages(c *gin.Context)
}
