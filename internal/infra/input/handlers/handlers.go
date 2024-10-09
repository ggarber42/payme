package handlers

import (
	"encoding/json"
	"net/http"
)

func healthCheckerHandler(w http.ResponseWriter, r *http.Request){
	data := map[string]bool{
		"ok": true,
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed health", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(jsonData))

}