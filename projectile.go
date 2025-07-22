package main

import (
	"fmt"
	"math"

	"github.com/Naveenaidu/gray/src/geom"
	"github.com/Naveenaidu/gray/src/world"
)

type Projectile struct {
	position geom.Point
	velocity geom.Vector
}

type Environment struct {
	gravity geom.Vector
	wind    geom.Vector
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.AddVector(proj.velocity)
	velocity := geom.AddVectors([]geom.Vector{proj.velocity, env.gravity, env.wind})
	return Projectile{*position, *velocity}
}

func projection(env Environment, proj Projectile, canvas *world.Canvas) Projectile {
	if proj.position.Y <= 0 {
		return proj
	}

	newproj := tick(env, proj)
	// Plot each projection on the canvas
	// The y coordinate is subtracted with canvas height to match our canvas coordinates
	proj_canvas_y := canvas.Height - int(newproj.position.Y)
	proj_canvas_x := int(newproj.position.X)

	// exit early if the projectile is going out of bounds of canvas
	if proj_canvas_y < 0 || proj_canvas_y >= canvas.Height || (proj_canvas_x >= canvas.Width) {
		return proj
	}

	fmt.Printf("Projectile Position: %v+\n", newproj.position)
	fmt.Printf("Projectile canvas Position: %d, %d\n", proj_canvas_x, proj_canvas_y)

	canvas.WritePixel(proj_canvas_x, proj_canvas_y, *world.Red)

	return projection(env, newproj, canvas)
}

func ThrowProjectile(env Environment, proj Projectile, canvas *world.Canvas) Projectile {
	return projection(env, proj, canvas)
}

// ---- Projectile Main ------
// func main() {
// 	// fmt.Println(quote.Glass())

// 	// start := Point{0, 1, 0}
// 	// velocity := Vector{1, 1.8, 0}.Normalize().ScalarMultiply(11.25)
// 	// gravity := Vector{0, -0.1, 0}
// 	// wind := Vector{-0.01, 0, 0}

// 	// canvas := NewCanvas(900, 500, *Black)

// 	// // Projectile
// 	// p := Projectile{position: start, velocity: *velocity}
// 	// e := Environment{gravity: gravity, wind: wind}
// 	// finalProjection := ThrowProjectile(e, p, canvas)
// 	// canvas.WriteToPPM("projectile.ppm")
// 	// fmt.Printf("\nfinal projectile position: %v+\n", finalProjection.position)

// }

func drawClock(radius float64, canvas *world.Canvas) {
	// Get the center
	centerX := canvas.Width / 2
	centerY := canvas.Height / 2
	center := geom.NewPoint(float64(centerX), float64(centerY), 0.0)
	fmt.Printf("\nCenter point %v+\n", center)
	canvas.WritePixel(int(center.X), int(center.Y), *world.Green)

	/*
		Learning: To draw a picture of your desire, you compute your drawing
		from your imagined origin taking unit points as reference and then
		compute the points to match your canvas dimensions
	*/
	refPointM := geom.NewPoint(0, 1, 0).ToMatrix()
	hourRotateM := geom.RotateZM(math.Pi / 6)
	for h := 0; h <= 11; h++ {
		// Rotate the previous hour point by pi/6
		nextHourPointM := geom.ChainTransforms([]*geom.Matrix{refPointM, hourRotateM})
		nextHourPoint := nextHourPointM.ToPoint()
		fmt.Printf("\n %d hour point %v", h+1, nextHourPoint)
		// canvas.WritePixel(int(nextHourPoint.X), int(nextHourPoint.Y), *Red)

		// Convert these points to our frame
		nextHourPointX := (nextHourPoint.X * radius) + float64(centerX)
		nextHourPointY := (nextHourPoint.Y * radius) + float64(centerY)
		fmt.Printf("\n next hour point %v %v\n", nextHourPointX, nextHourPointY)

		canvas.WritePixel(int(nextHourPointX), int(nextHourPointY), *world.Red)

		refPointM = nextHourPointM
	}

}

func main() {
	canvas := world.NewCanvas(100, 100, *world.Black)
	drawClock(15, canvas)
	canvas.WriteToPPM("clock.ppm")
}
