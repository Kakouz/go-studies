package interacao

const QUANTIDADE_REPETICOES = 5

func Repetir(caractere string, quantidadeRepeticoes int) string {
	var repeticoes string

	if quantidadeRepeticoes == 0 {
		 quantidadeRepeticoes = QUANTIDADE_REPETICOES
	}

	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}
