package main

import (
	"fmt"
)

func main() {
	num := 13 // 1
	//num := 14 // 4

	fmt.Println(isHappyNumber(num))

}

func isHappyNumber(num int) bool {
	var sum int
	for {

		for num != 0 {
			remainder := num % 10
			sum += remainder * remainder
			num = num / 10
		}
		num = sum

		// if sum is 1 digit break loop
		if sum < 9 {
			break
		}
		sum = 0
	}

	fmt.Println("sum=", sum)
	if sum == 1 {
		return true
	}
	return false
}
