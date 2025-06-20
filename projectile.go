package main

import (
	"fmt"
)

type Projectile struct {
	position Point
	velocity Vector
}

type Environment struct {
	gravity Vector
	wind    Vector
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.AddVector(proj.velocity)
	velocity := AddVectors([]Vector{proj.velocity, env.gravity, env.wind})
	return Projectile{*position, *velocity}
}

func projection(env Environment, proj Projectile, canvas *Canvas) Projectile {
	if proj.position.y <= 0 {
		return proj
	}

	newproj := tick(env, proj)
	// Plot each projection on the canvas
	// The y coordinate is subtracted with canvas height to match our canvas coordinates
	proj_canvas_y := canvas.height - int(newproj.position.y)
	proj_canvas_x := int(newproj.position.x)

	// exit early if the projectile is going out of bounds of canvas
	if proj_canvas_y < 0 || proj_canvas_y >= canvas.height || (proj_canvas_x >= canvas.width) {
		return proj
	}

	fmt.Printf("Projectile Position: %v+\n", newproj.position)
	fmt.Printf("Projectile canvas Position: %d, %d\n", proj_canvas_x, proj_canvas_y)

	canvas.WritePixel(proj_canvas_x, proj_canvas_y, *Red)

	return projection(env, newproj, canvas)
}

func ThrowProjectile(env Environment, proj Projectile, canvas *Canvas) Projectile {
	return projection(env, proj, canvas)
}

func main() {
	// fmt.Println(quote.Glass())

	// start := Point{0, 1, 0}
	// velocity := Vector{1, 1.8, 0}.Normalize().ScalarMultiply(11.25)
	// gravity := Vector{0, -0.1, 0}
	// wind := Vector{-0.01, 0, 0}

	// canvas := NewCanvas(900, 500, *Black)

	// // Projectile
	// p := Projectile{position: start, velocity: *velocity}
	// e := Environment{gravity: gravity, wind: wind}
	// finalProjection := ThrowProjectile(e, p, canvas)
	// canvas.WriteToPPM("projectile.ppm")
	// fmt.Printf("\nfinal projectile position: %v+\n", finalProjection.position)

	m1 := NewMatrix(2, 3, [][]float64{
		{1, 4}, {2, 5}, {3, 6},
	})

	m2 := NewMatrix(2, 2, [][]float64{
		{7, 9}, {8, 10},
	})

	result := m1.Multiply(*m2)
	fmt.Println(result)

}
