package main

import "fmt"

/*é uma estrutura com tipos pre definidos*/
type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	//vc tem duas formas de declara, essa:
	cliente1 := cliente{
		nome:      "joao",
		sobrenome: "Souza",
		fumante:   true,
	}

	//e essa:
	cliente2 := cliente{"Joana", "Silva", false}
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
