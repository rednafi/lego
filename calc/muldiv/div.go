package muldiv

import "math"

// Div two integers.
func Div(a, b int) (float64, error) {
	return math.Floor(float64(a) / float64(b)), nil
}
