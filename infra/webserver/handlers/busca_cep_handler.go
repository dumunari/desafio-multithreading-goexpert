package handlers

import (
	"desafio-multithreading/internal/entity"
	"desafio-multithreading/usecase/buscacep"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	cepParam := chi.URLParam(r, "cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := buscacep.BuscaCepMaisRapido(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := &entity.Response{
			Message: err.Error(),
			Service: "internal",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
