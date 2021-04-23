// https://www.youtube.com/watch?v=icoql2WKmbA&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=2
package main

import (
	"fmt"
)

func main() {
	const n = 5

	outerList := [n][]int{}

	firstList := []int{1}
	outerList[0] = firstList

	
	for i := 1; i < n; i++ {
		newList := []int {1}
		for j := 1; j < i; j++ {
			prevRow := outerList[i-1]

			newList = append(newList, prevRow[j-1] + prevRow[j])
		} 
		newList = append(newList, 1)
		outerList[i] = newList

	}

	fmt.Println(outerList) // [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]

}