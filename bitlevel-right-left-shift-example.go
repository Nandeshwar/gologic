package main

import (
	"fmt"
)

func main() {

	s1 := [5]byte{0x00, 0x06, 0x6F, 0x9A, 0x0F}
	a0 := int32(s1[0])
	a1 := int32(s1[1])
	a2 := int32(s1[2])
	fmt.Println(a0)
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Printf("\n%032b", a1)
	a1 = a1 << 4
	fmt.Printf("\n%032b", a1)

	fmt.Println("----------\n")

	fmt.Printf("\n%032b", a2)
	a2 = a2 << 4
	fmt.Printf("\n%032b", a2)
}

/*
output:
 0
6
111

00000000000000000000000000000110
00000000000000000000000001100000----------


00000000000000000000000001101111
00000000000000000000011011110000
*/
