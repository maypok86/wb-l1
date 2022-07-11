package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

/*
Да кто эти ваши встроенные методы языка?..
Это штуки из пакета sort или нужно прям алгос написать?
*/

func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(25)
	}
	return slice
}

func BinarySearch1(slice []int, x int) int {
	return sort.Search(len(slice), func(i int) bool {
		return slice[i] >= x
	})
}

func BinarySearch2(slice []int, value int) int {
	left := -1
	right := len(slice)
	for left < right-1 {
		middle := left + (right-left)/2
		if slice[middle] < value {
			left = middle
		} else {
			right = middle
		}
	}
	return right
}

func main() {
	slice := generateRandomSlice(10)
	sort.Ints(slice)
	fmt.Println(slice)
	x := slice[rand.Intn(len(slice))]
	fmt.Println(x, BinarySearch1(slice, x))
	fmt.Println(x, BinarySearch2(slice, x))
}
