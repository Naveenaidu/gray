package scene

import (
	"sort"

	"github.com/Naveenaidu/gray/src/core/color"
	"github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/lighting"
	"github.com/Naveenaidu/gray/src/material"
	"github.com/Naveenaidu/gray/src/rayt"
	"github.com/Naveenaidu/gray/src/shape"
)

type World struct {
	Light   lighting.Light
	Spheres []shape.Sphere
}

type Computation struct {
	T        float64
	Object   shape.Sphere
	Point    math.Point
	EyeV     math.Vector
	NormalV  math.Vector
	Inside   bool
	InShadow bool
}

func DefaultWorld() *World {
	pointLight := lighting.NewLight(*color.NewColor(1, 1, 1), *math.NewPoint(-10, 10, -10))

	s1 := shape.UnitSphere()
	s1.Material = material.Material{
		Color:     *color.NewColor(0.8, 1.0, 0.6),
		Ambient:   material.DefaultMaterial().Ambient,
		Diffuse:   0.7,
		Specular:  0.2,
		Shininess: material.DefaultMaterial().Shininess,
	}

	s2 := shape.UnitSphere()
	s2.Transform = *math.ChainTransforms([]*math.Matrix{math.ScaleM(0.5, 0.5, 0.5)})

	// Two concentric spheres, where the outermost is a unit sphere and the
	// innermost has a radius of 0.5
	spheres := []shape.Sphere{*s1, *s2}

	return &World{Light: pointLight, Spheres: spheres}
}

func IntersectWorld(world World, ray rayt.Ray) []rayt.Intersection {
	xs := []rayt.Intersection{}

	for _, s := range world.Spheres {
		sIntersections := ray.IntersectSphere(s)
		xs = append(xs, sIntersections...)
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})

	return xs

}

func PrepareComputations(intersection rayt.Intersection, ray rayt.Ray) *Computation {
	point := ray.Position(intersection.T)
	eyev := ray.Direction.Negate()
	normalV := lighting.NormalAt(intersection.Object, *point)
	inside := false
	inShadow := false

	// check if the ray is originating from inside of the sphere. If the eye
	// vector and the normal vector are in opposite direction than the ray is
	// originating from inside of sphere
	if normalV.DotProduct(*eyev) < 0 {
		inside = true
		normalV = *normalV.Negate()
	}

	return &Computation{
		T:        intersection.T,
		Object:   intersection.Object,
		Point:    *point,
		EyeV:     *eyev,
		NormalV:  normalV,
		Inside:   inside,
		InShadow: inShadow,
	}
}

func ShadeHit(world World, comps Computation) color.Color {
	return lighting.Lighting(comps.Object.Material, world.Light, comps.Point, comps.EyeV, comps.NormalV, comps.InShadow)
}

func ColorAt(world World, ray rayt.Ray) color.Color {
	color := color.Black
	intrs := IntersectWorld(world, ray)

	hit := ray.Hit(intrs)
	if hit == nil {
		return *color
	}
	comps := PrepareComputations(*hit, ray)
	hitColor := ShadeHit(world, *comps)

	return hitColor

}

func IsShadowed(world World, point math.Point) bool {
	v := world.Light.Position.Subtract(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	// Shadow ray (light - point)
	shadowRay := rayt.Ray{Origin: point, Direction: *direction}
	// shadow ray and the intersection of that ray with world
	intersections := IntersectWorld(world, shadowRay)

	h := shadowRay.Hit(intersections)
	if h != nil && h.T < distance {
		return true
	}
	return false

}
