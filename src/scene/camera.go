package scene

import (
	"math"

	"github.com/Naveenaidu/gray/src/core/color"
	coreMath "github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/rayt"
	"github.com/Naveenaidu/gray/src/rendering"
)

// transformation matrix that orients the world relative to the eye
/*
- from: you specify where you want the eye to be in the scene
- to: the point in the scene at which you want to look
- up: the vector indication which direction is up
*/
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

	var aspect, halfWidth, halfHeight float64
	// calculate the pixel size
	halfView := math.Tan(fieldOfView / 2)
	aspect = float64(hsize) / float64(vsize)

	if aspect >= 1.0 {
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

// Compute the world cooridates at the center of given pixel
func RayForPixel(camera Camera, px int, py int) *rayt.Ray {
	// offset from edge of canvas to the pixel center
	xOffset := (float64(px) + 0.5) * camera.PixelSize
	yOffset := (float64(py) + 0.5) * camera.PixelSize

	// untransformed coordinates of pixel in the worl space
	// (camera looks towards -z, so +x is to the left) i.e
	//		Positive X is to the left. and Positive Y is to the upwards
	worldX := camera.HalfWidth - xOffset
	worldY := camera.halfHeight - yOffset

	// The "transform" field of camera tells, how the camera looks at the world,
	// but we need the inverse, i.e how is the world looking at camera.
	// Transform the coordinates from camera to world.
	// (note camera is at z=-1)

	// pixel ← inverse(camera.transform) * point(world_x, world_y, -1)
	untransformedPixel := coreMath.NewPoint(worldX, worldY, -1)
	pixel := camera.Transform.Inverse().Multiply(*untransformedPixel.ToMatrix()).ToPoint()

	// origin ← inverse(camera.transform) * point(0, 0, 0)
	origin := camera.Transform.Inverse().Multiply(*coreMath.ObjectOrigin().ToMatrix()).ToPoint()

	direction := pixel.Subtract(*origin).Normalize()

	return &rayt.Ray{Origin: *origin, Direction: *direction}
}

func Render(camera Camera, world World) *rendering.Canvas {
	image := rendering.NewCanvas(camera.Hsize, camera.Vsize, *color.Black)

	for y := 0; y < camera.Vsize; y++ {
		for x := 0; x < camera.Hsize; x++ {
			ray := RayForPixel(camera, x, y)
			color := ColorAt(world, *ray)
			image.WritePixel(x, y, color)
		}
	}

	return image
}
