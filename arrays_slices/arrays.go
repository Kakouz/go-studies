package main

func Soma(numeros []int) int {
	sum := 0

	for _, numero := range numeros {
		sum += numero
	}

	return sum
}

func SomaTudo(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        somas = append(somas, Soma(numeros))
    }

    return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        if len(numeros) == 0 {
            somas = append(somas, 0)
        } else {
            final := numeros[1:]
            somas = append(somas, Soma(final))
        }
    }

    return somas
}