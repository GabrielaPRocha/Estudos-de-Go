package main

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTV50, comprarTV32, comprarSorvete
}

/*
func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t,tv32 %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}*/
