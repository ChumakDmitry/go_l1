package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	wg := sync.WaitGroup{}
	// 1 способ. Завершение горутины передачей данных в канал
	ch1 := make(chan int)
	go func() {
		wg.Add(1)
		for {
			select {
			case <-ch1:
				fmt.Println("Goroutine 1 stop")
				wg.Done()
				return
			default:
				fmt.Println("Goroutine 1 working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	ch1 <- 1
	time.Sleep(100 * time.Millisecond)

	// 2 способ. Закрытие канала

	go func() {
		wg.Add(1)
		for {
			value, ok := <-ch1
			if !ok {
				fmt.Println("Goroutine 2 stop")
				wg.Done()
				return
			}
			fmt.Printf("Goroutine 2 working. Value = %d\n", value)
		}
	}()

	ch1 <- 3
	ch1 <- 2
	ch1 <- 1
	close(ch1)
	time.Sleep(300 * time.Millisecond)

	// 3 способ. Установка ограничения по времени

	timeout := time.After(1 * time.Second)
	go func() {
		wg.Add(1)
		for {
			select {
			case <-timeout:
				fmt.Println("Goroutine 3 stop")
				wg.Done()
				return
			default:
				fmt.Println("Goroutine 3 working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)

	// 4 способ. Передача контекста
	ctx, stop := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(2)*time.Second))
	defer stop()

	go func() {
		wg.Add(1)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 4 stop")
				wg.Done()
				return
			default:
				fmt.Println("Goroutine 4 working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
