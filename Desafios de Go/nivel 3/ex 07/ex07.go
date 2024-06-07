package main

import "fmt"

func main() {
	x := 5
	if x < 1 {
		fmt.Println("menor que 1")
	} else if x == 5 {
		fmt.Println("igual a 5")
	} else {
		fmt.Println("diferente das opções")
	}
}
