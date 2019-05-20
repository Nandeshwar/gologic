package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go testContext(ctx)
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 120)

}

func testContext(ctx context.Context) {

	for i := 1; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println("There is an error: ", i)
			return

		default:
			fmt.Println("Context not closed yet. Doing noting")
		}

		time.Sleep(5 * time.Second)

		fmt.Println("This is testContext: ", i)
	}
}

/*
output when context is not cancelled
Context not closed yet. Doing noting
This is testContext:  1
Context not closed yet. Doing noting
This is testContext:  2
Context not closed yet. Doing noting
This is testContext:  3
Context not closed yet. Doing noting
This is testContext:  4
Context not closed yet. Doing noting
This is testContext:  5
Context not closed yet. Doing noting


output: when context cancelled
Context not closed yet. Doing noting
This is testContext:  1
context canceled
There is an error:  2

 */
