package main

import "fmt"

func main() {
	num1, num2 := twoSumWithSortedArr([]int{7, 2, 10, 3, 4}, 9)
	fmt.Println(num1, num2)

	num1, num2 = twoSumWithUnSortedArr([]int{7, 2, 10, 3, 4}, 9)
	fmt.Println(num1, num2)

	num1, num2, num3 := threeSum([]int{7, 1, 10, 1, 4}, 9)
	fmt.Println(num1, num2, num3)
}

func twoSumWithSortedArr(arr []int, target int) (int, int) {

	for i, j := 0, len(arr)-1; i < j; {
		if arr[i]+arr[j] < target {
			i++
		} else if arr[i]+arr[j] > target {
			j--
		} else if arr[i]+arr[j] == target {
			return arr[i], arr[j]
		}
	}
	return -1, -1
}

func twoSumWithUnSortedArr(arr []int, target int) (int, int) {
	m := map[int]int{}
	for i := 0; i < len(arr); i++ {
		v := target - arr[i]
		_, ok := m[arr[i]]
		if ok {
			return v, arr[i]
		}
		m[v] = arr[i]
	}
	return -1, -1
}

func threeSum(arr []int, target int) (int, int, int) {
	for i := 0; i < len(arr); i++ {
		a, b := twoSumWithSortedArr(arr[i+1:], target-arr[i])
		if a+b+arr[i] == target {
			return a, b, arr[i]
		}
	}
	return -1, -1, -1
}
