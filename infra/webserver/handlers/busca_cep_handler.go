package handlers

import (
	"desafio-multithreading/internal/entity"
	"desafio-multithreading/pkg"
	"desafio-multithreading/usecase/buscacep"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// GetCEP godoc
// @Summary      Get a cep
// @Description  Get a cep
// @Tags         cep
// @Accept       json
// @Produce      json
// @Param        cep   path      string  true  "cep"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  error
// @Failure      500  {object}  entity.Response
// @Router       /cep/{cep} [get]
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	cepParam := chi.URLParam(r, "cep")

	if !pkg.ValidateCepFormat(cepParam) {
		w.WriteHeader(http.StatusBadRequest)
		response := &entity.Response{
			Message: fmt.Sprintf("O cep informado %s está em um formato inválido.", cepParam),
			Service: "internal",
		}
		json.NewEncoder(w).Encode(response)
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
