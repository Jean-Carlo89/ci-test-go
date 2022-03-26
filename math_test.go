package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido:\nResultado: %d\nEsperado: %d", total, 30)
	}
}
