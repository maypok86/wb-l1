package main

import (
	"fmt"
	"strings"
)

type Reverser interface {
	ReverseString(string) string
}

type stringReverser struct{}

func (sr stringReverser) ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type wordsReverser struct{}

func (wr wordsReverser) ReverseWords(str string) string {
	words := strings.Fields(str)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

type wordsReverserAdapter struct {
	wr wordsReverser
}

func (wra wordsReverserAdapter) ReverseString(str string) string {
	return wra.wr.ReverseWords(str)
}

func testAdapter(reverser Reverser, str string) {
	fmt.Printf("'%s' - '%s'\n", str, reverser.ReverseString(str))
}

func main() {
	testAdapter(stringReverser{}, "snow dog sun")
	testAdapter(wordsReverserAdapter{}, "snow dog sun")
}
