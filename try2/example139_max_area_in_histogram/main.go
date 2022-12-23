package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*
					    |
					  | |
			    |     | |     |
				|     | | |   |
				| |   | | | | | |
				| | | | | | | | |
				4 2 1 5 6 3 2 4 2

		left smaller index : -1 -1 -1  2  3  2  2 6  2

						 a :  4  2  1  5  6  3  2  4  2
						      --------------------------
					  index:  0  1  2  3  4  5  6  7  8

		right smaller index:  1  2  9  5  5  6  9  8  9

					formula: (r[i] - l[i] - 1) * a[i]
					         0 : (1 - -1 - 1) * 4 = 4
							 1 : (2 - -1 - 1) * 2 = 4
							 2 : 9
							 3 : 10
							 4:  6
							 5:  9
							 6:  12
							 7: 4
							 8: 12
				max is 12
				output:
				leftSmallerIndexArr= [-1 -1 -1 2 3 2 2 6 2]
				rightSmallerIndexArr= [1 2 9 5 5 6 9 8 9]
				maxArea= 12
	*/
	a := []int{4, 2, 1, 5, 6, 3, 2, 4, 2}
	leftSmallerIndexArr := getLeftSmallerIndex(a)
	fmt.Println("leftSmallerIndexArr=", leftSmallerIndexArr)
	rightSmallerIndexArr := getRightSmallerIndex(a)
	fmt.Println("rightSmallerIndexArr=", rightSmallerIndexArr)

	maxArea := findMaxAreaInHistogram(a, leftSmallerIndexArr, rightSmallerIndexArr)
	fmt.Println("maxArea=", maxArea)

}

func getLeftSmallerIndex(a []int) []int {
	leftSmallerIndexArr := make([]int, len(a))
	ds := list.New()
	for i, v := range a {

		for ds.Len() != 0 {
			topElement := ds.Back().Value
			topIndex := topElement.(int)
			if a[topIndex] < v {
				leftSmallerIndexArr[i] = topIndex
				ds.PushBack(i)
				break
			} else {
				ds.Remove(ds.Back())
				continue
			}
		}

		if ds.Len() == 0 {
			ds.PushBack(i)
			leftSmallerIndexArr[i] = -1
			continue
		}
	}
	return leftSmallerIndexArr
}

func getRightSmallerIndex(a []int) []int {
	rightSmallerIndexArr := make([]int, len(a))
	ds := list.New()
	for i := len(a) - 1; i >= 0; i-- {

		for ds.Len() != 0 {
			topElement := ds.Back().Value
			topIndex := topElement.(int)
			if a[topIndex] < a[i] {
				rightSmallerIndexArr[i] = topIndex
				ds.PushBack(i)
				break
			} else {
				ds.Remove(ds.Back())
			}
		}

		if ds.Len() == 0 {
			ds.PushBack(i)
			rightSmallerIndexArr[i] = 9
		}
	}
	return rightSmallerIndexArr
}

func findMaxAreaInHistogram(a, left, right []int) int {
	var maxArea int
	for i := 0; i < len(a); i++ {
		area := (right[i] - left[i] - 1) * a[i]
		maxArea = max(area, maxArea)
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
