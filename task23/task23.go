package main

import "fmt"

// Удалить i-ый элемент из слайса.

// Изменяет исходный срез
func Remove1(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Не изменяет исходный срез
func Remove2(slice []int, index int) []int {
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	return append(result, slice[index+1:]...)
}

// Изменяет исходный срез
func Remove3(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 2}
	fmt.Println(Remove2(slice, 5))
	fmt.Println(slice)
}
