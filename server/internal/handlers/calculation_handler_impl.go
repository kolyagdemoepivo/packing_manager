package handlers

import (
	"net/http"
	"packing_manager/internal/services"
	"packing_manager/internal/types"
	"packing_manager/internal/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CalculataionHandlerImpl struct {
	CalculationService services.CalculateSerevice
}

func (ca CalculataionHandlerImpl) CalculatePackages(c *gin.Context) {
	totalItemsStr, exists := c.GetQuery("totalItems")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'totalItems' is required",
		})
		return
	}
	packagesStr, exists := c.GetQuery("packages")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'totalItems' is required",
		})
		return
	}

	totalItems, err := strconv.Atoi(totalItemsStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'totalItems' parameter: must be an integer",
		})
		return
	}

	if totalItems <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'totalItems' parameter: must be a positive integer",
		})
		return
	}

	splitedPackageStr := strings.Split(packagesStr, ",")
	splitedPackageStr = utils.RemoveDuplicates(splitedPackageStr)
	var packages []types.Package
	for _, pacakgeStr := range splitedPackageStr {
		pack, err := strconv.Atoi(pacakgeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid 'packages' parameter: must be a list of integers",
			})
			return
		}
		if pack <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid 'packages' parameter: must be a list of positive integers",
			})
			return
		}
		packages = append(packages, types.Package{Size: int64(pack)})
	}
	res := ca.CalculationService.CalculatePackages(int64(totalItems), packages)
	c.JSON(http.StatusOK, res)
}

func NewCalculateHandler(calculationService services.CalculateSerevice) CalculataionHandler {
	return &CalculataionHandlerImpl{CalculationService: calculationService}
}
