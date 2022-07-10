package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Fields(str)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func testReverseWords(str string) {
	fmt.Printf("'%s' - '%s'\n", str, ReverseWords(str))
}

func main() {
	testReverseWords("snow dog sun")
}
