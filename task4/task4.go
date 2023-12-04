package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func worker(ctx context.Context, name int, ch chan interface{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker #%d shutdown\n", name+1)
			wg.Done()
			return
		case value, ok := <-ch:
			if !ok {
				break
			}
			fmt.Printf("Worker #%d, value: %v\n", name+1, value)
		}
	}
}

func main() {
	array := []interface{}{2, "hi", 0.33, true}
	ch := make(chan interface{})

	fmt.Println("How many workers?")
	var num int
	_, err := fmt.Fscan(os.Stdin, &num)
	if err != nil {
		fmt.Println("Value is not a number")
		return
	}

	wg := sync.WaitGroup{}

	//Создаём контекст который будет завершён при нажатии Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for i := 0; i < num; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	for {
		select {
		case <-ctx.Done():
			close(ch)
			//ждём все горутины
			wg.Wait()
			return
		default:
			//Пишем в канал, пока контекст не завершился
			ch <- array[rand.Intn(len(array))]
		}
	}
}
