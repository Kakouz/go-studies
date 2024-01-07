package interacao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', resultado '%s'", esperado, repeticoes)
	}
}
