package main

import "fmt"

//Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

func main() {
	mapa := map[string][]string{
		"Pitty": []string{
			"Priscila Novas Leone", "Bahia", "cantora",
		},
		"Halsey": []string{
			"Ashley", "nova jersey", "cantora",
		},
		"Fernanda Montenegro": []string{
			"Arlette", "Rio de Janeiro", "Atriz",
		},
	}
	//acrescentar alguem:
	mapa["Tais Araujo"] = []string{"Tais Bianca", "Rio de Janeiro", "atriz"}
	//deletar alguem:
	delete(mapa, "Pitty")

	//mostrat o mapa
	for key, value := range mapa {
		fmt.Println(key)
		for indice, hobby := range value {
			fmt.Println("\t", indice, " - ", hobby)
		}
	}
}
