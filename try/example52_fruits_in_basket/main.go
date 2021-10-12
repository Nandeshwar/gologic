/*
  input:[3, 3, 3, 1,  2, 1, 1, 2, 3, 3, 4]
  output: 5

  a. find two different number pattern together
  b. then find max

  3 3 3 1     - 4
  1 2 1 1 2   - 5
  3 3 4       - 3

*/

package main

import (
	"math"
)

func findFruitsNumberInBasket(trees []int) int {
	max := 0
	currentMax := 0
	lastCnt := 0

	lastItem := -1
	secondLastItem := -1

	for _, currentItem := range trees {

		if currentItem != lastItem && currentItem != secondLastItem {
			currentMax = lastCnt
		}
		if currentItem != lastItem {
			lastCnt = 0
		}
		lastCnt++

		currentMax++
		max = int(math.Max(float64(currentMax), float64(max)))

		if currentItem != lastItem {
			secondLastItem = lastItem
			lastItem = currentItem
		}

	}
	return max
}
