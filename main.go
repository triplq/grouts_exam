package main

import (
	"sync"
)

var wg sync.WaitGroup

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

func main() {

}
