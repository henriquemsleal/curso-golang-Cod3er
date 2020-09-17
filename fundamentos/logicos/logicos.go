package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true) // trabalhou em dois dias
	fmt.Printf("\nSe trabalhou dois dias\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t,",
		tv50, tv32, sorvete, !sorvete) // !sorvete = saldável ou não

	tv50, tv32, sorvete = compras(true, false) // trabalhou no primeiro dia
	fmt.Printf("\nSe trabalhou no primeiro dia\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t,",
		tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(false, true) // trabalhou no segundo dia
	fmt.Printf("\nSe trabalhou no segundo dia\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t,",
		tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(false, false) // não trabalhou nos dois dias
	fmt.Printf("\nSe não trabalhou nos dois dias\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t,",
		tv50, tv32, sorvete, !sorvete)
}
