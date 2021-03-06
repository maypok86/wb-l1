package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
*/

/*
Проблемы две:

1. Глобальная переменная justString
Для исправления достаточно объявить её внутри функции
P.S. в общем случае глобальные переменные - не всегда плохо, но здесь они не нужны

2. Эта проблема, скорее всего, и подразумевалась
string в Go - слайс байтов, а символы utf-8 кодируются несколькими байтами, так как одного байта не хватает для представления символа unicode
Значит v[:100] сломается уже на кириллице. Выход: привести v к слайсу рун

func someFunc() string {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	return string(runes[:100])
}
*/

var justString string

func createHugeString(_ int) string {
	return "lalallaa"
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
