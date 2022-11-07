package main

import (
	"fmt"
)

var m = [][]byte{
	{'X', 'X', 'X', 'X'},
	{'X', 'O', 'O', 'X'},
	{'O', 'X', 'O', 'X'},
	{'X', 'O', 'X', 'X'},
}

/*
	output:
	X	X	X	X
	X	O	O	X
	O	X	O	X
	X	O	X	X
After replacing surrounded O to X

	X	X	X	X
	X	X	X	X
	O	X	X	X
	X	O	X	X
*/
var rLen = len(m)
var cLen = len(m[0])

func display(b [][]byte) {
	for i := 0; i < rLen; i++ {
		fmt.Println("")
		for j := 0; j < cLen; j++ {
			fmt.Print("\t", string(b[i][j]))
		}
	}
	fmt.Println()
}

func main() {
	display(m)
	fmt.Println("After replacing surrounded O to X")
	convertSurrounded_O_To_X(m)
	display(m)
}

func convertSurrounded_O_To_X(m [][]byte) {
	if rLen <= 2 || cLen <= 2 {
		return
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			// if last row or last column has 'O' change to 'A'
			if m[i][j] == 'O' && (i == 0 || i == rLen-1 || j == 0 || j == cLen-1) {
				m[i][j] = 'A'
			}
		}
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			// now change surrounded 'O' to X
			if m[i][j] == 'O' {
				m[i][j] = 'X'
			}

			// Then revert back 'A' to X
			if m[i][j] == 'A' {
				m[i][j] = 'O'
			}
		}

	}
}
