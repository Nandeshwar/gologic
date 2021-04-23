package main

import "fmt"

// https://www.youtube.com/watch?v=U6-X_QOwPcs&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-
func main() {
	arr1 := [4][5]int{
						{1, 1, 1, 1, 0},
						{1, 1, 0, 1, 0},
						{1, 1, 0, 0, 0},
						{0, 0, 0, 0, 0},
					}

	arr2 := [4][5]int{
						{1, 1, 0, 0, 0},
						{1, 1, 0, 0, 0},
						{0, 0, 1, 0, 0},
						{0, 0, 0, 1, 1},
					}

	fmt.Println(arr1)
	fmt.Println(arr2)

	cnt := 0
	for i, r := range arr1 {
		for j, _ := range r {
			if arr1[i][j] == 1 {
				cnt++
				callBSF(&arr1, i, j)
			}
		}
	} 

	cnt = 0
	fmt.Printf("count = %d\n", cnt)
	for i, r := range arr2 {
		for j, _ := range r {
			if arr2[i][j] == 1 {
				cnt++
				callBSF(&arr2, i, j)
			}
		}
	} 

	fmt.Printf("count = %d\n ", cnt)
	
	
}

func callBSF(arr *[4][5]int, i, j int) {
	if i < 0 || i >= len(arr) || j < 0 || j >= len(arr[i]) || arr[i][j] == 0 {
		return
	}

	arr[i][j] = 0
	callBSF(arr, i+1, j) // up
	callBSF(arr, i-1, j) // down
	callBSF(arr, i, j-1) //left 
	callBSF(arr, i, j+1) // right
}

/*
output
[[1 1 1 1 0] [1 1 0 1 0] [1 1 0 0 0] [0 0 0 0 0]]
[[1 1 0 0 0] [1 1 0 0 0] [0 0 1 0 0] [0 0 0 1 1]]
count = 0
count = 3
*/