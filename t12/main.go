package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func removeDuplicates(a []string) []string {
	set := make(map[string]struct{})

	var unique []string
	for _, x := range a {
		_, ok := set[x]
		if !ok {
			set[x] = struct{}{}
			unique = append(unique, x)
		}
	}
	return unique
}

func main() {
	// Я ничего не понял из условия, но раз просят множество, то давайте удалим дубликаты
	fmt.Println(removeDuplicates([]string{"cat", "cat", "dog", "cat", "tree"}))
}
