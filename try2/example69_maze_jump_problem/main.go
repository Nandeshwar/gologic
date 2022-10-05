package main

import (
	"fmt"
)

func main() {
	result := maze(1, 1, 3, 3)
	fmt.Println(result)
}

func maze(row, col, rowLen, colLen int) []string {
	if row == rowLen && col == colLen {
		return []string{""}
	}

	var finalResult []string

	for i := 1; i <= rowLen-row; i++ {
		result := maze(i+row, col, rowLen, colLen)
		for _, v := range result {
			finalResult = append(finalResult, "r"+v)
		}
	}

	for i := 1; i <= colLen-col; i++ {
		result := maze(row, col+i, rowLen, colLen)
		for _, v := range result {
			finalResult = append(finalResult, "c"+v)
		}
	}

	for i := 1; i <= rowLen-row && i <= colLen-col; i++ {
		result := maze(i+row, i+col, rowLen, colLen)
		for _, v := range result {
			finalResult = append(finalResult, "d"+v)
		}
	}

	return finalResult

}
