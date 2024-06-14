package main

import "fmt"

func main() {
	x := 5
	switch {
	case x < 1:
		fmt.Println("menor que 1")
	case x == 5:
		fmt.Println("igual a 5")
	default:
		fmt.Println("diferente das opções")
	}
}
