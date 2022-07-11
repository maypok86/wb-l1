package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name    string
	Surname string
}

func (h Human) PrintInfo() {
	fmt.Println("Name:", h.Name)
	fmt.Println("Surname:", h.Surname)
}

type Action struct {
	Human
}

func main() {
	h := Human{
		Name:    "Vasya",
		Surname: "Pupkin",
	}
	a := Action{
		Human: h,
	}
	h.PrintInfo()
	a.PrintInfo()
	a.Human.PrintInfo()
	fmt.Println(a.Name)
}
