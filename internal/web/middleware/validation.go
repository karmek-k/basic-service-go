package middleware

import (
	"net/http"

	"github.com/karmek-k/basic-service-go/internal/web/constants"
	"github.com/karmek-k/basic-service-go/internal/web/dto"
	"github.com/karmek-k/basic-service-go/pkg/validation"
)

// ValidArgsMiddleware validates calculator arguments in request context.
func ValidArgsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args, ok := r.Context().Value(constants.CalculatorArgsContextKey).(*dto.CalculatorArguments)

		if !ok {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
	
		if !(validation.InRange(args.A, constants.ArgumentMin, constants.ArgumentMax) && 
			validation.InRange(args.B, constants.ArgumentMin, constants.ArgumentMax)) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			
			return
		}

		next.ServeHTTP(w, r)
	})
}