package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11000.00,
		"Gabriela Silva": 15000.00,
		"Pedro Júnior":   2000.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.00 // vai adicionar. se existisse, alteraria o valor
	delete(funcsESalarios, "Inexistente")    // tenta deletar, mas não dá erro

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
