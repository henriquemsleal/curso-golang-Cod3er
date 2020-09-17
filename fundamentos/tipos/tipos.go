package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64 // abribui à variável o valor máximo do tipo int64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	var y = 49.99 // teste de tipo
	// por padrão o tipo de um literal flutuante é float64
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float 64
	fmt.Println("O tipo do literal y é", reflect.TypeOf(y))         // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo)) // tipo de bo
	fmt.Println(!bo)                                  // negação

	// string
	s1 := "Olá! Meu nome é Carlos"
	fmt.Println(s1)
	fmt.Println("O tamanho da string é", len(s1)) // largura da string

	// string com multiplas linhas
	s2 :=
		`Olá!
Estou aprendendo
Go Lang`
	fmt.Println(s2)
	fmt.Println("O tamanho de s2 é", len(s2))

	// char?
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
	// não tem tipo char em go

}
