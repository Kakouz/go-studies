package main

func Soma(numeros [5]int) int {
	result := 0
	for i := 0; i < 5; i++ {
		result += numeros[i]
	}
	return result
}
