package main

func fibonacci(n int, answer map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}
	
	v, ok := answer[n]
	if ok {
		return v
	}
	
	f := fibonacci(n-1, answer) + fibonacci(n-2, answer)
	answer[n] = f
	return f
}