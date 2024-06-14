package main

import "fmt"

/* - Crie uma função que retorne um int
- Crie outra função que retorne um int e uma string
- Chame as duas funções
- Demonstre seus resultados
*/

func retornaumint() int {
	return 27

}
func retornaNumeroeLetra() (int, string) {
	return 27, "teste"
}
func main() {
	fmt.Println(retornaumint())
	fmt.Println(retornaNumeroeLetra())

}
