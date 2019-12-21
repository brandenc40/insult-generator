package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brandenc40/insult-generator/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	insulter := handler.NewInsultGenerator()

	router := mux.NewRouter().StrictSlash(true)
	r := gin.Default()

	router.HandleFunc("/insult", insulter.GetInsult).Methods("GET")
	router.HandleFunc("/comeback", insulter.GetComeback).Methods("GET")
	router.HandleFunc("/compliment", insulter.GetCompliment).Methods("GET")

	fmt.Print("http://localhost:10000/\n")
	log.Fatal(http.ListenAndServe(":10000", router))
}
