package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		// 1a parte inicializa variável; 2a - condição; 3a - incremento
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	// if em for
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é impar")
		}
	}

	// laço infinito com for
	/*
		for {
			fmt.Println("Para sempre...")
			time.Sleep(time.Second)
		}*/
}
