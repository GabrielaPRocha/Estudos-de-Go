package main

import "fmt"

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior
*/

type pessoa struct {
	Nome      string
	Sobrenome string
	sabores   []string
}

func main() {
	mapa := make(map[string]pessoa)
	mapa["Oliver"] = pessoa{
		Nome:      "Oliver",
		Sobrenome: "Queen",
		sabores:   []string{"pistache", "morango", "baunilha"}}
	mapa["Barry"] = pessoa{
		Nome:      "Barry",
		Sobrenome: "Allen",
		sabores:   []string{"limão", "uva", "maracuja"}}

	for _, valor := range mapa {
		fmt.Println("Nome:", valor.Nome, "sorvetes:")
		for _, valor := range valor.Sobrenome {
			fmt.Println("-", valor)
		}
	}
	fmt.Println("\n")
}
