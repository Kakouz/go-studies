package main

import (
	"log"
	"net/http"
)

func NovoArmazenamentoJogadorEmMemoria() *ArmazenamentoJogadorEmMemoria {
	return &ArmazenamentoJogadorEmMemoria{map[string]int{}}
}

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func (s *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	s.armazenamento[nome]++
}

func (s *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return s.armazenamento[nome]
}

func main() {
	servidor := &ServidorJogador{NovoArmazenamentoJogadorEmMemoria()}

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}
