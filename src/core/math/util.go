package math

import (
	"math"
)

const EPSILON = 1e-5

func IsFloatEqual(a, b float64) bool {
	return math.Abs(a-b) <= EPSILON
}

func Clamp(f float64) float64 {
	if f > 1 {
		return 1
	}

	if f < 0 {
		return 1
	}

	return f
}
