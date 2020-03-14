package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func respondWithError(w http.ResponseWriter, msg string, sts int) {
	errorResponse := errorResponse{
		Message: msg,
	}

	errRespJSON, err := json.Marshal(errorResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sts)
	fmt.Fprint(w, string(errRespJSON))
}
