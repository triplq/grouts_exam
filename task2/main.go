package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/triplq/goruts_exam/utils"
)

func reader(ch <-chan int, wg *sync.WaitGroup) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	defer wg.Done()
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("No vals")
				return
			}
			fmt.Print(val, " ")
			timer.Reset(3 * time.Second)
		case <-timer.C:
			fmt.Println("no tasks 4long...")
		default:
			fmt.Println("\ndefault")
			time.Sleep(3 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := utils.Generator_by_range(100)

	go reader(ch, &wg)

	wg.Wait()
}
