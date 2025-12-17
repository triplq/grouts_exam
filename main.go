package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()

	return ch
}

func worker(jobs <-chan int, results chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-jobs:
				if !ok {
					return
				}
				results <- val * val
			case <-time.After(3 * time.Second):
				fmt.Println(errors.New("times out"))
			}
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	Workers := 5
	jobs := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	results := make(chan int, 5)

	for range Workers {
		worker(jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
