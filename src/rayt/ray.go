package rayt

import (
	"math"

	"github.com/Naveenaidu/gray/src/geom"
)

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

func (r Ray) IntersectSphere(s geom.Sphere) []float64 {
	intersections := []float64{}

	// We assume the spehre is at origin
	// vector from sphere origin, to the ray origin
	sphereToRay := (r.Origin).Subtract(s.Center)

	a := r.Direction.DotProduct(r.Direction)
	b := 2 * (r.Direction.DotProduct(*sphereToRay))
	c := sphereToRay.DotProduct(*sphereToRay) - math.Pow(s.Radius, 2)

	discriminant := math.Pow(b, 2) - 4*a*c

	// ray only intersects sphere if the discriminant is greater than zero
	if discriminant >= 0 {
		t1 := (-1*b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-1*b + math.Sqrt(discriminant)) / (2 * a)

		intersections = make([]float64, 2)
		intersections[0] = t1
		intersections[1] = t2
	}

	return intersections
}
