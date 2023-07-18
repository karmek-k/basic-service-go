package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/karmek-k/basic-service-go/internal/web/constants"
	"github.com/karmek-k/basic-service-go/internal/web/dto"
)

// ArgsFromParamsMiddleware provides calculator args in the request context,
// extracted from the query string.
//
// Writes a 400 error if they cannot be found or converted.
func ArgsFromParamsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args, err := getArgs(r.URL.Query())

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		// TODO use a key of different type
		ctx := context.WithValue(r.Context(), constants.CalculatorArgsContextKey, args)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getArgs(query url.Values) (*dto.CalculatorArguments, error) {
	queryA := query.Get("a")
	a, err := strconv.Atoi(queryA)

	if err != nil {
		return nil, fmt.Errorf("not an integer: %v", queryA)
	}

	queryB := query.Get("b")
	b, err := strconv.Atoi(queryB)

	if err != nil {
		return nil, fmt.Errorf("not an integer: %v", queryB)
	}

	return &dto.CalculatorArguments{
		A: a,
		B: b,
	}, nil
}