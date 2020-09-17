package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int // ponteiro de inteiro (vai ser nulo pq nao atribuiu nada)

	//valores zeros das variáves básicas
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)   // exibe o valor das variáveis = 0 0 false  <nil>
	fmt.Printf("\n%v %v %v %q %v", a, b, c, d, e) // q para string exibe com aspas

}
