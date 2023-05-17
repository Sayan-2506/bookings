package main

import (
	"net/http"

	"github.com/Sayan-2506/bookings/internal/config"
	"github.com/Sayan-2506/bookings/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	// "github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availiability)
	mux.Post("/search-availability", handlers.Repo.PostAvailiability)
	mux.Post("/search-availability-json", handlers.Repo.AvailiabilityJSON)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
