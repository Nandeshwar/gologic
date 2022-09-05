package main

import "fmt"

func main() {
	a := []int{1, 2, 1, 1, 3, 2, 1}
	// print number repeated more than n/2
	fmt.Println(nBy2RepeatedNumber(a))
}

func nBy2RepeatedNumber(a []int) int {
	cnt := 1
	ansIndex := 0
	maxRepeated := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] == a[ansIndex] {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			ansIndex = i
			cnt = 1
		}
	}

	maxRepeated = a[ansIndex]

	cnt = 0
	for _, v := range a {
		if v == maxRepeated {
			cnt++
		}
	}

	if cnt > len(a)/2 {
		return maxRepeated
	}

	return -1
}
