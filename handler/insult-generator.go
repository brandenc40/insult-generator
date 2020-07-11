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
	// APISuccess - success response
	APISuccess = "SUCCESS"
	// APIError - error response
	APIError = "ERROR"

	complimentTemplate = "You are %s and %s%s. %s - you're %s."
	insultTemplate     = "%s %s and %s%s. Now %s."
)

// InsultGenerator -
type InsultGenerator struct {
	data   *models.InsultData
	rand   rand.Source
	Logger logging.Logger
}

// NewInsultGenerator creates a new insultGenerator object
func NewInsultGenerator() (*InsultGenerator, error) {
	insultData, err := getInsultData("data/insults.json")
	if err != nil {
		return nil, err
	}
	return &InsultGenerator{
		data:   insultData,
		rand:   rand.NewSource(time.Now().Unix()),
		Logger: logging.NewLogger("insult-generator"),
	}, nil
}

func getInsultData(filename string) (*models.InsultData, error) {
	var data models.InsultData
	dataFile, err := os.Open(filename)
	defer dataFile.Close()
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(dataFile)
	jsonParser.Decode(&data)
	return &data, nil
}

// GetInsult -
func (g *InsultGenerator) GetInsult(w http.ResponseWriter, r *http.Request) {
	g.Logger.Debug("Endoint hit: /insult")

	insults := g.data.Insults
	rand := rand.New(g.rand)
	message := fmt.Sprintf(
		insultTemplate,
		insults[0][rand.Intn(len(insults[0]))],
		insults[1][rand.Intn(len(insults[1]))],
		insults[2][rand.Intn(len(insults[2]))],
		insults[3][rand.Intn(len(insults[3]))],
		insults[4][rand.Intn(len(insults[4]))],
	)
	response := models.APIResponse{
		Status:  APISuccess,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

// GetCompliment -
func (g *InsultGenerator) GetCompliment(w http.ResponseWriter, r *http.Request) {
	g.Logger.Debug("Endoint hit: /compliemnt")
	compliments := g.data.Compliments
	rand := rand.New(g.rand)
	message := fmt.Sprintf(
		complimentTemplate,
		compliments[0][rand.Intn(len(compliments[0]))],
		compliments[1][rand.Intn(len(compliments[1]))],
		compliments[2][rand.Intn(len(compliments[2]))],
		compliments[3][rand.Intn(len(compliments[3]))],
		compliments[4][rand.Intn(len(compliments[4]))],
	)
	response := models.APIResponse{
		Status:  APISuccess,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

// GetComeback -
func (g *InsultGenerator) GetComeback(w http.ResponseWriter, r *http.Request) {
	g.Logger.Debug("Endoint hit: /comeback")
	comebacks := g.data.Comebacks
	rand := rand.New(g.rand)
	message := comebacks[rand.Intn(len(comebacks[0]))]
	response := models.APIResponse{
		Status:  APISuccess,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}
