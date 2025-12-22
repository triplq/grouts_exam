package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/triplq/goruts_exam/utils"
)

func worker(sem chan struct{}, jobs <-chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	for {
		select {
		case val, ok := <-jobs:
			if !ok {
				return
			}
			sem <- struct{}{}
			results <- val * val
			<-sem
			timer.Reset(1 * time.Second)
		case <-timer.C:
			fmt.Println(errors.New("times out"))
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	Workers := 5
	wg.Add(Workers)

	sem := make(chan struct{}, 3)

	jobs := utils.Generator_by_slice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	results := make(chan int, 5)

	for range Workers {
		go worker(sem, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
