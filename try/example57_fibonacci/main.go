package main


func fibonacci(n, f1, f2 int, result []int) []int {
	if n == 0 {
		return result
	}

	f3 := f1 + f2
	f1 = f2
	f2 = f3

	result = append(result, f3)
	result = fibonacci(n-1, f1, f2, result)

	return result
}
