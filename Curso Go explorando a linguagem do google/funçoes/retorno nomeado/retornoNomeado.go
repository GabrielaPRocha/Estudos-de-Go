package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo pois é um retorno nomeado
}
func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
