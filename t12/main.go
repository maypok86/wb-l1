package main

import "fmt"

func unique(a []string) []string {
	set := make(map[string]struct{}, len(a))

	u := make([]string, 0, len(a))
	for _, x := range a {
		_, ok := set[x]
		if !ok {
			set[x] = struct{}{}
			u = append(u, x)
		}
	}
	return u
}

func main() {
	// Я ничего не понял из условия, но раз просят множество, то давайте удалим дубликаты
	fmt.Println(unique([]string{"cat", "cat", "dog", "cat", "tree"}))
}
