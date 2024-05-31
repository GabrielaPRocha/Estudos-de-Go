package main

import "fmt"

func main() {
	//	var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[13249080890] = "Pedro"
	aprovados[98765432109] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[98765432109])
	delete(aprovados, 98765432109)
	fmt.Println(aprovados[98765432109])
}
