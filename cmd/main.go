package main

import (
	"desafio-multithreading/infra/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", handlers.BuscaCepHandler)
	})
	http.ListenAndServe(":8080", r)
}
