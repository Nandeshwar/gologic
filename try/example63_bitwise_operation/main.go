package main

// https://www.youtube.com/watch?v=Db8OmMfzwl8package main

func multiplyBy2(num int) int {
	return num << 1
}

func divideBy2(num int) int {
	return num >> 1
}

func isEven(num int) bool {
	result := num & 1
	if result == 0 {
		return true
	}
	return false
}

func swapNumber(num1, num2 int) (int, int) {
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	return num1, num2
}
