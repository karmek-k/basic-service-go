package calculator_test

import (
	"testing"

	"github.com/karmek-k/basic-service-go/pkg/calculator"
	"github.com/karmek-k/basic-service-go/test"
)

func TestSubtract(t *testing.T) {
	test.AssertEqualInts(t, 0, calculator.Subtract(2, 2))
	test.AssertEqualInts(t, 4, calculator.Subtract(2, -2))
	test.AssertEqualInts(t, 0, calculator.Subtract(0, 0))
	test.AssertEqualInts(t, 9, calculator.Subtract(7, -2))
	test.AssertEqualInts(t, 0, calculator.Subtract(-50, -50))
	test.AssertEqualInts(t, -377, calculator.Subtract(123, 500))
}