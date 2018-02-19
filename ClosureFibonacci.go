package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	 // f1 and f2 will be treated as global in closure
	f1 := -1
	f2 := 1

	return func() int{
		f3 := f1 + f2
		f1 = f2
		f2 = f3
		return f3
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
output:
 0
1
1
2
3
5
8
13
21
34
 */