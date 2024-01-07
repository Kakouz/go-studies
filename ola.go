package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(name string, idioma string) string {

	if name == "" {
		return prefixoOlaPortugues + "mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol

	case frances:
		prefixo = prefixoOlaFrances
	}

	return prefixo + name
}
func main() {
	fmt.Println(Ola("Chris", ""))
}
