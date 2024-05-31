package main

import "math"

//Iniciando com letra maiusula é Publico, visivel fora do pacote
//Inicializando com letra miniscula é Privado (visivel apenas dentro do pacote)

type Ponto struct {
	//Ponto representa uma coordenada no plano cartesiano
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// responsavel por calcular a distancia linear dos dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
