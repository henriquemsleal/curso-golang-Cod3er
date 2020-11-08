package main

import "fmt"

var total, med float64

func media(numeros ...float64) float64 { // ... parâmetros variáveis
	// dentro da função, os parâmetros serão tratados como array
	for _, num := range numeros { // para percorrer os parâmetros, sem o índice
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	med = media(7.5, 10, 5, 6.8, 8.9, 9)
	fmt.Println("Total das notas:", total)
	// fmt.Printf("Média: %.2f", media(7.5, 10, 5, 6.8, 8.9, 9)) // da aula
	fmt.Printf("Média: %.2f", med) // da aula
}
