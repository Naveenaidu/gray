package main

import "fmt"

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

func projection(env Environment, proj Projectile) Projectile {
	if proj.position.y <= 0 {
		return proj
	}
	newproj := tick(env, proj)
	fmt.Printf("Projectile Position: %v+\n", newproj.position)
	return projection(env, newproj)
}

func ThrowProjectile(env Environment, proj Projectile) Projectile {
	return projection(env, proj)
}
