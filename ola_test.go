package main

import "testing"

func TestOla(t *testing.T) {

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris")
		esperado := "Olá, Chris"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, mundo"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})
}
