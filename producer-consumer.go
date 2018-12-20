package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	producerChannel := make(chan int, 5)
	receiverChannel := make(chan int, 7)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			producerChannel <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {

			k := <-producerChannel
			receiverChannel <- k
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-receiverChannel:
				fmt.Println(v)

			}

		}
	}()

	wg.Wait()
}
