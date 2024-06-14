package main

import (
	"fmt"
)

/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
  - "Nome", "Sobrenome", "Hobby favorito"

- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
func main() {
	ss := [][]string{
		[]string{"Gabriela", "Pereira", "Ouvir musica"},
		[]string{"Oliver", "Queen", "arco e Flecha"},
		[]string{"Barry", "Allen", "Correr"},
	}

	// estou indo dentro de cada dentro do x, e mostrando todos os itens
	for _, x := range ss {
		fmt.Println(x[0])
		for _, item := range x {
			fmt.Println("\t", item)
		}
	}
}
