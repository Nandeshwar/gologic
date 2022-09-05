package main

import (
	"fmt"
)

func main() {
	fmt.Println("all subsets=")
	findSubsets([]rune("abc"), 0, []string{})
	fmt.Println("-----------")
	findSubsets2([]rune("abc"), 0, [][]string{})
	fmt.Println("-----return result...")

	for _, v := range findSubsetsReturn([]rune("abc"), 0, []string{}, [][]string{}) {
		fmt.Println(v)
	}

	fmt.Println("return result with for loop....without recursion")

	for _, v := range findSubsetsReturnWithLoop([]rune("abc"), 0, []string{}) {
		fmt.Println(v)
	}

	fmt.Println("\n find3....")
	findSubsets3("abc", 0, "")
}

func findSubsets3(input string, i int, result string) {
	if i == len(input) {
		fmt.Println(result)
		return
	}

	appendedResult := result + string(input[i])
	findSubsets3(input, i+1, appendedResult)
	findSubsets3(input, i+1, result)
}

func findSubsets(input []rune, i int, result []string) {
	if i == len(input) {
		fmt.Println(result)
		return
	}

	appendedResult := append(result, string(input[i]))
	findSubsets(input, i+1, appendedResult)
	findSubsets(input, i+1, result)
}

func findSubsets2(input []rune, i int, result [][]string) {
	if i == len(input) {
		fmt.Println(result)
		return
	}

	appendedResult := append(result, []string{string(input[i])})
	findSubsets2(input, i+1, appendedResult)
	findSubsets2(input, i+1, result)
}

func findSubsetsReturn(input []rune, i int, result []string, finalResult [][]string) [][]string {
	if i == len(input) {
		finalResult = append(finalResult, result)
		return finalResult
	}

	appendedResult := append(result, string(input[i]))
	findSubsets(input, i+1, appendedResult)
	findSubsets(input, i+1, result)
	return finalResult
}

func findSubsetsReturnWithLoop(input []rune, i int, result []string) [][]string {
	outerLoop := [][]string{
		{},
	}

	for _, v := range input {
		l := len(outerLoop)
		for j := 0; j < l; j++ {
			tmpArr := outerLoop[j]
			tmpArr = append(tmpArr, string(v))
			outerLoop = append(outerLoop, tmpArr)
		}
	}
	return outerLoop
}
