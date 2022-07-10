package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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

type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func QuicksortInts(slice []int) {
	sort.Ints(slice)
	fmt.Println(slice)
}

func QuicksortIntSlice(is IntSlice) {
	sort.Sort(is)
	fmt.Println(is)
}

func partition(slice []int, left, right int) int {
	middle := slice[left+(right-left)/2]
	i := left
	j := right
	for i <= j {
		for slice[i] < middle {
			i++
		}
		for slice[j] > middle {
			j--
		}
		if i >= j {
			break
		}
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	return j
}

func quicksort(slice []int, left, right int) {
	if left < right {
		index := partition(slice, left, right)
		quicksort(slice, left, index)
		quicksort(slice, index+1, right)
	}
}

func QuicksortOwn(slice []int) {
	quicksort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func main() {
	slice1 := generateRandomSlice(30)
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	QuicksortIntSlice(slice1)
	slice3 := make([]int, len(slice2))
	copy(slice3, slice2)
	QuicksortIntSlice(slice2)
	QuicksortOwn(slice3)
}
