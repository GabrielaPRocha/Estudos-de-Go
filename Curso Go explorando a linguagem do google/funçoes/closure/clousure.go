package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}

/*
Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto
ex X vale 1 2 3
se eu fizer um closure na função A ele vai valer 1 2 3
se eu fizer um closure na funcção B ele também sera 1 2 3
- Exemplo de closure:
    - func i() func() int { x := 0; return func() int { x++; return x } }
    - Quando fizermos a := i() teremos um scope, um valor para x.
    - Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
- Closures nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código.
*/
