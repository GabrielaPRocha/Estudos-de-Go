package main

/*- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.*/
import (
	"fmt"
)

func main() {

	mapa := map[string][]string{
		"Kara Zoerl": []string{
			"voar", "laser nos olhos", "super força",
		},
		"Oliver Queen": []string{
			"Voltar dos mortos", "inteligencia", "planejamento",
		},
		"Natasha Romanoff": []string{
			"Espiã", "disfarce", "lealdade",
		},
	}

	for key, value := range mapa {
		fmt.Println(key)
		for indice, hobby := range value {
			fmt.Println("\t", indice, " - ", hobby)
		}
	}

}
