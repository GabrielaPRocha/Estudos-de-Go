package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)
	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang"}
	fmt.Println(coisa2)
}
