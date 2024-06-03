package matematica

import "testing"

const erroPadrão = "Valor esperado %v, mas o valor encontrado fo i%v"

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsperado {
		t.Errorf(erroPadrão, valorEsperado, valor)
	}

}
