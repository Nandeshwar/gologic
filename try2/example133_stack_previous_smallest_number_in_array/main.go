package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := []int{4, 10, 5, 15, 3}
	// expected output
	// -1, 4, 4, 5, -1

	fmt.Println(getPreviousSmallestElement(a))
}

func getPreviousSmallestElement(a []int) []int {
	var result []int
	stack := list.New()

	for _, v := range a {
		for stack.Len() != 0 {
			element := stack.Back().Value
			top := element.(int)
			if top < v {
				result = append(result, top)
				stack.PushBack(v)
				break
			} else {
				stack.Remove(stack.Back())
			}
		}

		if stack.Len() == 0 {
			result = append(result, -1)
			stack.PushBack(v)
			continue
		}
	}
	return result
}

/*
1.
   Previous smallest number
bash-3.2$ go run main.go
input: a := []int{4, 10, 5, 15, 3}
[-1 4 4 5 -1]

2. Previous greater number:
    Small change in if logic, instead of < than check put > than check

3. next smallest number
   Similar logic but traverse list from end

*/
