package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brandenc40/insult-generator/handler"
	"github.com/brandenc40/insult-generator/logging"
	"github.com/gorilla/mux"
)

func main() {
	logger := logging.NewLogger("main")
	logger.SetLevel(logging.DEBUG)

	insulter, err := handler.NewInsultGenerator()
	if err != nil {
		logger.Critical("Unable to build the InsultGenerator handler")
		logger.Critical(err.Error())
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		logger.Warning("The $PORT env variale is not set, defaulting to 10000")
		port = "10000"
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<H1>Insult generator API</H1>Available endoints are:</br> - /insult</br> - /comeback</br> - /compliment")
	})

	router.HandleFunc("/insult", insulter.GetInsult).Methods("GET")
	router.HandleFunc("/comeback", insulter.GetComeback).Methods("GET")
	router.HandleFunc("/compliment", insulter.GetCompliment).Methods("GET")

	logger.Info("Starting app at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
