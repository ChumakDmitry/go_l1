package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func CheckUniq(str string) bool {
	m := make(map[rune]bool)
	str = strings.ToLower(str)
	for _, sym := range str {
		if m[sym] {
			return false
		}
		m[sym] = true
	}

	return true
}

func main() {
	str := "123abcd☺☺"
	fmt.Println(CheckUniq(str))
}
