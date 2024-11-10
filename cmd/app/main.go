package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/services"
	"awesomeProject/internal/storages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Configuration not found: %v", err)
	}

	db, err := storages.NewMongo(cfg.Database.URI, cfg.Database.Name)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	r := repositories.NewRepository(db, cfg)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Post("/", h.ShortenUrl)
	mux.Get("/{shorten_url}", h.RedirectUrl)

	log.Printf("The server is running at %v", cfg.Server.BindAddr)
	err = http.ListenAndServe(cfg.Server.BindAddr, mux)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
