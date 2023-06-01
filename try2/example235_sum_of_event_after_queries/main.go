package main

import (
	"fmt"
)

func main() {
	
	
	a := []int{1, 2, 3, 4}
	
	/*
		loop queries:
		   1st is value 
		   2nd is index
		   
		add query value to array index of query if even
	*/
	queries := [][]int{
		{1, 0},
		{-3, 1},
		{-4, 0},
		{2, 3},
	}
	finalArray := sumOfEventAfterQueries(a, queries)
	fmt.Println("result array=", finalArray)

	// expectation: [8,6,2,4]
}

func sumOfEventAfterQueries(a []int, queries [][]int) []int {
	var result []int
	var sum int

	for _, v := range a {
		if v&1 == 0 {
			sum += v
		}
	}

	for _, at := range queries {
		i := at[1]
		v := at[0]

		if a[i]&1 == 0 {   // subtract old value if even
			sum -= a[i]   
		}
		a[i] += v   // add new value

		if a[i]&1 == 0 {  // if new value is even, add in result
			sum += a[i]
		}

		result = append(result, sum)
	}

	return result
}
