package material

import "github.com/Naveenaidu/gray/src/geom"

type Sphere struct {
	Center    geom.Point
	Radius    float64
	Transform geom.Matrix
	Material  Material
}

// Sphere with radius 1 and centered at origin (0,0,0)
func UnitSphere() *Sphere {
	return &Sphere{
		Center:    *geom.NewPoint(0, 0, 0),
		Radius:    1.0,
		Transform: *geom.IdentityMatrix(),
		Material:  DefaultMaterial(),
	}
}
