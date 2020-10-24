package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64 // cria com valores zeros do tipo informado
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ { // len para obter o tamanho do array
		total += notas[i]
	}

	media := total / float64(len(notas))          // converter o tamanho de inteiro para float64
	fmt.Println("A média das notas é", media)     // várias casas decimais
	fmt.Printf("A média das notas é %.2f", media) // ponto flutuante com duas casas decimais
}
