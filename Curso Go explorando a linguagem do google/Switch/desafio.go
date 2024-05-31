package main

import "fmt"

func nota(n float64) string {
	switch {
	case n > 9 && n <= 10:
		return "A"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}

}
func main() {
	fmt.Println(nota(10))
}
