package services

import (
	"packing_manager/internal/types"
	"packing_manager/internal/utils"
	"sort"
)

type CalculateServiceImpl struct {
}

func (c CalculateServiceImpl) CalculatePackages(totalItemsNumber int64, packages []types.Package) map[int64]int64 {
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Size > packages[j].Size
	})
	res := utils.CalculatePackages(totalItemsNumber, packages)
	i := int64(1)
	for res == nil {
		res = utils.CalculatePackages(totalItemsNumber+i, packages)
		i += 1
	}
	return res
}

func NewCalculationService() CalculateSerevice {
	return &CalculateServiceImpl{}
}
