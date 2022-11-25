package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{
		{1, 5},
		{2, 4},
		{6, 9},
		{7, 8},
		{7, 10},
		{10, 11},
		{12, 13},
		{14, 15},
	}

	mergedOverlaps := mergeOverlaps(a)
	fmt.Println("mergedOverlaps=", mergedOverlaps)
}

func mergeOverlaps(a [][]int) [][]int {
	var result [][]int

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	goodRecord := a[0]
	goodEnd := goodRecord[1]

	for i := 1; i < len(a); i++ {
		// if 2nd start date < 1st end date - overlap detected
		if a[i][0] <= a[i-1][1] {
			goodEnd = max(goodEnd, a[i][1])

			// corner case: 1st is good record, then overlap and last record is overlap --- add good record
			// update endTime
			if i == len(a)-1 {
				goodRecord[1] = goodEnd
				result = append(result, goodRecord)
			}

		} else {
			// if no overlap - add 1st good record
			// update endTime
			goodRecord[1] = goodEnd
			result = append(result, goodRecord)
			goodRecord = a[i]
			goodEnd = goodRecord[1]

			// corner case: if last record is good record -  add good record
			// update endTime
			if i == len(a)-1 {
				result = append(result, goodRecord)
			}

		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
