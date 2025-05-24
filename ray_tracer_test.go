package main

import (
	"math"
	"testing"
)

func TestAddTuples(t *testing.T) {
	// Case 1: Adding vector to point gives a new point
	a1 := NewPoint(3, -2, 5)
	a2 := NewVector(-2, 3, 1)
	a3 := a1.AddVector(*a2)

	if !a3.IsEqual(*NewPoint(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 1)", a3)
	}

	// Case 2: Adding two vectors gives a new vector
	c2v1 := NewVector(3, -2, 5)
	c2v2 := NewVector(-2, 3, 1)
	c2v3 := c2v1.Add(*c2v2)
	if !c2v3.IsEqual(*NewVector(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 0)", c2v3)
	}

}

func TestSubtractTuples(t *testing.T) {

	// Case 1: Subtract two points, gives a new vector
	c1p1 := NewPoint(3, 2, 1)
	c1p2 := NewPoint(5, 6, 7)
	c1v1 := c1p1.Subtract(*c1p2)

	if !c1v1.IsEqual(*NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", c1v1)
	}

	// Case 2: Subtract a vector from a point, gives a new point
	// This is equivalent to walking from a point in the direction of the vector
	c2p1 := NewPoint(3, 2, 1)
	c2v1 := NewVector(5, 6, 7)
	c2p2 := c2p1.SubtractVector(*c2v1)
	if !c2p2.IsEqual(*NewPoint(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 1)", c2p2)
	}

	// Case 3: Subtract two vectors, gives a new vector
	c3v1 := NewVector(3, 2, 1)
	c3v2 := NewVector(5, 6, 7)
	c3v3 := c3v1.Subtract(*c3v2)
	if !c3v3.IsEqual(*NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", c3v3)
	}

	// Case 4: Subtract a vector from itself, gives a zero vector
	c4v1 := NewVector(3, 2, 1)
	c4v2 := c4v1.Subtract(*c4v1)
	if !c4v2.IsEqual(*NewVector(0, 0, 0)) {
		t.Errorf("got: %+v, want: (0, 0, 0, 0)", c4v2)
	}

}

func TestNegateVector(t *testing.T) {
	a1 := NewVector(1, -2, 3)

	if !a1.Negate().IsEqual(*NewVector(-1, 2, -3)) {
		t.Errorf("got: %+v, want: (-1, 2, -3, 0)", a1.Negate())
	}
}

func TestScalarMuliply(t *testing.T) {
	p1 := NewPoint(1, -2, 3)
	p2 := p1.ScalarMultiply(3)
	if !p2.IsEqual(*NewPoint(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 1)", p2)
	}
	v1 := NewVector(1, -2, 3)
	v2 := v1.ScalarMultiply(3)
	if !v2.IsEqual(*NewVector(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 0)", v2)
	}

}

func TestScalarDivide(t *testing.T) {
	p1 := NewPoint(1, -2, 3)
	p2 := p1.ScalarDivide(2)
	if !p2.IsEqual(*NewPoint(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 1)", p2)
	}
	v1 := NewVector(1, -2, 3)
	v2 := v1.ScalarDivide(2)
	if !v2.IsEqual(*NewVector(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 0)", v2)
	}

}

func TestVectorMagnitude(t *testing.T) {

	v := NewVector(1, 0, 0)
	if !isFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = NewVector(0, 1, 0)
	if !isFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = NewVector(0, 0, 1)
	if !isFloatEqual(v.Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Magnitude())
	}

	v = NewVector(1, 2, 3)
	if !isFloatEqual(v.Magnitude(), math.Sqrt(14)) {
		t.Errorf("got: %f, want: %f", v.Magnitude(), math.Sqrt(14))
	}

	v = NewVector(-1, -2, -3)
	if !isFloatEqual(v.Magnitude(), math.Sqrt(14)) {
		t.Errorf("got: %f, want: %f", v.Magnitude(), math.Sqrt(14))
	}

}

func TestNormalizeVector(t *testing.T) {
	var v, v_normalized *Vector

	v = NewVector(4, 0, 0)
	v_normalized = v.Normalize()
	if !isFloatEqual(v_normalized.x, 1.0) {
		t.Errorf("got: %f, want: 1.0", v_normalized.x)
	}
	if !isFloatEqual(v_normalized.y, 0.0) {
		t.Errorf("got: %f, want: 0.0", v_normalized.y)
	}
	if !isFloatEqual(v_normalized.z, 0.0) {
		t.Errorf("got: %f, want: 0.0", v_normalized.z)
	}

	v = NewVector(1, 2, 3)
	v_normalized = v.Normalize()
	if !isFloatEqual(v_normalized.x, 0.26726) {
		t.Errorf("got: %f, want: 0.26726", v_normalized.x)
	}
	if !isFloatEqual(v_normalized.y, 0.53452) {
		t.Errorf("got: %f, want: 0.53452", v_normalized.y)
	}
	if !isFloatEqual(v_normalized.z, 0.80178) {
		t.Errorf("got: %f, want: 0.80178", v_normalized.z)
	}

	// Check the magnitude of the normalized vector
	v = NewVector(1, 2, 3)
	if !isFloatEqual(v.Normalize().Magnitude(), 1.0) {
		t.Errorf("got: %f, want: 1.0", v.Normalize().Magnitude())
	}
}

func TestDotProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	dotProduct := a.DotProduct(*b)

	if dotProduct != 20.0 {
		t.Errorf("got: %f, want: 20.0", dotProduct)
	}

}

func TestCrossProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	abCrossProduct := a.CrossProduct(*b)

	if !abCrossProduct.IsEqual(*NewVector(-1, 2, -1)) {
		t.Errorf("got: %+v, want: (-1, 2, -1)", abCrossProduct)
	}

	baCrossProduct := b.CrossProduct(*a)

	if !baCrossProduct.IsEqual(*NewVector(1, -2, 1)) {
		t.Errorf("got: %+v, want: (-1, 2, -1)", baCrossProduct)
	}

}
