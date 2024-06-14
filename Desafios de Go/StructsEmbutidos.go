package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}
type profissao struct {
	pessoa
	profissao string
	idade     int
}

func main() {
	pessoa1 := pessoa{"Joana", "Silva"}
	pessoa2 := profissao{
		pessoa:    pessoa{"Jose", "Oliveira"},
		profissao: "Administrador",
		idade:     25,
	}
	fmt.Println(pessoa1)

	fmt.Println(pessoa2)
}
