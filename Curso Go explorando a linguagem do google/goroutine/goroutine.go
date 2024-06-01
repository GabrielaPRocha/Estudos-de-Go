package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (intereação %d)\n", pessoa, texto, i+1)
	}

}
func main() {
	fale("Maria", "Poq vc não falou comigo", 3)
	fale("Joao", "Só posso falar depois de vc", 1)
}
