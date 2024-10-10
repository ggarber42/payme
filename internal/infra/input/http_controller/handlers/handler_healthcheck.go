package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthCheckerHandler(w http.ResponseWriter, r *http.Request){
	data := struct {
		Status string `json:"status"`
		Timestamp string `json:"timestamp"`
	}{
		Status: "Ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed health", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))

}
