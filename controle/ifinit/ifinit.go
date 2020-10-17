package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	// pacote rand de random, número aleatório, nano segundo da data atual
	// passar como fonte para gerar número aleatório
	r := rand.New(s)  // gera número novo
	return r.Intn(10) // número até 10
}

func main() {
	// bloco de inicialização com variável interna chama método e recebe número aleatório até 10
	if i := numeroAleatorio(); i > 5 { // bloco de inicialização tb suportado pelo switch, depois condição
		fmt.Printf("%v\nGanhou!!!", i)
	} else {
		fmt.Printf("%v\nPerdeu!", i)
	}

}
