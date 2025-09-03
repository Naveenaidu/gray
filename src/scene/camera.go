package scene

import "github.com/Naveenaidu/gray/src/core/math"

// transformation matrix that orients the world relative to the eye
func ViewTransform(from math.Point, to math.Point, up math.Vector) *math.Matrix {
	forward := to.Subtract(from).Normalize()
	upNormalized := up.Normalize()
	left := forward.CrossProduct(*upNormalized)
	trueUp := left.CrossProduct(*forward)

	orientation := math.NewMatrix(4, 4, [][]float64{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})

	// move the scene into place before orienting it
	translationM := math.TranslationM(-from.X, -from.Y, -from.Z)
	return orientation.Multiply(*translationM)
}
