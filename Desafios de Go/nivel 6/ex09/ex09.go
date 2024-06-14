package main

import "fmt"

//- Callback: passe uma função como argumento a outra função.

func main() {
	Receptora(Enviadora)

}

func Enviadora() {
	fmt.Println("Essa função está enviando para a receptora")
}
func Receptora(x func()) {
	fmt.Println("Essa função vai receber a outra: ")
	x()
}
