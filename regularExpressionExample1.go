package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("p([a-z]+)ch")
	foundItems := r.FindAllString("hello peach pichu peach nandepeachshwar", -1)
	fmt.Println(foundItems)

	/*
		output:
		[peach pich peach peach]
	*/

	r, _ = regexp.Compile("([a-z|0-9|_]+)=([a-z|0-9|_]+)")
	//r, _ = regexp.Compile("(\\w+)=(\\w+)")
	foundItems = r.FindAllString("This is my id=1 and this is my shoe_size=10", 4)
	fmt.Println(foundItems)
	/*
		output:
		[id=1 shoe_size=10]
	*/
}
