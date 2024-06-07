package main

import "fmt"

/*- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
        U+0041 'A'
        U+0041 'A'
        U+0041 'A'
    66
        U+0042 'B'
        U+0042 'B'
        U+0042 'B' */
func main() {
	for x := 65; x <= 90; x++ {
		fmt.Print(x)
		for y := 1; y <= 3; y++ {
			fmt.Printf("%#U\n", x)
			fmt.Printf("%#U\n", x)
			fmt.Printf("%#U\n", x)
		}
	}
}
