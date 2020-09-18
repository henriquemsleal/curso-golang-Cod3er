package main

import "fmt"

func main() {
	i := 1
	fmt.Println("i =", i)

	var p *int = nil // cria variável do tipo ponteiro com valor nulo

	p = &i // pegando o endereço da variável i
	*p++   // desreferenciando (pegando o valor '*p') e incrementando i pelo ponteiro
	fmt.Println("i =", i)
	i++ // incrementa novamente
	fmt.Println("i =", i)

	// Go não tem aritimética de ponteiros
	// exemplo: 'p++'

	fmt.Println("p = endereço de i:", p, "\n*p = valor de i:", *p, "\ni =", i, "\n&i = endereço de i:", &i)
}
