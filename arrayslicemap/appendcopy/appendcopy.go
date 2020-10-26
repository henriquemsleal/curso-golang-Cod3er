package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)
	// não é possível
	// 1o argumento para append precisa ser um slice
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copia para slice2 conteúdo de slice1
	// a fonte de dados não pode ser um array
	// copy não aumenta quantidade de elementos
	fmt.Println(slice2)
}
