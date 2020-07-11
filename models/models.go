package models

// InsultData loaded from the data.json file and used to
// populate the varied responses
type InsultData struct {
	Insults     [][]string `json:"insults"`
	Compliments [][]string `json:"compliments"`
	Comebacks   []string   `json:"comebacks"`
}

// APIResponse is the response returned to the API user
type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
