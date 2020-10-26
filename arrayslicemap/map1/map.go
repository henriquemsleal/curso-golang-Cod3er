package main

import "fmt"

func main() {
	/*var aprovados map[int]string
	cria estrutura do tipo map cuja chave é inteiro e valor string
	vai dar erro, porque mapas devem ser inicializados
	*/
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[45698712355] = "Ana"
	aprovados[14725836988] = "Manoel"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[98765432100]) // imprime o valor da chave
	delete(aprovados, 98765432100)
	fmt.Println(aprovados[98765432100]) // imprime vazio, porque deletou

	fmt.Println(aprovados)
}
