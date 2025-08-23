package color

import (
	core "github.com/Naveenaidu/gray/src/core/math"
)

type Color struct {
	R, G, B float64
}

var Red = NewColor(1.0, 0.0, 0.0)
var Green = NewColor(0.0, 1.0, 0.0)
var Blue = NewColor(0.0, 0.0, 1.0)
var Black = NewColor(0.0, 0.0, 0.0)

func NewColor(r float64, g float64, b float64) *Color {
	return &Color{r, g, b}
}

func (c1 Color) IsEqual(c2 Color) bool {
	return core.IsFloatEqual(c1.R, c2.R) &&
		core.IsFloatEqual(c1.G, c2.G) &&
		core.IsFloatEqual(c1.B, c2.B)
}

func (c *Color) Clamp() *Color {
	c.R = core.Clamp(c.R)
	c.G = core.Clamp(c.G)
	c.B = core.Clamp(c.B)
	return c
}

func (c1 Color) add(c2 Color) *Color {
	return &Color{c1.R + c2.R, c1.G + c2.G, c1.B + c2.B}
}

func AddColors(vlist []Color) *Color {
	ColorSum := Color{}
	for _, vec := range vlist {
		ColorSum = *ColorSum.add(vec)
	}
	return &ColorSum
}

func (c1 Color) subtract(c2 Color) *Color {
	return &Color{c1.R - c2.R, c1.G - c2.G, c1.B - c2.B}
}

func SubtractColors(vlist []Color) *Color {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.subtract(vlist[i])
	}
	return &result
}

func (c1 Color) multiply(c2 Color) *Color {
	return &Color{c1.R * c2.R, c1.G * c2.G, c1.B * c2.B}
}

func MultiplyColors(vlist []Color) *Color {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.multiply(vlist[i])
	}
	return &result
}

func (c1 Color) ScalarMultiply(scalar float64) *Color {
	return &Color{c1.R * scalar, c1.G * scalar, c1.B * scalar}
}
