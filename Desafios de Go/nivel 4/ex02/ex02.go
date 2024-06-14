package main

import "fmt"

/* Usando uma literal composta
cie uma slice de tipo int
atribua 10 valores a ela
utile range para demostrar todos esses valores
e utilieze format printng para demonstrar seu tipo*/
func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)

}
