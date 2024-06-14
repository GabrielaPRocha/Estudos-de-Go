package main

import "fmt"

/*- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.*/
func variadico(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func variadico2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func main() {
	si := []int{30, 30}
	fmt.Println(variadico(si...))
	fmt.Println(variadico2(si))

}
