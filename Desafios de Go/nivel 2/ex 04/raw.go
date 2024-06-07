package main

import "fmt"

/*- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.*/

func main() {
	x := `você usa o  acento agudo
						e ai consegue
						fazer essa formatação
												meio diferente`
	fmt.Println(x)
}
