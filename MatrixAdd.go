package main


import "fmt"

func main() {
	num1 := [][] int { {1, 2},
	                   {3, 4}}

	num2 := [][] int { {5, 6},
	                   {7, 8}}

	num3 :=make([][] int, 2)
	for i := 0; i < 2; i++ {
		num3[i] = make([] int, 2)
		for j := 0; j < 2; j++ {
			num3[i][j] = num1[i][j] + num2[i][j]
		}
	}

	for i := 0; i < 2; i++ {
		fmt.Println("\n")

		for j :=0; j < 2; j++ {
			fmt.Printf("%d ", num3[i][j])
		}
	}
}
