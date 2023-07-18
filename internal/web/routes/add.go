package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/karmek-k/basic-service-go/internal/web/constants"
	"github.com/karmek-k/basic-service-go/internal/web/dto"
	"github.com/karmek-k/basic-service-go/pkg/calculator"
)

// Add adds two numbers
func Add(w http.ResponseWriter, r *http.Request) {
	args := r.Context().Value(constants.CalculatorArgsContextKey).(*dto.CalculatorArguments)

	result := calculator.Add(args.A, args.B)

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf("{\"result\": %d}", result))
}