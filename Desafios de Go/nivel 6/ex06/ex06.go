package main

import "fmt"

// crie e utilize uma função anonima

func main() {
	func() {
		fmt.Println("Função anonima")
	}()
}
