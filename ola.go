package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(name string, idioma string) string {

	if name == "" {
		name = "mundo"
	}

	return prefixoDeSaudacao(idioma) + name
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("Chris", ""))
}
