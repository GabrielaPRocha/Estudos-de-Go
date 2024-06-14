package main

import "fmt"

/*- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
type pessoa struct {
	Nome      string
	Sobrenome string
	sabores   string
}

func main() {
	pessoa1 := pessoa{"Kara", "Zoerl", "Baunilha"}
	pessoa2 := pessoa{"Felicity", "Smoke", "Morango"}
	for _, v := range pessoa1.Sobrenome {
		fmt.Println("\t", v)
	}
	for _, v := range pessoa2.Sobrenome {
		fmt.Println("\t", v)
		fmt.Println(pessoa2.Sobrenome)
	}
}
