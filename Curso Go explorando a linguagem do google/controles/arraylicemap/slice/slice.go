package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice
	// quando vc decide fixo é um array, tamanho variavel é um slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	//você pode pensar no slice como uma estrutura de ponteiros que aponta para os itens de um array e sabe onde iniciar, tamanho do slice e a capacidade do array.
	//slice não é um array! Slide define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println(a2, s2)
	// slice tamanho e um ponteiro para um elemento de array
	s3 := a2[:2]
	fmt.Println(a2, s3)
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
