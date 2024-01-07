package main

func Soma(numeros [5]int) int {
	sum := 0
	
	for _, numero := range numeros {
		sum += numero
	}

	return sum
}
