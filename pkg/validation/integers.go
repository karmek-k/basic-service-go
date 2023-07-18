package validation

// InRange returns true if `value` is in the range between `min` and `max`.
func InRange(value, min, max int) bool {
	return value >= min && value <= max
}
