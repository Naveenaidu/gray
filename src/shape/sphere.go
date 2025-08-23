package shape

import (
	core "github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/material"
)

type Sphere struct {
	Center    core.Point
	Radius    float64
	Transform core.Matrix
	Material  material.Material
}

// Sphere with radius 1 and centered at origin (0,0,0)
func UnitSphere() *Sphere {
	return &Sphere{
		Center:    *core.NewPoint(0, 0, 0),
		Radius:    1.0,
		Transform: *core.IdentityMatrix(),
		Material:  material.DefaultMaterial(),
	}
}
