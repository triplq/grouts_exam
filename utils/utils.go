package utils

func Generator_by_slice(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()

	return ch
}

func Generator_by_range(num int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range num {
			ch <- i
		}
	}()

	return ch
}
