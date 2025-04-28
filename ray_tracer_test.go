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

	if !isPoint(a) {
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

	if !isVector(a) {
		t.Errorf("got: Point, want: Vector")
	}
}

func TestCreatePoint(t *testing.T) {
	point := point(4, -4, 3)
	if !isPoint(point) {
		t.Errorf("expected point, got something else")
	}
}

func TestCreateVector(t *testing.T) {
	vector := vector(4, -4, 3)
	if !isVector(vector) {
		t.Errorf("expected vector, got something else")
	}
}
