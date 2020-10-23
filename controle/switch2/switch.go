package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switch true 1a expressão verdadeira
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
