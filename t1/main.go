package main

import "fmt"

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
	Name string
}

func main() {
	h := Human{
		Name:    "Vasya",
		Surname: "Pupkin",
	}
	a := Action{
		Human: h,
		Name:  "Petya",
	}
	h.PrintInfo()
	a.PrintInfo()
	fmt.Println(a.Name)
}
