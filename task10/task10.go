package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
//градусов. Последовательность в подмножноствах не важна.
//
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	var set = make(map[int][]float32)
	temper := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, value := range temper {
		switch {
		case value < -20:
			set[-20] = append(set[-20], value)
		case value > -20 && value < -10:
			set[-10] = append(set[-10], value)
		case value > -10 && value < 10:
			set[0] = append(set[0], value)
		case value > 10 && value < 20:
			set[10] = append(set[10], value)
		case value > 20 && value < 30:
			set[20] = append(set[20], value)
		case value > 30:
			set[30] = append(set[30], value)
		}
	}

	for index, value := range set {
		fmt.Printf("%d: %.1f\n", index, value)
	}
}
