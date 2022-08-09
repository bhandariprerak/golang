package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/prerakbhandari/mongoApi/controller"
)

func Routes() *mux.Router {
	fmt.Println("welcome to router")

	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	return r

}
