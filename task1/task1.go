package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

import "fmt"

type Human struct {
	name string
	age  int
}

func NewHuman(name string, age int) *Human {
	return &Human{
		name: name,
		age:  age,
	}
}

func (h *Human) PrintName() {
	fmt.Println(h.name)
}

func (h *Human) PrintAllData() {
	fmt.Printf("Name: %s, age: %d\n", h.name, h.age)
}

type Action struct {
	Human
}

func NewAction(name string, age int) *Action {
	return &Action{
		Human{
			name: name,
			age:  age,
		},
	}
}

func main() {
	h := NewHuman("Ivan", 15)
	h.PrintAllData()
	a := NewAction("Dima", 20)
	a.PrintAllData()
	a.name = "Vova"
	a.PrintAllData()
	fmt.Println(h)
	fmt.Println(a)
}
