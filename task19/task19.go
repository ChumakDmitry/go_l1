package main

import (
	"fmt"
	"unicode/utf8"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

// Вариант 1. Переводим в срез рун и иеняем местами крайние элементы
// смещаясь в центр
func Upheaval1(str string) string {
	runeArray := []rune(str)

	j := len(runeArray) - 1
	for i := 0; i <= j; i++ {
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
		j--
	}

	return string(runeArray)
}

// Вариант 2. Инициализируем массив рун и проходимся по строке range,
// записывая элементы в конец, постемпенно смещаясь к началу.
func Upheaval2(str string) string {
	strLen := utf8.RuneCountInString(str)
	result := make([]rune, strLen)
	i := 1

	for _, c := range str {
		result[strLen-i] = c
		i++
	}

	return string(result)
}

func main() {
	s := "☺ ☻ ☹ シ "
	//s := "1234567890"
	fmt.Println(s)

	result := Upheaval1(s)
	fmt.Println(result)

	result = Upheaval2(s)
	fmt.Println(result)
}
