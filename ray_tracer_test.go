package main

import "testing"

func TestTupleIsPoint(t *testing.T) {
	a := Tuple{x: 4.3, y: -4.2, z: 3.1, w: 1.0}
	if a.x != 4.3 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.x, 4.3)
	}
	if a.y != -4.2 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.y, -4.2)
	}
	if a.z != 3.1 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.z, 3.1)
	}
	if a.w != 1.0 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.w, 1.0)
	}

	if !a.isPoint() {
		t.Errorf("got: Vector, want: Point")
	}
}

func TestTupleIsVector(t *testing.T) {
	a := Tuple{x: 4.3, y: -4.2, z: 3.1, w: 0.0}
	if a.x != 4.3 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.x, 4.3)
	}
	if a.y != -4.2 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.y, -4.2)
	}
	if a.z != 3.1 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.z, 3.1)
	}
	if a.w != 0.0 {
		t.Errorf("Coordinate value was incorrect, got: %f, want: %f", a.w, 0.0)
	}

	if !a.isVector() {
		t.Errorf("got: Point, want: Vector")
	}
}

func TestCreatePoint(t *testing.T) {
	point := NewPoint(4, -4, 3)
	if !point.isPoint() {
		t.Errorf("expected point, got something else")
	}
}

func TestCreateVector(t *testing.T) {
	vector := NewVector(4, -4, 3)
	if !vector.isVector() {
		t.Errorf("expected vector, got something else")
	}
}

func TestAddTuples(t *testing.T) {
	// Case: Adding two points, gives a new point
	a1 := NewPoint(3, -2, 5)
	a2 := NewVector(-2, 3, 1)

	a3, err := a1.add(*a2)
	if err != nil {
		t.Errorf("Error while adding tuples: %v", err)
	}

	if !a3.isEqual(*NewPoint(1, 1, 6)) {
		t.Errorf("got: %+v, want: (1, 1, 6, 1)", a3)
	}
}

func TestSubtractTuples(t *testing.T) {
	var p1, p2, v1 *Tuple
	// Case 1: Subtract two points, gives a new vector
	p1 = NewPoint(3, 2, 1)
	p2 = NewPoint(5, 6, 7)
	v1, err := p1.subtract(*p2)
	if err != nil {
		t.Errorf("Error while subtracting tuples: %v", err)
	}

	if !v1.isEqual(*NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", v1)
	}

	// Case 2: Subtract a vector from a point, gives a new point
	// This is equivalent to walking from a point in the direction of the vector
	p1 = NewPoint(3, 2, 1)
	v1 = NewVector(5, 6, 7)
	p2, err = p1.subtract(*v1)
	if err != nil {
		t.Errorf("Error while subtracting tuples: %v", err)
	}
	if !p2.isEqual(*NewPoint(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 1)", p2)
	}

	// Case 3: Subtract a point from a vector, gives an error
	v1 = NewVector(3, 2, 1)
	p2 = NewPoint(5, 6, 7)
	_, err = v1.subtract(*p2)
	if err == nil {
		t.Errorf("Expected error while subtracting point from vector, got nil")
	}

	// Case 4: Subtract two vectors, gives a new vector
	v1 = NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	v3, err := v1.subtract(*v2)
	if err != nil {
		t.Errorf("Error while subtracting tuples: %v", err)
	}
	if !v3.isEqual(*NewVector(-2, -4, -6)) {
		t.Errorf("got: %+v, want: (-2, -4, -6, 0)", v3)
	}

	// Case 5: Subtract a vector from itself, gives a zero vector
	v1 = NewVector(3, 2, 1)
	v2 = NewVector(3, 2, 1)
	v3, err = v1.subtract(*v2)
	if err != nil {
		t.Errorf("Error while subtracting tuples: %v", err)
	}
	if !v3.isEqual(*NewVector(0, 0, 0)) {
		t.Errorf("got: %+v, want: (0, 0, 0, 0)", v3)
	}

}

func TestNegateVector(t *testing.T) {
	a1 := NewVector(1, -2, 3)

	if !a1.negate().isEqual(*NewVector(-1, 2, -3)) {
		t.Errorf("got: %+v, want: (-1, 2, -3, 0)", a1.negate())
	}
}

func TestScalarMuliply(t *testing.T) {
	var a1, a2, a3, a4 *Tuple
	a1 = NewPoint(1, -2, 3)
	a2 = a1.scalarMultiply(3)
	if !a2.isEqual(*NewPoint(3, -6, 9)) {
		t.Errorf("got: %+v, want: (3, -6, 9, 1)", a2)
	}

	a3 = &Tuple{x: 1, y: -2, z: 3, w: 3}
	a4 = a3.scalarMultiply(0.5)
	if !a4.isEqual(Tuple{0.5, -1, 1.5, 1.5}) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 3)", a4)
	}
}

func TestScalarDivide(t *testing.T) {
	var a1, a2, a3, a4 *Tuple
	a1 = NewPoint(1, -2, 3)
	a2 = a1.scalarDivide(2)
	if !a2.isEqual(*NewPoint(0.5, -1, 1.5)) {
		t.Errorf("got: %+v, want: (0.5, -1, 1.5, 1)", a2)
	}

	a3 = &Tuple{x: 1, y: -2, z: 3, w: 3}
	a4 = a3.scalarDivide(0.5)
	if !a4.isEqual(Tuple{2, -4, 6, 6}) {
		t.Errorf("got: %+v, want: (2, -4, 6, 3)", a4)
	}
}
