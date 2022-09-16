package main

import (
	"fmt"
	"math"
)

func main() {
	building := []int{3, 1, 2, 4, 0, 1, 3, 2}
	fmt.Println(trappedWater(building))
}

func trappedWater(building []int) int {

	leftWallMaxHeights := make([]int, len(building))
	rightWallMaxHeights := make([]int, len(building))

	// preprocessing: left wall has max value
	leftWallMaxHeights[0] = building[0]
	for i := 1; i < len(building); i++ {
		if building[i] > leftWallMaxHeights[i-1] {
			leftWallMaxHeights[i] = building[i]
		} else {
			leftWallMaxHeights[i] = leftWallMaxHeights[i-1]
		}
	}

	// right wall has max value
	rightWallMaxHeights[len(rightWallMaxHeights)-1] = building[len(building)-1]
	for i := len(building) - 2; i >= 0; i-- {
		if building[i] > rightWallMaxHeights[i+1] {
			rightWallMaxHeights[i] = building[i]
		} else {
			rightWallMaxHeights[i] = rightWallMaxHeights[i+1]
		}
	}

	fmt.Println("leftWallMaxHeights=", leftWallMaxHeights)
	fmt.Println("rightWallMaxHeights=", rightWallMaxHeights)

	trappedRainWater := 0
	for i := 0; i < len(building); i++ {
		trappedRainWater += int(math.Min(float64(leftWallMaxHeights[i]), float64(rightWallMaxHeights[i]))) - building[i]
	}
	return trappedRainWater
}

/*
output:
  bash-3.2$ go run main.go
leftWallMaxHeights= [3 3 3 4 4 4 4 4]
rightWallMaxHeights= [4 4 4 4 3 3 3 2]
8
*/
