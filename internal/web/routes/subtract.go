package routes

import (
	"io"
	"net/http"
)

// Subtract subtracts two numbers (|n| <= 1000)
func Subtract(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "subtracting two numbers")
}