package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func ReverseWord1(str string) string {
	// Разделяем строку по пробелам
	stringArray := strings.Split(str, " ")
	j := len(stringArray) - 1

	// Меняем местами крайние элементы
	for i := 0; i <= j; i++ {
		stringArray[i], stringArray[j] = stringArray[j], stringArray[i]
		j--
	}

	// Конкатинируем строки и возвращаем
	return strings.Join(stringArray, " ")
}

func ReverseWord2(str string) string {
	// Разделяем строку по пробелам
	stringArray := strings.Fields(str)

	// Объявляем структуру Builder, которая используется для конкатинации строк
	// В нулевом значении уже готова к использованию
	// Является лучшей практикой для конкатинации строк т.к. минимизирует затрачиваемую память
	var b strings.Builder

	for i := len(stringArray) - 1; i >= 0; i-- {
		b.WriteString(stringArray[i])
		b.WriteString(" ")
	}

	return b.String()
}

func main() {
	s := "snow dog sun"
	fmt.Println(ReverseWord1(s))
	fmt.Println(ReverseWord2(s))
}
