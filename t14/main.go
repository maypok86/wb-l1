package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func GetValueType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan struct{}:
		fmt.Println("chan")
	default:
		panic("unknown type")
	}
}

func main() {
	GetValueType(10)
	GetValueType("vasya")
	GetValueType(true)
	GetValueType(make(chan struct{}))
}
