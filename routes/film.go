package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")

}
