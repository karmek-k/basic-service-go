package web

import (
	"net/http"

	"github.com/karmek-k/basic-service-go/internal/web/middleware"
	"github.com/karmek-k/basic-service-go/internal/web/routes"
)

// CreateCalculatorHandler creates a `http.Handler` with registered calculator routes
func CreateCalculatorHandler() http.Handler {
	mux := http.NewServeMux()
	
	handler := middleware.ArgsFromParamsMiddleware(mux)
	// TODO fix; always returns 400
	// handler = middleware.ValidArgsMiddleware(handler)

	mux.HandleFunc("/add", routes.Add)
	mux.HandleFunc("/subtract", routes.Subtract)

	return handler
}