package scene

import (
	"math"

	coreMath "github.com/Naveenaidu/gray/src/core/math"
)

// transformation matrix that orients the world relative to the eye
func ViewTransform(from coreMath.Point, to coreMath.Point, up coreMath.Vector) *coreMath.Matrix {
	forward := to.Subtract(from).Normalize()
	upNormalized := up.Normalize()
	left := forward.CrossProduct(*upNormalized)
	trueUp := left.CrossProduct(*forward)

	orientation := coreMath.NewMatrix(4, 4, [][]float64{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})

	// move the scene into place before orienting it
	translationM := coreMath.TranslationM(-from.X, -from.Y, -from.Z)
	return orientation.Multiply(*translationM)
}

type Camera struct {
	// horizontal size (in pixels) of the canvas that the picture will be rendered to
	Hsize int
	// canvas vertical size (in pixels)
	Vsize int
	// angle that describes how much the camera can see
	FieldOfView float64
	// matrix representing how the world should be oriented relative to the camera
	Transform coreMath.Matrix
	// size of the pixel on the canvas
	PixelSize  float64
	HalfWidth  float64
	halfHeight float64
}

func NewCamera(hsize int, vsize int, fieldOfView float64) *Camera {

	// calculate the pixel size
	halfView := math.Tan(fieldOfView / 2)
	aspect := hsize / vsize

	var halfWidth, halfHeight float64

	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / float64(aspect)
	} else {
		halfWidth = halfView * float64(aspect)
		halfHeight = halfView
	}

	pixelSize := halfWidth * 2 / float64(hsize)

	return &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fieldOfView,
		Transform:   *coreMath.IdentityMatrix(),
		PixelSize:   pixelSize,
		HalfWidth:   halfWidth,
		halfHeight:  halfHeight,
	}
}
