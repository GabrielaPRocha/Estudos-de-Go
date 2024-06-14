package main

import "fmt"

/*- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.*/
func retorno() func() {
	return func() {
		fmt.Println("Função que retorna uma função")
	}
}
func main() {
	x := retorno()
	x()
}
