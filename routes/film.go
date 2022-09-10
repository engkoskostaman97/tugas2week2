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

	r.HandleFunc("/films", h.FindFilms).Methods("GET")         //get alll
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")       //select
	r.HandleFunc("/film", h.CreateFilm).Methods("POST")        // add
	r.HandleFunc("/film/{id}", h.UpdateFilm).Methods("PATCH")  // edite
	r.HandleFunc("/film/{id}", h.DeleteFilm).Methods("DELETE") // delete

}
