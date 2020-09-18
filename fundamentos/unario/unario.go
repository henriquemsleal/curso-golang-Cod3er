package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Println("x e y =", x, y)

	// go incremento e decremento só tem forma pós-fixada
	x++ // x += 1 ou x = x + 1
	y-- // y -= 1 ou y = y - 1

	fmt.Println("incrementa x =", x, "decrementa y =", y)
}
