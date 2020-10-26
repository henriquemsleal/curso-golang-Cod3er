package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3) // aponta para o mesmo array interno criado em s1
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)
	// mesmo array interno é capaz de ser referenciado por vários slices
}
