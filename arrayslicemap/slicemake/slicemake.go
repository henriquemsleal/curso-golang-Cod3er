package main

import "fmt"

func main() {
	s := make([]int, 10)
	// cria slice com método make e, internamente, ele já cria o array associado.
	fmt.Println(s)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	// nova versão na variável, com 10 elementos e array interno com 20 posições
	s[9] = 13
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	// não vai dar erro. o slice cresce automaticamente e referencia arrays diferentes.
	fmt.Println(s, len(s), cap(s))

}
