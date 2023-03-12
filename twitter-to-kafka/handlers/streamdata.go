package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"twitter-to-kafka/packages/model"
)

type Data struct {
	l *log.Logger
}

func NewData(l *log.Logger) *Data {
	return &Data{l}
}

func (h *Data) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Bad request", http.StatusNotFound)
		return
	}

	inputData := model.InputData{}

	//Deserializing to inputData
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		h.l.Println("[ERROR] Deserializing Input Data", err)
		http.Error(rw, "Error reading input data", http.StatusBadRequest)
		return
	}

	err = inputData.Validate()
	if err != nil {
		h.l.Println("[ERROR] Validating Input Data", err)
		http.Error(
			rw,
			fmt.Sprintf("Error reading input data"),
			http.StatusBadRequest,
		)
		return
	}

	h.l.Printf("Data %s\n", inputData)
}
