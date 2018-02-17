package main

import "fmt"

func main(){
	// any type value can be assigned (because all types implement the empty interface)
	var whatever [3]interface{}
	whatever[0] = 100
	whatever[1] = "Nandeshwar"
	whatever[2] = true

	//To use these values, we must "assert" what type they are.
	var x int
	x = whatever[0].(int) + 100
	fmt.Println(x)
	//If you need to verify the assert works use:

	v, ok := whatever[0].(int)
	if !ok {
		fmt.Println("Type convertion error")
	}
	fmt.Println(v)

    //Note: assert is only for interfaces, use type conversion when casting type:

	z := int64(x)
	fmt.Println(z)
}
