package main

import (
	"fmt"
)

func main() {
	printStairsPath(4, "")
}

// current stair (eg 4)
// how we can reach to stair 1
/*
bash-3.2$ go run main.go

 1111  - 1 jump

 112

 121

 13

 211

 22

 31
*/
func printStairsPath(currentPosInStair int, path string) {

	if currentPosInStair < 0 {
		return
	}

	if currentPosInStair == 0 {
		fmt.Println("\n", path)
		return
	}

	printStairsPath(currentPosInStair-1, path+"1")
	printStairsPath(currentPosInStair-2, path+"2")
	printStairsPath(currentPosInStair-3, path+"3")

}
