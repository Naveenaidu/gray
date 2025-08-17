package rayt

import (
	"math"
	"sort"

	"github.com/Naveenaidu/gray/src/geom"
	"github.com/Naveenaidu/gray/src/material"
)

type Ray struct {
	Origin    geom.Point
	Direction geom.Vector
}

// TODO: Should "Intersection" be part of Ray struct
type Intersection struct {
	T      float64
	Object material.Sphere
}

func NewIntersection(t float64, obj material.Sphere) Intersection {
	return Intersection{T: t, Object: obj}
}

func (r Ray) Position(t float64) *geom.Point {
	// newPosition = r.origin + r.direction * t
	// The new point that lies at the distance "t" along the ray
	newPosition := r.Origin.AddVector(*r.Direction.ScalarMultiply(t))
	return newPosition
}

func (r Ray) IntersectSphere(s material.Sphere) []Intersection {
	intersections := []Intersection{}

	//apply the inverse of the sphere trasnformation to  ray
	transformedRay := r.Transform(s.Transform.Inverse())

	// We assume the spehre is at origin
	// vector from sphere origin, to the ray origin
	sphereToRay := (transformedRay.Origin).Subtract(s.Center)

	a := transformedRay.Direction.DotProduct(transformedRay.Direction)
	b := 2 * (transformedRay.Direction.DotProduct(*sphereToRay))
	c := sphereToRay.DotProduct(*sphereToRay) - math.Pow(s.Radius, 2)

	discriminant := math.Pow(b, 2) - 4*a*c

	// ray only intersects sphere if the discriminant is greater than zero
	if discriminant >= 0 {
		t1 := (-1*b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-1*b + math.Sqrt(discriminant)) / (2 * a)

		intersections = make([]Intersection, 2)
		intersections[0] = NewIntersection(t1, s)
		intersections[1] = NewIntersection(t2, s)
	}

	return intersections
}

func (r Ray) Hit(intersections []Intersection) *Intersection {
	// If "t" is negative then it means the intersection happened behind the
	// origin of ray, ignore those intersections
	nonNegativeIntersections := []Intersection{}

	for i := range intersections {
		if intersections[i].T > 0 {
			nonNegativeIntersections = append(nonNegativeIntersections, intersections[i])
		}
	}

	if len(nonNegativeIntersections) == 0 {
		return nil
	}

	sort.Slice(nonNegativeIntersections, func(i, j int) bool {
		return nonNegativeIntersections[i].T < nonNegativeIntersections[j].T
	})

	return &nonNegativeIntersections[0]

}

func (r Ray) Transform(m *geom.Matrix) Ray {
	transformedRayPointM := geom.ChainTransforms([]*geom.Matrix{r.Origin.ToMatrix(), m})
	transformedDirectionM := geom.ChainTransforms([]*geom.Matrix{r.Direction.ToMatrix(), m})
	return Ray{
		Origin: *transformedRayPointM.ToPoint(),
		// Direction: r.Direction,
		Direction: *transformedDirectionM.ToVector(),
	}
}
