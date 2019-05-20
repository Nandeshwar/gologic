package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	go testContext(ctx)
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
output: context will be closed in 2 second

Context not closed yet. Doing noting
This is testContext:  1
context deadline exceeded
There is an error:  2

*/
