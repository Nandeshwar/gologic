package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World")
	selectChannelExample()

	go func() {

		defer func() {
			if arg := recover(); arg != nil {
				fmt.Println(arg)
			}
		}()

		b, _ := strconv.Atoi("0")
		s := 1 / b
		fmt.Println("this is go routine1")
		fmt.Println(s)

	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end of the program")
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
