package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 { // no go mesmo com uma sequencia de codigo é necessario bloco, não é necessario parenteses
		fmt.Println("aprovado com nota: ", nota)
	} else {
		fmt.Println("reprovado por nota: ", nota)
	}
}

/*func main() {
	imprimirResultado(7.3)
	imprimirResultado(2.1)
}
*/
