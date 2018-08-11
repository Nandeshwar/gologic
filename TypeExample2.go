package main

import (
	"fmt"
)

type closeFunc func()

type S struct {
}

func(s S) Close() {
	fmt.Println("This is close of S")
	
}


func main() {
	s := S{}
	
	//k := closeFunc(s.Close)
	// The below two line is similar to one line above  
	var k closeFunc
	k = s.Close
	
	k()
}

/*
output
This is close of S
*/
