package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	numbers := []int64{2, 4, 6, 7, 8, 10}

	var wg sync.WaitGroup

	var sum int64
	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(number int64) {
			defer wg.Done()

			atomic.AddInt64(&sum, number*number)
		}(number)
	}
	wg.Wait()

	fmt.Println(sum)
}
