package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prerakbhandari/mongoApi/router"
)

func main() {
	fmt.Println("Welcome to mongo db in go")
	fmt.Println("Server is getting started...")
	r := router.Routes()
	fmt.Println("port is up and running at : http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", r))

}
