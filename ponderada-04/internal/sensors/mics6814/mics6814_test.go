package mics6814

import(
	"testing"
)

func TestCreateGasesValues(t *testing.T) {
	data := CreateGasesValues("test")
	if data.CarbonMonoxide > gasesRange["carbon_monoxide"].MinValue || data.CarbonMonoxide < gasesRange["carbon_monoxide"].MaxValue {
		t.Errorf("Valor de CO fora do intervalo esperado")
	}
	if data.NitrogenDioxide > gasesRange["nitrogen_dioxide"].MinValue || data.NitrogenDioxide < gasesRange["nitrogen_dioxide"].MaxValue {
		t.Errorf("Valor de NO2 fora do intervalo esperado")
	}
	if data.Ethanol > gasesRange["ethanol"].MinValue || data.Ethanol < gasesRange["ethanol"].MaxValue {
		t.Errorf("Valor de EtOH fora do intervalo esperado")
	}
	if data.Hydrogen > gasesRange["hydrogen"].MinValue || data.Hydrogen < gasesRange["hydrogen"].MaxValue {
		t.Errorf("Valor de H2 fora do intervalo esperado")
	}
	if data.Ammonia > gasesRange["ammonia"].MinValue || data.Ammonia < gasesRange["ammonia"].MaxValue {
		t.Errorf("Valor de NH3 fora do intervalo esperado")
	}
	if data.Methane > gasesRange["methane"].MinValue || data.Methane < gasesRange["methane"].MaxValue {
		t.Errorf("Valor de CH4 fora do intervalo esperado")
	}
	if data.Propane > gasesRange["propane"].MinValue || data.Propane < gasesRange["propane"].MaxValue {
		t.Errorf("Valor de C3H8 fora do intervalo esperado")
	}
	if data.IsoButane > gasesRange["iso_butane"].MinValue || data.IsoButane < gasesRange["iso_butane"].MaxValue {
		t.Errorf("Valor de i-C4H10 fora do intervalo esperado")
	}
}