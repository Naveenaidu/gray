package geom

type Sphere struct {
	Center    Point
	Radius    float64
	Transform Matrix
}

// Sphere with radius 1 and centered at origin (0,0,0)
func UnitSphere() *Sphere {
	return &Sphere{Center: *NewPoint(0, 0, 0), Radius: 1.0, Transform: *IdentityMatrix()}
}
