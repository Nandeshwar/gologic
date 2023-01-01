package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := [][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	// max area of rectangle = 15
	/*
	  steps:
	   1. 1st row in tmp arr and find its max area
	   2. when next row is processed, update tmp row. increment temp row - column by 1
	      if current row's column is zero, corresponding value in temp row will be 0
	   3. and keep update max area variable and that will be answer
	*/

	var maxArea int
	tmpArr := make([]int, len(a[0]))
	copy(tmpArr, a[0])
	area := findAreaInHistogram(tmpArr)
	maxArea = max(area, maxArea)

	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == 0 {
				tmpArr[j] = 0
			} else {
				tmpArr[j] += 1
			}
		}
		area = findAreaInHistogram(tmpArr)

		maxArea = max(area, maxArea)
	}

	fmt.Println("max area of rectangle=", maxArea)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findAreaInHistogram(a []int) int {
	maxArea := 0
	stack := list.New()

	left := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		for stack.Len() != 0 {
			topElement := stack.Back().Value
			top := topElement.(int)
			if a[top] < a[i] {
				left[i] = top
				stack.PushBack(i)
				break
			} else {
				stack.Remove(stack.Back())
				continue
			}
		}

		if stack.Len() == 0 {
			stack.PushBack(i)
			left[0] = -1
		}
	}

	right := make([]int, len(a))
	stack.Init()

	for i := len(a) - 1; i >= 0; i-- {
		for stack.Len() != 0 {
			topElement := stack.Back().Value
			top := topElement.(int)
			if a[top] < a[i] {
				right[i] = top
				stack.PushBack(i)
				break
			} else {
				stack.Remove(stack.Back())
				continue
			}
		}

		if stack.Len() == 0 {
			right[i] = 5
			stack.PushBack(i)
		}
	}

	for i := 0; i < len(a); i++ {
		area := (right[i] - left[i] - 1) * a[i]
		maxArea = max(maxArea, area)
	}

	return maxArea
}
