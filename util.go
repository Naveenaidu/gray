package main

import (
	"math"
)

const floatEqualityThreshold = 1e-5

func isFloatEqual(a, b float64) bool {
	return math.Abs(a-b) <= floatEqualityThreshold
}
