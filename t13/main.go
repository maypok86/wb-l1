package main

import "fmt"

func main() {
	a := 4
	b := 5

	b = b - a
	a = a + b
	b = a - b
	fmt.Println(a, b)
}
