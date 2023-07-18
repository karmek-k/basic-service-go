package calculator_test

import (
	"testing"

	"github.com/karmek-k/basic-service-go/pkg/calculator"
	"github.com/karmek-k/basic-service-go/test"
)

func TestAdd(t *testing.T) {
	test.AssertEqualInts(t, 4, calculator.Add(2, 2))
	test.AssertEqualInts(t, 0, calculator.Add(0, 0))
	test.AssertEqualInts(t, 5, calculator.Add(7, -2))
	test.AssertEqualInts(t, -100, calculator.Add(-50, -50))
	test.AssertEqualInts(t, 623, calculator.Add(123, 500))
}