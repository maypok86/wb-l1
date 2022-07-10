package main

import "fmt"

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func testReverseString(str string) {
	fmt.Printf("%s - %s\n", str, ReverseString(str))
}

func main() {
	testReverseString("главрыба")
}
