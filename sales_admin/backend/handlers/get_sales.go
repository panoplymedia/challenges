package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clevengermatt/challenges/sales_admin/backend/db"
)

// GetSales will return the sales data from the database as JSON
func GetSales(w http.ResponseWriter, r *http.Request) {

	sales, _ := db.GetSales()

	salesJSON, err := json.Marshal(sales)
	if err != nil {
		respondWithError(w, "Oops... Someting went wrong", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(salesJSON))
}
