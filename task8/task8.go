package main

import (
	"fmt"
	"os"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var number int64
	var i uint8

	fmt.Println("Введите число")
	fmt.Fscan(os.Stdin, &number)
	fmt.Println("Введите номер бита (от 0), который необходимо инвертировать")
	fmt.Fscan(os.Stdin, &i)

	if i > 64 && i < 0 {
		fmt.Println("Столько битов нет в числе")
		return
	}

	fmt.Printf("Введённое десятичное число: %d\nБитовая запись числа: %064b\n", number, number)
	number ^= 1 << i
	fmt.Printf("Полученное десятичное число: %d\nБитовая запись числа: %064b", number, number)
}
