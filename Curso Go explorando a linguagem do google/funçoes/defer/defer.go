package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	if numero >= 5000 {
		fmt.Println("Valor Alto ...")
		return 5000
	} else {
		fmt.Println("Valor Baixo...")
		return numero
	}
}
func main() {
	fmt.Println(obterValorAprovado(6000))
}
