package main

import (
	"fmt"
	"sort"
)

func main() {
	greedyChild := []int{1, 3, 2}
	breadSize := []int{1, 1}

	sort.Ints(greedyChild)
	sort.Ints(breadSize)

	j := len(breadSize) - 1
	counter := 0
	for i := len(greedyChild) - 1; i >= 0; i-- {
		if j >= 0 {
			if greedyChild[i] <= breadSize[j] {
				j--
				counter++
			}
		}
	}

	fmt.Println("children satisfied count=", counter)

}
