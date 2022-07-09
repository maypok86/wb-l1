package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5., 7.}
	m := make(map[int][]float64)
	for _, temperature := range temperatures {
		i := int(temperature) / 10 * 10
		fmt.Println(int(temperature), i)
		m[i] = append(m[i], temperature)
	}
	fmt.Println(m)
}
