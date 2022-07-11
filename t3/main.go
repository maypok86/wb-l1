package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func main() {
	numbers := []int64{2, 4, 6, 7, 8, 10}

	var wg sync.WaitGroup

	var sum int64
	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(number int64) {
			defer wg.Done()

			// избегаем data race
			// можно ещё использовать mutex, но для атомарных операций с числами
			// правильнее использовать atomic
			atomic.AddInt64(&sum, number*number)
		}(number)
	}
	wg.Wait()

	fmt.Println(sum)
}
