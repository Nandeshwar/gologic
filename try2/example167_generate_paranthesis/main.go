package main

import "fmt"

func main() {
	n := 3
	result := []string{}
	generateParanthesis(n, "(", 1, 0, &result)
	fmt.Println("final result=", result)
	/*
				output
		((()))
		(()())
		(())()
		()(())
		()()()
		final result= [((())) (()()) (())() ()(()) ()()()]
	*/
}

func generateParanthesis(n int, bracket string, openingBracketCount, closingBracketCount int, result *[]string) {
	if len(bracket) == 2*n {
		*result = append(*result, bracket)
		fmt.Println(bracket)
		return
	}

	if openingBracketCount < n {
		generateParanthesis(n, bracket+"(", openingBracketCount+1, closingBracketCount, result)
	}

	if closingBracketCount < openingBracketCount {
		generateParanthesis(n, bracket+")", openingBracketCount, closingBracketCount+1, result)

	}
}
