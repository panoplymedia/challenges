package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getSalesTotalResponse struct {
	Total int `json:"total"`
}

// GetSalesTotal will return the sum of sales
func GetSalesTotal(w http.ResponseWriter, r *http.Request) {

	resp := getSalesTotalResponse{
		Total: 0,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		respondWithError(w, "Opps... something went wrong.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(respJSON))
}
