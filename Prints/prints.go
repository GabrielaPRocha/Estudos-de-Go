package main

import "fmt"

func main() {
	//O print imprime tudo na mesma linha
	fmt.Print("primeira")
	fmt.Print("segunda")
	//O println imprime com quebra de linha
	fmt.Println("primeira linha")
	fmt.Println("segunda linha")
	//O printf é o print formatado, você tem mais poderes para alterar, ex tirar casa decimal etc
	x := 1.2345678
	fmt.Printf("O valor de x é %.2f", x)
}
