package main

import (
	"http-pokeapi/internal/pokeapi"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type apiconfig struct {
	client pokeapi.Client
}

func main() {
	time.Sleep(2 * time.Second)
	cfg := &apiconfig{
		client: pokeapi.Client{},
	}
	port := "8080"
	defer cfg.client.CloseIdleConnections()

	r := chi.NewRouter()
	s := chi.NewRouter()

	r.Mount("/api", s)

	s.Get("/locations", cfg.FetchLocations)
	s.Get("/locations/{location}", cfg.FetchPokemons)
	s.Get("/pokemon/{name}", cfg.AboutPokemon)

	servmux := corsmiddleware(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: servmux,
	}
	log.Printf("The server is live on port %s\n", port)

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v", err)
	}
}
