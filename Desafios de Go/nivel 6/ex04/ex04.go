package main

import "fmt"

/*- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.*/
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) demonstrar() {
	fmt.Println("Nome", p.nome, p.sobrenome, "\n essa pessoa tem:", p.idade, "anos de idade")
}
func main() {
	pessoa1 := pessoa{"Emma", "Swan", 27}
	pessoa1.demonstrar()
}
