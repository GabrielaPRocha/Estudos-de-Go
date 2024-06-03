package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - indices: esperado (%d) <> econtrado (%d)"

func TestIndex(t *testing.T) {
	testes := []struct {
		texto     string
		parte     string
		esperando int
	}{
		{"Cod3r é show", "Cod3t", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Leonardo", "o", 2},
	}
	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)
		if atual != teste.esperando {
			t.Error(msgIndex, teste.texto, teste.parte, teste.esperando, atual)
		}
	}
}
