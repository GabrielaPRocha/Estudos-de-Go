package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	fmt.Println(x)
	y--
	fmt.Print(y)
}

// Go faz incrimento apenas postfix
//é um pos fixado, você não consegue usar algo como x++ == y--
