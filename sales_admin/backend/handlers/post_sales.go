package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/clevengermatt/challenges/sales_admin/backend/db"
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

	file, handler, err := r.FormFile("file")
	if err != nil {
		respondWithError(w, "Missing Parameters: file", http.StatusBadRequest)
		return
	}
	defer file.Close()

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

	sales, err := salesFrom(file)
	if err != nil {
		respondWithError(w, "Invalid Parameters: file, improperly formatted csv data", http.StatusBadRequest)
		return
	}

	err = db.StoreSales(sales)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	salesJSON, err := json.Marshal(sales)
	if err != nil {
		respondWithError(w, "Oops... Someting went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(salesJSON))
}

func salesFrom(csvFile multipart.File) ([]db.Sale, error) {

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var sale db.Sale
	var sales []db.Sale

	for index, columns := range data {

		// skip the row of column names
		if index == 0 {
			continue
		}

		itemPrice, err := strconv.ParseFloat(columns[2], 64)
		if err != nil {
			return nil, err
		}

		quanitity, err := strconv.Atoi(columns[3])
		if err != nil {
			return nil, err
		}

		sale = db.Sale{
			CustomerName:    columns[0],
			ItemDescription: columns[1],
			ItemPrice:       itemPrice,
			Quantity:        quanitity,
			MerchantName:    columns[4],
			MerchantAddress: columns[5],
		}

		sales = append(sales, sale)
	}

	return sales, nil
}
