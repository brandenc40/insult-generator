package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brandenc40/insult-generator/handler"
	"github.com/gorilla/mux"
)

func main() {
	insulter := handler.NewInsultGenerator()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/insult", insulter.GetInsult).Methods("GET")
	router.HandleFunc("/comeback", insulter.GetComeback).Methods("GET")
	router.HandleFunc("/compliment", insulter.GetCompliment).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
