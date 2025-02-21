package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return an response in JSON for the requester
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

// Erro return an erro in JSON format
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json: "erro"`
	}{
		Erro: erro.Error(),
	})
}
