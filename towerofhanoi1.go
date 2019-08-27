package main

import (
	"fmt"
)

func main() {
	// Source Peg     : A
	// Destination Peg: C
	// Aux Peg        : B
	tHanoi1(3, "A", "C", "B")
}

func tHanoi1(n int, sourcePeg, destinationPeg, auxPeg string) {
	if n == 1 {
		fmt.Printf("\n %s -> %s", sourcePeg, destinationPeg)
	} else {

		// Move n-1 from source to Aux
		tHanoi1(n-1, sourcePeg, auxPeg, destinationPeg)
		//Move last item from source to destination
		fmt.Printf("\n %s -> %s", sourcePeg, destinationPeg)
		// Move n-1 from aux to destination
		tHanoi1(n-1, auxPeg, destinationPeg, sourcePeg)
	}
}

/*
 A -> C
 A -> B
 C -> B
 A -> C
 B -> A
 B -> C
 A -> C
*/
