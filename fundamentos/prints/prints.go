package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	//não consegue concatenar string e número
	x := 3.141516
	//exemplo fmt.Println("O valor de x é " + x)
	//uma forma é converter em outra var
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	//outra forma é com vírgula
	fmt.Println("O valor de x é", x)
	//outra com Printf
	fmt.Printf("O valor de x é %.2f", x) //%f tipo float e %.2f float com 2 casas decimais

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%d %.4f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
