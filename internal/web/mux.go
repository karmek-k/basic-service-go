package web

import (
	"net/http"

	"github.com/karmek-k/basic-service-go/internal/web/routes"
)

// CreateCalculatorMux creates a mux with registered calculator routes
func CreateCalculatorMux() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/add", routes.Add)
	mux.HandleFunc("/subtract", routes.Subtract)

	return mux
}