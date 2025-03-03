package services

import "packing_manager/internal/types"

type CalculateSerevice interface {
	CalculatePackages(totalItemsNumber int64, packages []types.Package) map[int64]int64
}
