package main

import "fmt"

/*- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
  - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
  - Do quinto ao último item do slice (incluindo o último item!)
  - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
  - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
  - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item*/
func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x[:3])
	fmt.Println(x[3:10])
	fmt.Println(x[2:7])
	fmt.Println(x[3:9])
	fmt.Println(x[2 : len(x)-1])

}
