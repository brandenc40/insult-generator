package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/brandenc40/insult-generator/logging"
	"github.com/brandenc40/insult-generator/models"
)

const (
	// API response statuses
	ApiSuccess = "SUCCESS"
	ApiError   = "ERROR"
)

type insultGenerator struct {
	data   models.InsultData
	rand   rand.Source
	Logger logging.Logger
}

// NewInsultGenerator creates a new insultGenerator object
func NewInsultGenerator() insultGenerator {
	return insultGenerator{
		data:   getInsultData("data/insults.json"),
		rand:   rand.NewSource(time.Now().Unix()),
		Logger: logging.NewLogger("insult-generator"),
	}
}

func getInsultData(filename string) models.InsultData {
	var data models.InsultData
	dataFile, err := os.Open(filename)
	defer dataFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(dataFile)
	jsonParser.Decode(&data)
	return data
}

func (g *insultGenerator) GetInsult(w http.ResponseWriter, r *http.Request) {
	g.Logger.Info("Endoint hit: /insult")
	insult_template := "%s %s and %s%s. Now %s."
	insults := g.data.Insults
	rand := rand.New(g.rand)
	message := fmt.Sprintf(
		insult_template,
		insults[0][rand.Intn(len(insults[0]))],
		insults[1][rand.Intn(len(insults[1]))],
		insults[2][rand.Intn(len(insults[2]))],
		insults[3][rand.Intn(len(insults[3]))],
		insults[4][rand.Intn(len(insults[4]))],
	)
	response := models.ApiResponse{
		Status:  ApiSuccess,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

func (g *insultGenerator) GetCompliment(w http.ResponseWriter, r *http.Request) {
	g.Logger.Info("Endoint hit: /compliemnt")
	compliment_template := "You are %s and %s%s. %s - you're %s."
	compliments := g.data.Compliments
	rand := rand.New(g.rand)
	message := fmt.Sprintf(
		compliment_template,
		compliments[0][rand.Intn(len(compliments[0]))],
		compliments[1][rand.Intn(len(compliments[1]))],
		compliments[2][rand.Intn(len(compliments[2]))],
		compliments[3][rand.Intn(len(compliments[3]))],
		compliments[4][rand.Intn(len(compliments[4]))],
	)
	response := models.ApiResponse{
		Status:  ApiSuccess,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

func (g *insultGenerator) GetComeback(w http.ResponseWriter, r *http.Request) {
	g.Logger.Info("Endoint hit: /comeback")
	comebacks := g.data.Comebacks
	rand := rand.New(g.rand)
	message := comebacks[rand.Intn(len(comebacks[0]))]
	response := models.ApiResponse{
		Status:  ApiSuccess,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}
