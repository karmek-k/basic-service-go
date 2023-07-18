package main

import (
	"net/http"

	"github.com/karmek-k/basic-service-go/internal/web"
)

func main() {
	r := web.CreateCalculatorHandler()

	http.ListenAndServe(":8000", r)
}