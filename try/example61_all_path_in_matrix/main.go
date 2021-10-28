package main

// https://www.youtube.com/watch?v=46zD5d9y9b4

// find count of all posible paths to reach to destination from source. only left -> right or top to bottom path is allowed
func findAllPossiblePathsCountInMatrix(m [][]int) int {
	return findCountOfPossbilePathInMatrix(m, len(m), len(m[0]))
}

func findCountOfPossbilePathInMatrix(m [][]int, row, column int) int {
	if row == 1 || column == 1 {
		return 1
	}

	return findCountOfPossbilePathInMatrix(m, row-1, column) + findCountOfPossbilePathInMatrix(m, row, column-1)
}
