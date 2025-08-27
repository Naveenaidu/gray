package scene

import (
	"github.com/Naveenaidu/gray/src/core/color"
	"github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/lighting"
	"github.com/Naveenaidu/gray/src/material"
	"github.com/Naveenaidu/gray/src/shape"
)

type World struct {
	Light   lighting.Light
	Spheres []shape.Sphere
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
