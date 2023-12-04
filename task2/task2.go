package main

//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func FindSquaresWG(numbers []int) []int {
	squares := make([]int, len(numbers))
	wg := sync.WaitGroup{}

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		// Здесь необходимо явно передать индекс в горутину
		// чтобы использовалось правильное значение
		go func(index int) {
			defer wg.Done()

			square := numbers[index] * numbers[index]
			squares[index] = square
		}(i)
	}
	wg.Wait()

	return squares
}

func main() {
	var numbers = []int{2, 4, 6, 8, 10}
	if len(numbers) == 0 {
		fmt.Println("Error, array is empty")
		return
	}

	squares := FindSquaresWG(numbers)
	if squares == nil {
		return
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Number: %d, Square = %d\n", numbers[i], squares[i])
	}
}
