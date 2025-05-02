package main

import (
	"math"
)

const floatEqualityThreshold = 1e-5

func isFloatEqual(a, b float32) bool {
	return math.Abs(float64(a)-float64(b)) <= floatEqualityThreshold
}
