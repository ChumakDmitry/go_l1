package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func main() {
	data := make(map[string]int)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ind int) {
			defer wg.Done()
			mutex.Lock()
			data[strconv.Itoa(ind)] = ind
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	for index, value := range data {
		fmt.Printf("Index = %s, Value = %d\n", index, value)
	}
}
