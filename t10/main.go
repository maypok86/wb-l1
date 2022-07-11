package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func testTemperatures(temperatures ...float64) {
	m := make(map[int][]float64)
	for _, temperature := range temperatures {
		i := int(temperature) / 10 * 10
		fmt.Println(int(temperature), i)
		m[i] = append(m[i], temperature)
	}
	fmt.Println(m)
}

func main() {
	testTemperatures(-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5)
	fmt.Println()
	// Есть проблемы с группой вокруг нуля она получается (-10, 10)
	// И как нужно разбивать на группы в этом случае я не понимаю
	testTemperatures(-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -1., 9.)
}
