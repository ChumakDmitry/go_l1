package main

import (
	"fmt"
)

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func typeOf(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "i dont now"
	}
}

func ofType(i interface{}) {
	fmt.Printf("%T\n", i)
}

func main() {
	var item interface{}

	item = 3
	ofType(item)
	fmt.Println(typeOf(item))

	item = true
	ofType(item)
	fmt.Println(typeOf(item))

	item = make(chan int)
	ofType(item)
	fmt.Println(typeOf(item))

	item = "int"
	ofType(item)
	fmt.Println(typeOf(item))
}
