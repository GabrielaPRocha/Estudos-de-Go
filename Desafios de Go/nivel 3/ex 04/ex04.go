package main

import "fmt"

func main() {
	anoAtual := 2024
	anoNasci := 2003
	for {
		if anoAtual <= anoNasci {
			break
		}
		anoAtual--
		fmt.Print(anoAtual)

	}

}
