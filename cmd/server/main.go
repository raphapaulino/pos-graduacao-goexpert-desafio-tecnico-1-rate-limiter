package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/raphapaulino/pos-graduacao-goexpert-desafio-tecnico-1-rate-limiter/cmd/configs"
	"github.com/raphapaulino/pos-graduacao-goexpert-desafio-tecnico-1-rate-limiter/limiter"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	rateLimiter := limiter.NewRateLimiter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(rateLimiter.LimitHandler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página Inicial!"))
	})
	fmt.Println("Server is starting on port: ", configs.WebServerPort)

	http.ListenAndServe("127.0.0.1:"+configs.WebServerPort, router)
}
