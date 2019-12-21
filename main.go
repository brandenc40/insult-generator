package main

import (
	"fmt"
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
		port = "10000"
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<H1>Insult generator API</H1>Available endoints are:</br> - /insult</br> - /comeback</br> - /compliment")
	})

	router.HandleFunc("/insult", insulter.GetInsult).Methods("GET")
	router.HandleFunc("/comeback", insulter.GetComeback).Methods("GET")
	router.HandleFunc("/compliment", insulter.GetCompliment).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
