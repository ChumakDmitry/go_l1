package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
// с использованием конкурентных вычислений.

func FindEqualsSquareChanel(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	ch := make(chan int, len(numbers))
	var sum int

	go func() {
		for _, value := range numbers {
			ch <- value * value
		}
		close(ch)
	}()

	for value := range ch {
		sum += value
	}

	return sum
}

func FindEqualsSquare(numbers []int) int {
	sum := 0
	wg := sync.WaitGroup{}

	for _, value := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sum += num * num
		}(value)
	}
	wg.Wait()

	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	if len(numbers) == 0 {
		fmt.Println("Array is empty")
		return
	}

	//result := FindEqualsSquareChanel(numbers)
	result := FindEqualsSquare(numbers)

	fmt.Printf("Result = %d", result)
}
