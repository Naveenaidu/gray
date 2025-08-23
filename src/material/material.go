package material

import (
	color "github.com/Naveenaidu/gray/src/core/color"
)

type Material struct {
	Color     color.Color
	Ambient   float64 // ranges between 0 and 1
	Diffuse   float64 // ranges between 0 and 1
	Specular  float64 // ranges between 0 and 1
	Shininess int     // ranges between 10 and 200
}

func DefaultMaterial() Material {
	return Material{
		Color:     *color.NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
