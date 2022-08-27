package main

import (
	"fmt"
)

func main() {
	towerOfHenoi(3, "t1", "t3", "t2") // move from t1 to t3 using t2
	fmt.Println("Another way")
	th(3, "t1", "t3", "t2")
}

/*
	move n-1 ie. 1 and 2 from t1 -> t2 using t3
	move 3 to destination
	move n-1 i.e. 1 and 2 from t2 -> t3 using t1
*/
func towerOfHenoi(n int, t1, t3, t2 string) {
	if n == 0 {
		return
	}

	towerOfHenoi(n-1, t1, t2, t3)
	fmt.Printf("\nmove disk=%d from %s to %s", n, t1, t3)
	towerOfHenoi(n-1, t2, t3, t1)
}

/*
output:
move disk=1 from t1 to t3
move disk=2 from t1 to t2
move disk=1 from t3 to t2
move disk=3 from t1 to t3
move disk=1 from t2 to t1
move disk=2 from t2 to t3
move disk=1 from t1 to t3
*/

func th(n int, source, dest, tmp string) {
	if n == 0 {
		return
	}
	th(n-1, source, tmp, dest)
	fmt.Printf("\n move disk=%d from %s to %s", n, source, dest)
	th(n-1, tmp, dest, source)
}
