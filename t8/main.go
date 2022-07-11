package main

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(number int64, position int, value int) int64 {
	// с помощью 1 << position получаем число с 1 в i бите и остальными битами равными 0
	if value == 1 {
		return number | (1 << position)
	}
	return number &^ (1 << position)
}

func main() {
	fmt.Println(setBit(100, 5, 1))
}
