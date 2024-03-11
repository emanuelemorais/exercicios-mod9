package rxwlib900

import(
	"testing"
)

func TestCreateGasesValues(t *testing.T) {
	data := CreateGasesValues("test")
	if data.Radiation > radiationRange["radiation"].MinValue || data.Radiation < radiationRange["radiation"].MaxValue {
		t.Errorf("Valor de radiação fora do intervalo esperado")
	}
}