package main

import (
	"fmt"
	"math/big"
	"os"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func main() {
	fmt.Println("Input first number number")
	var (
		temp   string
		first  = new(big.Int)
		second = new(big.Int)
	)

	_, err := fmt.Fscan(os.Stdin, &temp)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	first.SetString(temp, 10)

	fmt.Println("Input second number number")

	_, err = fmt.Fscan(os.Stdin, &temp)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	second.SetString(temp, 10)

	fmt.Println("Input operation")

	_, err = fmt.Fscan(os.Stdin, &temp)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	switch temp {
	case "+":
		fmt.Printf("Result = %v", first.Add(first, second))
	case "*":
		fmt.Printf("Result = %v", first.Mul(first, second))
	case "/":
		fmt.Printf("Result = %v", first.Div(first, second))
	case "-":
		fmt.Printf("Result = %v", first.Sub(first, second))
	}
}
