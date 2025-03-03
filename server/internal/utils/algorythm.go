package utils

import (
	"math"
	"packing_manager/internal/types"
)

func CalculatePackages(totalItemsNumber int64, packages []types.Package) map[int64]int64 {
	dp := make([]int64, totalItemsNumber+1)
	res := make(map[int64]int64)
	for i := range totalItemsNumber + 1 {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0
	packageUsed := make([]int64, totalItemsNumber+1)
	for i := int64(1); i < totalItemsNumber+1; i++ {
		for _, pack := range packages {
			if i >= pack.Size && dp[i-pack.Size]+1 < dp[i] {
				dp[i] = dp[i-pack.Size] + 1
				packageUsed[i] = pack.Size
			}
		}
	}

	if dp[totalItemsNumber] == math.MaxInt32 {
		return nil
	}

	remaining := totalItemsNumber
	for remaining > 0 {
		pack := packageUsed[remaining]
		res[pack] += 1
		remaining -= pack
	}

	return res
}
