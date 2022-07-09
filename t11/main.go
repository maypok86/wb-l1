package main

import "fmt"

func intersect(a, b []int) []int {
	aSet := make(map[int]struct{})
	bSet := make(map[int]struct{})

	for _, x := range a {
		aSet[x] = struct{}{}
	}
	for _, x := range b {
		bSet[x] = struct{}{}
	}

	out := make([]int, 0, len(aSet)+len(bSet))
	for x := range aSet {
		_, ok := bSet[x]
		if ok {
			out = append(out, x)
		}
	}
	return out
}

func main() {
	fmt.Println(intersect([]int{1, 2, 3, 4}, []int{1, 2, 5, 6, 1}))
}