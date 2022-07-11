package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func IsUnique(str string) bool {
	m := make(map[rune]struct{})
	for _, ch := range strings.ToLower(str) {
		_, ok := m[ch]
		if ok {
			return false
		}
		m[ch] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(IsUnique("abcd"))
	fmt.Println(IsUnique("abCdefAaf"))
	fmt.Println(IsUnique("aabcd"))
	fmt.Println(IsUnique("Россия"))
	fmt.Println(IsUnique("мышь"))
}
