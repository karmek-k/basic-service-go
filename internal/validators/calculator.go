package validators

import (
	"github.com/karmek-k/basic-service-go/internal/constants"
	"github.com/karmek-k/basic-service-go/internal/web/dto"
	"github.com/karmek-k/basic-service-go/pkg/validation"
)

// IsArgsValid returns true if arguments are of supported range
func IsArgsValid(args *dto.CalculatorArguments) bool {
	return validation.InRange(args.A, constants.ArgumentMin, constants.ArgumentMax) &&
		validation.InRange(args.B, constants.ArgumentMin, constants.ArgumentMax)
}
