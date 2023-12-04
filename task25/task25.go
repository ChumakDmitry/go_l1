package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	x := 5
	Sleep(5)
	fmt.Println(x)
}
