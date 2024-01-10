package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	selectChannelExample()

	m := make(map[string]string)
	k := fmt.Sprintf("%d", 60)
	m[k] = "nks"
	k2 := fmt.Sprintf("%s", "60")
	v, _ := m[k2]
	fmt.Println(v)
}

func selectChannelExample() {
	fmt.Println("Hello World")

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	close(ch)

	for i := 0; i < 10; i++ {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("channel closed")
				ch = nil
			}

		default:
			fmt.Println("No item yet")
		}
	}
}
