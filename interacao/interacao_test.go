package interacao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {

	t.Run("repete a string 5 vezes", func(t *testing.T) {
		repeticoes := Repetir("a", 0)
		esperado := "aaaaa"

		if repeticoes != esperado {
			t.Errorf("esperado '%s', resultado '%s'", esperado, repeticoes)
		}
	})

	t.Run("repete a string 20 vezes", func(t *testing.T) {
		repeticoes := Repetir("a", 20)
		esperado := "aaaaaaaaaaaaaaaaaaaa"

		if repeticoes != esperado {
			t.Errorf("esperado '%s', resultado '%s'", esperado, repeticoes)
		}
	})

}

func ExampleRepetir() {
	repeticoes := Repetir("a", 5)
	fmt.Println(repeticoes)
	// Output: aaaaa
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 0)
	}
}
