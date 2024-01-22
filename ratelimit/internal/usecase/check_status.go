package usecase

import (
	"encoding/json"
	"net/http"
)

type CheckStatusResponse struct {
	Messsage string `json:"message"`
}

func CheckStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := CheckStatusResponse{
		Messsage: "OK",
	}

	json.NewEncoder(w).Encode(response)
}
