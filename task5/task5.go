package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

// Реализация через context.WithDeadline
func ReadFunc(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("End reading")
			fmt.Printf("Time now = %v", time.Now())
			wg.Done()
			return
		case value := <-ch:
			fmt.Printf("value = %d\n", value)
		}
	}
}

func main() {
	var timer int
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Введите кол-во секунд работы программы")
	fmt.Fscan(os.Stdin, &timer)

	if timer == 0 {
		return
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(timer)*time.Second))
	defer cancel()

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("Time now = %v\n", time.Now())
	go ReadFunc(ctx, ch, &wg)

	for {
		select {
		case <-ctx.Done():
			close(ch)
			wg.Wait()
			return
		default:
			ch <- numbers[rand.Intn(len(numbers))]
			time.Sleep(10 * time.Millisecond)
		}
	}
}
