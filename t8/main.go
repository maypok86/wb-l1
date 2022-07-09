package main

import "fmt"

func setBit(number int64, position int, value int) int64 {
	if value == 1 {
		return number | (1 << position)
	}
	return number &^ (1 << position)
}

func main() {
	fmt.Println(setBit(100, 5, 1))
}
