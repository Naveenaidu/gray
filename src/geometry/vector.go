package geometry

import (
	"math"

	"github.com/Naveenaidu/gray/src/util"
)

type Vector struct {
	X, Y, Z float64
}

func NewVector(x float64, y float64, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v1 Vector) IsEqual(v2 Vector) bool {
	return util.IsFloatEqual(v1.X, v2.X) &&
		util.IsFloatEqual(v1.Y, v2.Y) &&
		util.IsFloatEqual(v1.Z, v2.Z)
}

func (v1 Vector) add(v2 Vector) *Vector {
	return &Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func AddVectors(vlist []Vector) *Vector {
	vectorSum := Vector{}
	for _, vec := range vlist {
		vectorSum = *vectorSum.add(vec)
	}
	return &vectorSum
}

func (v1 Vector) subtract(v2 Vector) *Vector {
	return &Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func SubtractVectors(vlist []Vector) *Vector {
	result := vlist[0]
	for i := 1; i < len(vlist); i++ {
		result = *result.subtract(vlist[i])
	}
	return &result
}

func (v1 Vector) Negate() *Vector {
	return &Vector{-v1.X, -v1.Y, -v1.Z}
}

func (v1 Vector) ScalarMultiply(scalar float64) *Vector {
	return &Vector{v1.X * scalar, v1.Y * scalar, v1.Z * scalar}
}

func (v1 Vector) ScalarDivide(scalar float64) *Vector {
	return &Vector{v1.X / scalar, v1.Y / scalar, v1.Z / scalar}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v1 Vector) Normalize() *Vector {
	v1_magnitude := v1.Magnitude()
	return &Vector{v1.X / v1_magnitude, v1.Y / v1_magnitude, v1.Z / v1_magnitude}
}

func (v1 Vector) DotProduct(v2 Vector) float64 {
	return (v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z)
}

func (v1 Vector) CrossProduct(v2 Vector) *Vector {
	return &Vector{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vector) ToMatrix() *Matrix {
	return NewMatrix(4, 1, [][]float64{{v1.X}, {v1.Y}, {v1.Z}, {0.0}})
}

func (v1 Vector) Scale(x float64, y float64, z float64) *Vector {
	scaleM := ScaleM(x, y, z)
	vectorM := v1.ToMatrix()
	scaledVectorM := scaleM.Multiply(*vectorM)
	return scaledVectorM.ToVector()
}
