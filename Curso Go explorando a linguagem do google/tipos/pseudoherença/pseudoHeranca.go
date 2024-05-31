package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}
type ferrari struct {
	carro       // cria um campo anonimo, então tudo que tem em carros estara em ferrari, não está herdando
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Printf("a ferrari %s esta com o turbo ligado? %v\n", f.nome, f.turboLigado)
}
