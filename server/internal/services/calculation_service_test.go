package services_test

import (
	"packing_manager/internal/services"
	"packing_manager/internal/types"
	"testing"
)

func AssertMapsEqual(map1, map2 map[int64]int64) bool {
	if map1 == nil && map2 == nil {
		return true
	}

	if map1 == nil || map2 == nil {
		return false
	}

	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		value2, exists := map2[key]
		if !exists || value1 != value2 {
			return false
		}
	}

	return true
}
func TestCalculationService_CalculatePackages(t *testing.T) {
	calculationService := services.NewCalculationService()
	res := calculationService.CalculatePackages(12001, []types.Package{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	})
	if !AssertMapsEqual(res, map[int64]int64{
		5000: 2,
		2000: 1,
		250:  1,
	}) {
		t.Fatal("Test Packages 5000, 2000, 250; TotalItems: 12001 failed")
	}
	res = calculationService.CalculatePackages(500000, []types.Package{
		{Size: 23},
		{Size: 31},
		{Size: 53},
	})
	if !AssertMapsEqual(res, map[int64]int64{
		23: 2,
		31: 7,
		53: 9429,
	}) {
		t.Fatal("Test Packages 23, 31, 53; TotalItems: 500000 failed")
	}
	res = calculationService.CalculatePackages(5000, []types.Package{
		{Size: 5001},
		{Size: 278},
		{Size: 1111},
	})
	if !AssertMapsEqual(res, map[int64]int64{
		278:  2,
		1111: 4,
	}) {
		t.Fatal("Test Packages 5001, 278, 1111; TotalItems: 5000 failed")
	}
}
