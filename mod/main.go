package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":8000", r) // this might give error so we use log.fatal to cover if it fails.
	log.Fatal(http.ListenAndServe(":8000", r))
}

func greeter() {
	fmt.Println("hey there user")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey prerak</h1>"))
}
