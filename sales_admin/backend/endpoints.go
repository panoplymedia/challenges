package main

import (
	"net/http"

	"github.com/clevengermatt/challenges/sales_admin/backend/handlers"
)

// Endpoint is an API endpoint's handler, HTTP method, and url path
type Endpoint struct {
	Handler func(http.ResponseWriter, *http.Request)
	Methods []string
	Path    string
}

// The collection of available endpoints on the API
var endpoints = []Endpoint{
	Endpoint{
		Handler: handlers.PostSales,
		Methods: []string{"POST"},
		Path:    "/sales/",
	},
	Endpoint{
		Handler: handlers.GetSales,
		Methods: []string{"GET"},
		Path:    "/sales/",
	},
}
