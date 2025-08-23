package math

import "math"

/*
Create a translation matrix. Example matrix:

	1.00   0.00   0.00   X
	0.00   1.00   0.00   Y
	0.00   0.00   1.00   Z
	0.00   0.00   0.00   1.00

Points and vectors are represented as single colum matrix, something like:

	PX.00
	PY.00
	PZ.00
	1.00

Muliplying translation matrix with point matrix gives: P' = T * P, resembles
the below operation:

	PX' = X' + PX
	PY' = Y' + PY
	PZ' = Z' + PZ

Note: Translations to vector are not supported. Vector is an arrow, moving it
around space does not change the direction it points to
*/
func TranslationM(x float64, y float64, z float64) *Matrix {
	return NewMatrix(4, 4, [][]float64{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	})
}

/*
Scaling Matrix:

# Translation moves a point by adding to it, scaling moves it by multiplication

Muliplying scaling matrix with point/vector matrix gives: P' = T * P, resembles

	PX' = X' * PX
	PY' = Y' * PY
	PZ' = Z' * PZ
*/
func ScaleM(x float64, y float64, z float64) *Matrix {
	return NewMatrix(4, 4, [][]float64{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	})
}

func RotateXM(r float64) *Matrix {
	return NewMatrix(4, 4, [][]float64{
		{1, 0, 0, 0},
		{0, math.Cos(r), (-1 * math.Sin(r)), 0},
		{0, math.Sin(r), math.Cos(r), 0},
		{0, 0, 0, 1},
	})
}

func RotateYM(r float64) *Matrix {
	return NewMatrix(4, 4, [][]float64{
		{math.Cos(r), 0, math.Sin(r), 0},
		{0, 1, 0, 0},
		{(-1 * math.Sin(r)), 0, math.Cos(r), 0},
		{0, 0, 0, 1},
	})
}

func RotateZM(r float64) *Matrix {
	return NewMatrix(4, 4, [][]float64{
		{math.Cos(r), (-1 * math.Sin(r)), 0, 0},
		{math.Sin(r), math.Cos(r), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

func ShearM(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) *Matrix {
	return NewMatrix(4, 4, [][]float64{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	})
}

func ChainTransforms(transformations []*Matrix) *Matrix {
	chainTransformM := IdentityMatrix()

	for t := len(transformations) - 1; t >= 0; t-- {
		chainTransformM = chainTransformM.Multiply(*transformations[t])
	}

	return chainTransformM
}
