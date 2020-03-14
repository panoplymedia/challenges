package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// PostSales will parse and store sales data from a cvs file.
func PostSales(w http.ResponseWriter, r *http.Request) {

	ct := "multipart/form-data"
	reqCT := r.Header.Get("Content-Type")
	reqCT = strings.Split(reqCT, ";")[0]

	if strings.ToLower(reqCT) != ct {
		respondWithError(w, "Extpected Content-Type: multipart/form-data", http.StatusUnsupportedMediaType)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		respondWithError(w, "Invalid Parameters: file exceeds 10MB maximum file size", http.StatusBadRequest)
		return
	}

	_, handler, err := r.FormFile("file")
	if err != nil {
		respondWithError(w, "Missing Parameters: file", http.StatusBadRequest)
		return
	}

	fileName := handler.Filename

	comps := strings.Split(fileName, ".")

	isCSV := false
	if len(comps) > 0 {
		ext := comps[len(comps)-1]
		isCSV = strings.ToLower(ext) == "csv"
	}

	if !isCSV {
		respondWithError(w, "Invalid Parameters: file, expected type .csv", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "")
}
