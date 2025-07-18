package world

import "github.com/Naveenaidu/gray/src/util"

type Color struct {
	r, g, b float64
}

var Red = NewColor(1.0, 0.0, 0.0)
var Green = NewColor(0.0, 1.0, 0.0)
var Blue = NewColor(0.0, 0.0, 1.0)
var Black = NewColor(0.0, 0.0, 0.0)

func NewColor(r float64, g float64, b float64) *Color {
	return &Color{r, g, b}
}

func (c1 Color) IsEqual(c2 Color) bool {
	return util.IsFloatEqual(c1.r, c2.r) &&
		util.IsFloatEqual(c1.g, c2.g) &&
		util.IsFloatEqual(c1.b, c2.b)
}

func (c *Color) clamp() *Color {
	c.r = util.Clamp(c.r)
	c.g = util.Clamp(c.g)
	c.b = util.Clamp(c.b)
	return c
}

func (c1 Color) add(c2 Color) *Color {
	return &Color{c1.r + c2.r, c1.g + c2.g, c1.b + c2.b}
}

func AddColors(vlist []Color) *Color {
	ColorSum := Color{}
	for _, vec := range vlist {
		ColorSum = *ColorSum.add(vec)
	}
	return &ColorSum
}

func (c1 Color) subtract(c2 Color) *Color {
	return &Color{c1.r - c2.r, c1.g - c2.g, c1.b - c2.b}
}

func SubtractColors(vlist []Color) *Color {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.subtract(vlist[i])
	}
	return &result
}

func (c1 Color) multiply(c2 Color) *Color {
	return &Color{c1.r * c2.r, c1.g * c2.g, c1.b * c2.b}
}

func MultiplyColors(vlist []Color) *Color {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.multiply(vlist[i])
	}
	return &result
}
