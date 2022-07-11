package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать конвейер чисел.
Даны два канала:
в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func read(numbers ...int) <-chan int {
	in := make(chan int, len(numbers))
	go func() {
		for _, x := range numbers {
			in <- x
		}
		close(in)
	}()
	return in
}

func double(in <-chan int) <-chan int {
	out := make(chan int, len(in))
	go func() {
		for x := range in {
			out <- 2 * x
		}
		close(out)
	}()
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := make([]int, 10)
	for i := range a {
		a[i] = rand.Int()
	}

	for x := range double(read(a...)) {
		fmt.Println(x)
	}
}
