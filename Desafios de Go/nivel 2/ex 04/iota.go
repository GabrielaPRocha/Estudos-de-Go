package main

/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.*/
import (
	"fmt"
)

const (
	_ = 2024 + iota
	x
	y
	z
	a
)

func main() {
	fmt.Print(x, y, z, a)
}
