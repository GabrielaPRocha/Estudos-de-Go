package main

import "fmt"

/* Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.*/

func main() {
	x := struct {
		modelo     string
		cor        []string
		acessorios map[string]string
	}{
		modelo: "Ecosport",
		cor:    []string{"Preto", "Azul", "Vermelho"},
		acessorios: map[string]string{
			"Radio":              "multimidia",
			"direção hidraulica": "direção eletrica",
		},
	}
	fmt.Println(x)

}
