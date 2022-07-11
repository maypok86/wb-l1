package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Удалить i-ый элемент из слайса.
*/

func testDeleteByIndex(slice []int, index int, deleteByIndex func([]int, int) []int) {
	fmt.Println(deleteByIndex(slice, index))
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(numbers))
	fmt.Println(numbers)
	fmt.Printf("Index: %d\n", index)

	testDeleteByIndex(copiedNumbers, index, func(slice []int, index int) []int {
		return append(slice[:index], slice[index+1:]...)
	})

	testDeleteByIndex(numbers, index, func(slice []int, index int) []int {
		for i := index; i < len(slice)-1; i++ {
			slice[i] = slice[i+1]
		}
		return slice[:len(slice)-1]
	})
}
