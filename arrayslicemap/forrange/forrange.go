package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // notação: é array, compilador conta os elementos

	for i, numero := range numeros { /*
			com range retorna dois elementos. para i o índice, para numero o elemento do array
			para cada índice e elemento imprima
		*/
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // ignora o índice
		fmt.Println(num)
	}
}
