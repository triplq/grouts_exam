package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		m.Lock()
		go func() {
			counter++
			m.Unlock()
		}()
	}
	fmt.Println(counter)
}
