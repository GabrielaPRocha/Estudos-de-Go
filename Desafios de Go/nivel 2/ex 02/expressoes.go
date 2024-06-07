package main

import "fmt"

/*- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.
==
!=
>
<
=>
<=
*/

func main() {
	x := (5 == 42)
	y := (5 != 42)
	z := (5 <= 42)
	w := (5 < 42)
	a := (5 >= 42)
	b := (5 > 42)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", x, y, z, w, a, b)

}
