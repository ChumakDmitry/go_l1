package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func SwapValue(a, b int) (int, int) {
	return b, a
}

func main() {
	a := 10
	b := 20
	fmt.Printf("a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("Swap #1: a = %d, b = %d\n", a, b)

	a, b = SwapValue(a, b)
	fmt.Printf("Swap #2: a = %d, b = %d", a, b)
}
