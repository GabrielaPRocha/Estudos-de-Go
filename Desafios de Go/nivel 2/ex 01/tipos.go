package main

import "fmt"

//Escreva um programa que mostre um número em decimal, binário, e hexadecimal
func main() {
	x := 42
	fmt.Printf("%b\n %d\n %#x", x, x, x)
}
