package main

import (
	"fmt"
	"sort"
)

/*
- Partindo do código abaixo, ordene os []user por idade e sobrenome.
  - https://play.golang.org/p/BVRZTdlUZ_

- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/
type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type idade []user

func (p idade) Len() int                { return len(p) }
func (p idade) Less(indice, j int) bool { return p[indice].Age < p[j].Age }
func (p idade) Swap(indice, j int)      { p[indice], p[j] = p[j], p[indice] }

type sobrenome []user

func (p sobrenome) Len() int                { return len(p) }
func (p sobrenome) Less(indice, j int) bool { return p[indice].Last < p[j].Last }
func (p sobrenome) Swap(indice, j int)      { p[indice], p[j] = p[j], p[indice] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(idade(users))
	fmt.Println(users)
	sort.Sort(sobrenome(users))
	fmt.Println(users)
	for _, v := range users {
		sort.Strings(v.Sayings)

	}

}
