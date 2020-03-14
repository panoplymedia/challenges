package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/clevengermatt/challenges/sales_admin/backend/db"
)

func main() {

	// give the database some time to start before trying to connect
	time.Sleep(2 * time.Second)

	// Creat the database table
	err := db.CreateSaleTable()
	if err != nil {
		log.Fatal(err)
	}

	// start listening for api requests
	r := mux.NewRouter()
	r.StrictSlash(true)

	for _, endpoint := range endpoints {
		mthds := strings.Join(endpoint.Methods, ", ")
		r.HandleFunc(endpoint.Path, endpoint.Handler).Methods(mthds)
	}

	http.Handle("/", r)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
