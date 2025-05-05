package main

import (
	"errors"
	"fmt"
	"math"

	"rsc.io/quote"
)

type Tuple struct {
	x, y, z, w float64
}

func (t1 Tuple) isEqual(t2 Tuple) bool {
	return isFloatEqual(t1.x, t2.x) &&
		isFloatEqual(t1.y, t2.y) &&
		isFloatEqual(t1.z, t2.z) &&
		isFloatEqual(t1.w, t2.w)
}

func NewPoint(x float64, y float64, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func NewVector(x float64, y float64, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func (t1 Tuple) isPoint() bool {
	return t1.w == 1.0
}

func (t1 Tuple) isVector() bool {
	return t1.w == 0.0
}

// t1 +  t2
func (t1 Tuple) add(t2 Tuple) (*Tuple, error) {
	// Adding two Points do not make sense. Points can only be added to a vector
	if t1.isPoint() && t2.isPoint() {
		return nil, errors.New("two points cannot be added")
	}
	return &Tuple{t1.x + t2.x, t1.y + t2.y, t1.z + t2.z, t1.w + t2.w}, nil
}

// t1 - t2
func (t1 Tuple) subtract(t2 Tuple) (*Tuple, error) {
	// Subtracting a point from a vector is not a standard vector operation
	if t1.isVector() && t2.isPoint() {
		return nil, errors.New("subtracting a point from a vector is not allowed")
	}
	return &Tuple{t1.x - t2.x, t1.y - t2.y, t1.z - t2.z, t1.w - t2.w}, nil
}

func (t1 Tuple) negate() *Tuple {
	// For points and vectors, negate the x, y, z coordinates
	if t1.isPoint() || t1.isVector() {
		return &Tuple{-t1.x, -t1.y, -t1.z, t1.w}
	} else {
		return &Tuple{-t1.x, -t1.y, -t1.z, -t1.w}
	}

}

func (t1 Tuple) scalarMultiply(scalar float64) *Tuple {
	// For points and vectors, multiply the x, y, z coordinates by the scalar
	if t1.isPoint() || t1.isVector() {
		return &Tuple{t1.x * scalar, t1.y * scalar, t1.z * scalar, t1.w}
	} else {
		// For other tuples, multiply all coordinates by the scalar
		return &Tuple{t1.x * scalar, t1.y * scalar, t1.z * scalar, t1.w * scalar}
	}
}

func (t1 Tuple) scalarDivide(scalar float64) *Tuple {
	// For points and vectors, divide the x, y, z coordinates by the scalar
	if t1.isPoint() || t1.isVector() {
		return &Tuple{t1.x / scalar, t1.y / scalar, t1.z / scalar, t1.w}
	} else {
		// For other tuples, divide all coordinates by the scalar
		return &Tuple{t1.x / scalar, t1.y / scalar, t1.z / scalar, t1.w / scalar}
	}
}

func (t1 Tuple) magnitude() float64 {

	// "magnitude" oeration is only allowed for vectors
	// We could have returned an error here but doing so would make
	// usage of this function cumbersome. So we return NaN instead.
	// ASK: Is this a good idea?
	if t1.isPoint() {
		return math.NaN()
	}
	return math.Sqrt(t1.x*t1.x + t1.y*t1.y + t1.z*t1.z + t1.w*t1.w)
}

// Converts an arbitrary vector into a unit vector.
// This keep calculations anchored relative to a common scale (the unit vector)
func (t1 Tuple) normalize() *Tuple {
	// "normalize" oeration is only allowed for vectors
	// We could have returned an error here but doing so would make
	// usage of this function cumbersome. So we return NaN instead.
	// ASK: Is this a good idea?
	if t1.isPoint() {
		return &Tuple{math.NaN(), math.NaN(), math.NaN(), 1}
	}

	t1_magnitude := t1.magnitude()
	normalized_x := t1.x / t1_magnitude
	normalized_y := t1.y / t1_magnitude
	normalized_z := t1.z / t1_magnitude
	normalized_w := t1.w / t1_magnitude

	return &Tuple{normalized_x, normalized_y, normalized_z, normalized_w}
}

func main() {
	fmt.Println(quote.Glass())
}
