package buscacep

import (
	"context"
	"desafio-multithreading/internal/dto"
	"desafio-multithreading/internal/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func BuscaCepMaisRapido(cep string) (*entity.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resultCh := make(chan *entity.Response, 1)
	errCh := make(chan error, 1)

	go buscaViaCEP(cep, resultCh, errCh)
	go buscaBrasilApiCEP(cep, resultCh, errCh)

	select {
	case respostaRapida := <-resultCh:
		return respostaRapida, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func buscaViaCEP(cep string, resultCh chan *entity.Response, errorCh chan error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		errorCh <- err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorCh <- err
	}
	var viaCep dto.ViaCEP
	err = json.Unmarshal(body, &viaCep)
	if err != nil {
		errorCh <- err
	}

	response := &entity.Response{
		Message: fmt.Sprintf("O CEP %s é da cidade de %s, no estado de %s", viaCep.Cep, viaCep.Localidade, viaCep.Estado),
		Service: "ViaCEP",
	}

	resultCh <- response
}

func buscaBrasilApiCEP(cep string, resultCh chan *entity.Response, errorCh chan error) {
	resp, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		errorCh <- err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorCh <- err
	}

	var brasilApi dto.BrasilApi
	err = json.Unmarshal(body, &brasilApi)
	if err != nil {
		errorCh <- err
	}

	response := &entity.Response{
		Message: fmt.Sprintf("O CEP %s é da cidade de %s, no estado de %s", brasilApi.Cep, brasilApi.Localidade, brasilApi.Estado),
		Service: "BrasilAPI",
	}

	resultCh <- response
}
