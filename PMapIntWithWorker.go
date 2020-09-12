package main

import (
	"fmt"
	"sync"

	"github.com/logic-building/functional-go/fp"
)

func main() {
	newList := PMapIntWithWorker(func(val int) int {
		return val * val
	}, fp.RangeInt(1, 101), 10)

	fmt.Println(newList)
}

func PMapIntWithWorker(f func(int) int, list []int, worker int) []int {
	if f == nil {
		return []int{}
	}

	chJobs := make(chan int, len(list))
	for _, v := range list {
		chJobs <- v
	}
	close(chJobs)

	chResult := make(chan int, worker/3)

	var wg sync.WaitGroup
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func(chResult chan int, chJobs chan int) {
			defer wg.Done()
			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}
	return newList
}
