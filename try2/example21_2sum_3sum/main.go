package main

import "fmt"

func main() {
	num1, num2 := twoSumWithSortedArr([]int{2, 3, 4, 7, 10}, 9)
	fmt.Println(num1, num2)

	num1, num2 = twoSumWithUnSortedArr([]int{7, 2, 10, 3, 4}, 9)
	fmt.Println(num1, num2)

	num1, num2, num3 := threeSum([]int{1, 1, 4, 7, 10}, 9)
	fmt.Println(num1, num2, num3)

	fmt.Println("Three sum algo2")
	num1, num2, num3 = threeSumAlgo2([]int{1, 1, 4, 7, 10}, 9)
	fmt.Println(num1, num2, num3)
}

func twoSumWithSortedArr(arr []int, target int) (int, int) {

	for i, j := 0, len(arr)-1; i < j; {
		if target < arr[i]+arr[j] {
			j--
		} else if target > arr[i]+arr[j] {
			i++
		} else if target == arr[i]+arr[j] {
			return arr[i], arr[j]
		}
	}
	return -1, -1
}

func twoSumWithUnSortedArr(arr []int, target int) (int, int) {
	m := map[int]int{}
	for i := 0; i < len(arr); i++ {

		v, ok := m[target-arr[i]]
		if ok {
			return arr[i], v
		}
		m[arr[i]] = arr[i]
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

func threeSumAlgo2(a []int, target int) (int, int, int) {

	for i := 0; i < len(a)-2; i++ {
		j := i + 1
		k := len(a) - 1

		for j < k {
			sumHere := a[i] + a[j] + a[k]
			if sumHere == target {
				return a[i], a[j], a[k]
			} else if sumHere > target {
				k--
			} else {
				i++
			}
		}
	}

	return -1, -1, -1
}
