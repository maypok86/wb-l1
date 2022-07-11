package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	numbers := []int{2, 4, 6, 7, 8, 10}

	var wg sync.WaitGroup

	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(number int) {
			defer wg.Done()

			fmt.Println(number * number)
		}(number)
	}
	wg.Wait()
}
