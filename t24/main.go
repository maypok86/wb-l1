package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) DistanceTo(other Point) float64 {
	distX := p.x - other.x
	distY := p.y - other.y
	return math.Sqrt(distX*distX + distY*distY)
}

func main() {
	p := NewPoint(1.25, 0.34)
	fmt.Println("Distance:", p.DistanceTo(NewPoint(6.89, 7.01)))
}
