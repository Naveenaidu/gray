package geom

import "github.com/Naveenaidu/gray/src/util"

type Point struct {
	X, Y, Z float64
}

func NewPoint(x float64, y float64, z float64) *Point {
	return &Point{x, y, z}
}

func (p1 Point) IsEqual(p2 Point) bool {
	return util.IsFloatEqual(p1.X, p2.X) &&
		util.IsFloatEqual(p1.Y, p2.Y) &&
		util.IsFloatEqual(p1.Z, p2.Z)
}

func (p1 Point) AddVector(v1 Vector) *Point {
	return &Point{p1.X + v1.X, p1.Y + v1.Y, p1.Z + v1.Z}
}

func (p1 Point) Subtract(p2 Point) *Vector {
	return &Vector{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

func (p1 Point) SubtractVector(v1 Vector) *Point {
	return &Point{p1.X - v1.X, p1.Y - v1.Y, p1.Z - v1.Z}
}

func (p1 Point) Negate() *Point {
	return &Point{-p1.X, -p1.Y, -p1.Z}
}

func (p1 Point) ScalarMultiply(scalar float64) *Point {
	return &Point{p1.X * scalar, p1.Y * scalar, p1.Z * scalar}
}

func (p1 Point) ScalarDivide(scalar float64) *Point {
	return &Point{p1.X / scalar, p1.Y / scalar, p1.Z / scalar}
}

func (p1 Point) ToMatrix() *Matrix {
	return NewMatrix(4, 1, [][]float64{{p1.X}, {p1.Y}, {p1.Z}, {1.0}})
}

func (p1 Point) Translate(x float64, y float64, z float64) *Point {
	translationM := TranslationM(x, y, z)
	pointM := p1.ToMatrix()
	translatedPointM := translationM.Multiply(*pointM)
	return translatedPointM.ToPoint()
}

func (p1 Point) Scale(x float64, y float64, z float64) *Point {
	scaleM := ScaleM(x, y, z)
	pointM := p1.ToMatrix()
	scaledPointM := scaleM.Multiply(*pointM)
	return scaledPointM.ToPoint()
}

func (p1 Point) Shear(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) *Point {
	shearM := ShearM(xy, xz, yx, yz, zx, zy)
	pointM := p1.ToMatrix()
	shearedPointM := shearM.Multiply(*pointM)
	return shearedPointM.ToPoint()
}
