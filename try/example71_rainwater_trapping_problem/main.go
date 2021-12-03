package main

import (
	"github.com/logic-building/functional-go/fp"
)

func getRainWaterTrappedBlock(buildingHeight []int) int {
	totalGap := 0
	left := make([]int, len(buildingHeight))
	right := make([]int, len(buildingHeight))

	left[0] = buildingHeight[0]
	right[len(buildingHeight)-1] = buildingHeight[len(buildingHeight)-1]

	for i := 1; i < len(buildingHeight); i++ {
		left[i] = fp.MaxInt([]int{buildingHeight[i], buildingHeight[i-1]})
	}

	for i := len(buildingHeight) - 2; i >= 0; i-- {
		right[i] = fp.MaxInt([]int{buildingHeight[i], buildingHeight[i+1]})
	}

	for i := 0; i < len(buildingHeight); i++ {
		gap := fp.MinInt([]int{left[i], right[i]}) - buildingHeight[i]
		if gap > 0 {

			totalGap += gap
		}
	}
	return totalGap
}
