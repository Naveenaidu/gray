package rayt

import "github.com/Naveenaidu/gray/src/geom"

type Ray struct {
	Origin    geom.Point
	Direction geom.Vector
}

func (r Ray) Position(t float64) *geom.Point {
	// newPosition = r.origin + r.direction * t
	// The new point that lies at the distance "t" along the ray
	newPosition := r.Origin.AddVector(*r.Direction.ScalarMultiply(t))
	return newPosition
}
