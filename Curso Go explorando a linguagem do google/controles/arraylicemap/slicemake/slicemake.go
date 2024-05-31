package main

import "fmt"

func main() {
	s := make([]int, 10) // criamos o slice atraves do make
	s[9] = 12            // ele mesmo cria o array
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
	// slice vai crescendo internamente, então se vc cola elemento a mais ele "cria outro array com o dobro de capacidade"
}
