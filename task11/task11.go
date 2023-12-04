package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func CreateMap(arr []int) map[int]bool {
	//создание множества, у которого ключи это элементы массива,
	//а значение - true
	set := make(map[int]bool)
	for _, value := range arr {
		set[value] = true
	}

	return set
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7}
	set2 := []int{5, 6, 7, 3, 8, 12}

	map1 := CreateMap(set1)
	map2 := CreateMap(set2)

	result := make(map[int]bool)

	//Делаем перебор по ключам одного множества и смотрим, есть
	// ли во втором множистве такой же ключ, т.е. по значению ключа - true
	for key := range map1 {
		if map2[key] {
			result[key] = true
		}
	}

	for key := range result {
		fmt.Printf("%d, ", key)
	}
}
