package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/canis-sirius/bookings/pkg/handlers"

	"github.com/go-chi/chi"

	"github.com/canis-sirius/bookings/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}