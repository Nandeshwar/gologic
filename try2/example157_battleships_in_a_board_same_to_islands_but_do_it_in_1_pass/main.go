package main

import (
	"fmt"
)

func main() {
	a := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}

	// can be done with depth first search like count number of islands

	count := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i][j] == '.' {
				continue
			}

			if i > 0 && a[i-1][j] == 'X' {
				continue
			}

			if j > 0 && a[i][j-1] == 'X' {
				continue
			}
			count++
		}
	}

	fmt.Println("count=", count)
}
