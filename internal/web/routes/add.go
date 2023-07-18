package routes

import (
	"io"
	"net/http"
)

// Add adds two numbers (|n| <= 1000)
func Add(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "adding two numbers")
}