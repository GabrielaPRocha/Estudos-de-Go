package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "and64" {
		t.Skip("Não funciona em arquitetura and64")
	}
	//...
	t.Fail()
}
