package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(name string) string {
	if name == "" {
		return prefixoOlaPortugues + "mundo"
	}

	return prefixoOlaPortugues + name
}
func main() {
	fmt.Println(Ola("Chris"))
}
