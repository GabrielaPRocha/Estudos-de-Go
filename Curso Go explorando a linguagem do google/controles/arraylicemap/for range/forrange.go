package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta a quantidade de elementos, retirando os ... ele vira um slice
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}
	for _, num := range numeros { // colocar o _ faz com que ele ignore o indice do array
		fmt.Println(num)
	}

}
