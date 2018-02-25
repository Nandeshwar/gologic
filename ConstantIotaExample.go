package main

import "fmt"

type Season uint8

// Assign 0 to 3 - Spring, Summer, Autumn, Winter
const (
	Spring = Season(iota)
	Summer
	Autumn
	Winter
)


func main() {
	// This will generate maximum number of given type
	const MaxUint = ^uint(0)
	const MaxUint64 = ^uint64(0)

	fmt.Println(MaxUint)
	fmt.Println(MaxUint64)

	fmt.Println(Spring)
	fmt.Println(Winter)
}

// This Allows to print string (spring, summer, autumn, winter, other) instead of integer 0-3
func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	default:
		return "unknown"

	}
}

/*
output:

18446744073709551615
18446744073709551615
spring
winter

 */