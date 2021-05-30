package main

import (
	"fmt"
	"math"
)

func main() {
	list := []int{-1, 2, 1, -4}
	target := 1
	expectedClosestToTarget := 2

	closest := list[0] + list[1] + list[2]
	out:
	for i := 0; i < len(list) - 1; i++ {
		j := i + 1
		k := len(list) - 1

		for j < k {
			sum := list[i] + list[j] + list[k]

			if math.Abs(float64(sum) - float64(target)) < math.Abs(float64(closest) - float64(target)) {
					closest = sum
			}

			if sum == closest {
				break out
			} else if sum < closest {
				j++
			} else {
				k--
			}
		}
	}

	fmt.Println(closest)
	fmt.Println(closest == expectedClosestToTarget)
}