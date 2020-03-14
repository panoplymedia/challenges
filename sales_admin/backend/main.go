package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/clevengermatt/challenges/sales_admin/backend/db"
	"github.com/gorilla/mux"
)

func main() {
	startDatabase()
	startAPI()
}

func startAPI() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	for _, endpoint := range endpoints {
		mthds := strings.Join(endpoint.Methods, ", ")
		r.HandleFunc(endpoint.Path, endpoint.Handler).Methods(mthds)
	}

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func startDatabase() {
	err := db.CreateConnectionPool()
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateSaleTable()
	if err != nil {
		log.Fatal(err)
	}
}
