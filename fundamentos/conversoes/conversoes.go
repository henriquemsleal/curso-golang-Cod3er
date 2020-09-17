package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // tipo float64
	y := 2   // tipo inteiro

	// se tentar fmt.Println(x / y) vai dar erro

	// precisa converter
	fmt.Println(x / float64(y))

	nota := 6.9            // tipo float64
	notaFinal := int(nota) // converte p inteiro (despreza decimais)
	fmt.Println("A nota final é", notaFinal)

	// cuidado...

	// concatenar inteiro precisa converter
	fmt.Println("Teste " + string(123)) // código de '{' na tabela asc
	fmt.Println("Teste " + string(97))  // código de 'a' na tabela asc

	// maneira correta int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") // em go a função retorna 2 valores, sendo a 2a um erro
	fmt.Println(num - 122)        // int - int = 1

}
