package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, 2, 1, -4} // ans 2
	targetSum := 1
	result := threeSumClosestNumber(a, targetSum)
	fmt.Println("result=", result)
}

func threeSumClosestNumber(a []int, targetSum int) int {
	sort.Ints(a)

	fmt.Println("sorted Arr=", a)
	result := a[0] + a[1] + a[2]

	for i := 0; i < len(a)-2; i++ {
		j := i + 1
		k := len(a) - 1

		for j < k {
			sumHere := a[i] + a[j] + a[k]

			if Abs(sumHere-targetSum) < Abs(result-targetSum) {
				result = sumHere
			}

			if targetSum == sumHere {
				return sumHere
			} else if targetSum < sumHere {
				k--
			} else if targetSum > sumHere {
				j++
			}

			// if sumHere == targetSum {
			// 	return sumHere
			// } else if sumHere < targetSum {
			// 	j++
			// } else if sumHere > targetSum {
			// 	k--
			// }
		}
	}
	return result
}

func Abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}
