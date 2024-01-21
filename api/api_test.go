package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes        map[string]int
	registrosVitorias []string
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func (e *EsbocoArmazenamentoJogador) RegistrarVitoria(nome string) {
	e.registrosVitorias = append(e.registrosVitorias, nome)
}

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		}, nil,
	}
	servidor := &ServidorJogador{&armazenamento}

	t.Run("retorna pontuacao de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoRequisicao(t, resposta.Body.String(), "20")
	})

	t.Run("retorna pontuacao de Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoRequisicao(t, resposta.Body.String(), "10")
	})

	t.Run("retorna 404 para jogador não encontrado", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Jorge")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Code
		esperado := http.StatusNotFound

		if recebido != esperado {
			t.Errorf("recebido status %d esperado %d", recebido, esperado)
		}
	})
}

func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil,
	}
	servidor := &ServidorJogador{&armazenamento}

	t.Run("registra vitorias na chamada ao método HTTP POST", func(t *testing.T) {
		jogador := "Maria"

		requisicao := novaRequisicaoRegistrarVitoriaPost(jogador)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusAccepted)

		if len(armazenamento.registrosVitorias) != 1 {
			t.Errorf("verifiquei %d chamadas a RegistrarVitoria, esperava %d", len(armazenamento.registrosVitorias), 1)
		}

		if armazenamento.registrosVitorias[0] != jogador {
			t.Errorf("não registrou o vencedor corretamente, recebi '%s', esperava '%s'", armazenamento.registrosVitorias[0], jogador)
		}
	})
}
func TestRegistrarVitoriasEBuscarEstasVitorias(t *testing.T) {
	armazenamento := NovoArmazenamentoJogadorEmMemoria()
	servidor := ServidorJogador{armazenamento}
	jogador := "Maria"

	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))

	resposta := httptest.NewRecorder()
	servidor.ServeHTTP(resposta, novaRequisicaoObterPontuacao(jogador))
	verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

	verificarCorpoRequisicao(t, resposta.Body.String(), "3")
}

func novaRequisicaoRegistrarVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisição é inválido, recebido '%s' esperado '%s'", recebido, esperado)
	}
}

func verificarRespostaCodigoStatus(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("não recebeu código de status HTTP esperado, recebido %d, esperado %d", recebido, esperado)
	}
}
