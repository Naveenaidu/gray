package main

import (
	"fmt"
	"math"

	color "github.com/Naveenaidu/gray/src/core/color"
	core "github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/lighting"
	"github.com/Naveenaidu/gray/src/material"
	"github.com/Naveenaidu/gray/src/rayt"
	"github.com/Naveenaidu/gray/src/rendering"
	"github.com/Naveenaidu/gray/src/shape"
)

type Projectile struct {
	position core.Point
	velocity core.Vector
}

type Environment struct {
	gravity core.Vector
	wind    core.Vector
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.AddVector(proj.velocity)
	velocity := core.AddVectors([]core.Vector{proj.velocity, env.gravity, env.wind})
	return Projectile{*position, *velocity}
}

func projection(env Environment, proj Projectile, canvas *rendering.Canvas) Projectile {
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

	canvas.WritePixel(proj_canvas_x, proj_canvas_y, *color.Red)

	return projection(env, newproj, canvas)
}

func ThrowProjectile(env Environment, proj Projectile, canvas *rendering.Canvas) Projectile {
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

func drawClock(radius float64, canvas *rendering.Canvas) {
	// Get the center
	centerX := canvas.Width / 2
	centerY := canvas.Height / 2
	center := core.NewPoint(float64(centerX), float64(centerY), 0.0)
	fmt.Printf("\nCenter point %v+\n", center)
	canvas.WritePixel(int(center.X), int(center.Y), *color.Green)

	/*
		Learning: To draw a picture of your desire, you compute your drawing
		from your imagined origin taking unit points as reference and then
		compute the points to match your canvas dimensions
	*/
	refPointM := core.NewPoint(0, 1, 0).ToMatrix()
	hourRotateM := core.RotateZM(math.Pi / 6)
	for h := 0; h <= 11; h++ {
		// Rotate the previous hour point by pi/6
		nextHourPointM := core.ChainTransforms([]*core.Matrix{refPointM, hourRotateM})
		nextHourPoint := nextHourPointM.ToPoint()
		fmt.Printf("\n %d hour point %v", h+1, nextHourPoint)
		// canvas.WritePixel(int(nextHourPoint.X), int(nextHourPoint.Y), *Red)

		// Convert these points to our frame
		nextHourPointX := (nextHourPoint.X * radius) + float64(centerX)
		nextHourPointY := (nextHourPoint.Y * radius) + float64(centerY)
		fmt.Printf("\n next hour point %v %v\n", nextHourPointX, nextHourPointY)

		canvas.WritePixel(int(nextHourPointX), int(nextHourPointY), *color.Red)

		refPointM = nextHourPointM
	}

}

// ----- Draw Clock Main ---------
// func main() {
// 	canvas := world.NewCanvas(100, 100, *world.Black)
// 	drawClock(15, canvas)
// 	canvas.WriteToPPM("clock.ppm")
// }

// TODO: Update code to reflect the one in the textbook?
func drawSphereWithLight() {
	canvas := rendering.NewCanvas(100, 100, *color.Black)

	sphere := shape.UnitSphere()
	sphere.Transform = *core.ChainTransforms([]*core.Matrix{
		core.ScaleM(30, 30, 30),
		core.TranslationM(50, 50, 0),
	})
	sphere.Material = material.DefaultMaterial()
	sphere.Material.Color = *color.NewColor(1, 0.2, 1)

	// assumed ray origin
	rayOrigin := core.NewPoint(50, 50, 50)

	lightPosition := core.NewPoint(-5, 5, 55)
	lightColor := color.NewColor(1, 1, 1)
	light := lighting.NewLight(*lightColor, *lightPosition)

	for h := 0; h < canvas.Height; h++ {
		for w := 0; w < canvas.Width; w++ {
			pixel := core.NewPoint(float64(w), float64(h), 0.0)
			rayDirection := pixel.Subtract(*rayOrigin).Normalize()
			ray := rayt.Ray{Origin: *rayOrigin, Direction: *rayDirection}

			// TODO: A better interface possible ?
			// Intersect the ray with sphere
			intersections := ray.IntersectSphere(*sphere)
			hit := ray.Hit(intersections)

			if hit != nil {
				// for now, just color the pixel where the ray hit
				point := ray.Position(hit.T)
				normal := lighting.NormalAt(hit.Object, *point)
				eye := ray.Direction.Reverse()
				color := lighting.Lighting(hit.Object.Material, light, *point, *eye, normal)

				canvas.WritePixel(int(pixel.X), int(pixel.Y), color)
			}

		}
	}

	canvas.WriteToPPM("circle.ppm")

}

func main() {
	drawSphereWithLight()
}
