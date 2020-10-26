package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  15456,
			"Gustavo Pereira": 8456,
		},
		"J": {
			"João Carlos": 11325,
			"José Filho":  4000,
		},
		"P": {
			"Pedro Silva": 5000,
		},
	}

	for nome, salario := range funcsPorLetra {
		fmt.Println(nome, salario)
	}

	delete(funcsPorLetra, "P")

	for nome, salario := range funcsPorLetra {
		fmt.Println(nome, salario)
	}
}
