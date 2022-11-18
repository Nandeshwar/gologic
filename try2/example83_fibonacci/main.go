package main

import "fmt"

func main() {
	fmt.Println(fib1(10, -1, 1))
	fmt.Println(fib2(10, -1, 1, []int{}))
	fmt.Println(fib3(9, map[int]int{}))
	fmt.Println("optimized algorithm")
	fmt.Println("\n", fib(10))
}

func fib(n int) int {
	a := 0
	b := 1
	c := -1

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	fmt.Print(0)
	fmt.Print(" ", 1, " ")
	n = n - 2
	for n >= 1 {
		n--

		c = a + b
		a = b
		b = c
		fmt.Print(" ", c)
	}
	return c
}

func fib1(n, v1, v2 int) []int {
	if n == 0 {
		return []int{}
	}

	f1 := v1
	f2 := v2
	f3 := f1 + f2
	result := []int{f3}

	result = append(result, fib1(n-1, f2, f3)...)
	return result

}

func fib2(n, f1, f2 int, result []int) []int {
	if n == 0 {
		return result
	}

	f3 := f1 + f2
	f1 = f2
	f2 = f3

	result = append(result, f3)
	result = fib2(n-1, f1, f2, result)

	return result
}

func fib3(n int, answer map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	v, ok := answer[n]
	if ok {
		return v
	}

	f := fib3(n-1, answer) + fib3(n-2, answer)
	answer[n] = f
	return f
}
