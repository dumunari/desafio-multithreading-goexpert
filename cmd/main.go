package main

import (
	"desafio-multithreading/infra/webserver/handlers"
	"net/http"

	_ "desafio-multithreading/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// @title Busca CEP Server
// @version 1.0
// @description This is a busca cep server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", handlers.BuscaCepHandler)
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	http.ListenAndServe(":8080", r)
}
