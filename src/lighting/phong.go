package lighting

import (
	"github.com/Naveenaidu/gray/src/geom"
	"github.com/Naveenaidu/gray/src/world"
)

/*
To calculate the normal for a sphere at a point, we do the following:
1. Convert the sphere to object world
2. Calculate the normal in object world
3. Convert the object world normal into world coordinate system
*/
func NormalAt(s geom.Sphere, p geom.Point) geom.Vector {
	// FIXME: Check if the transform can be invertible
	// get the sphere to be at the origin in object world
	invertTransformM := s.Transform.Inverse()
	objectPoint := invertTransformM.Multiply(*p.ToMatrix()).ToPoint()
	objectNormal := objectPoint.Subtract(*geom.ObjectOrigin())
	// use transpose of inverse matrix to convert vector in object space to
	// world space
	// world_normal ← transpose(inverse(sphere.transform)) * object_normal
	worldNormal := invertTransformM.Transpose().Multiply(*objectNormal.ToMatrix()).ToVector()

	return *worldNormal.Normalize()
}

func Reflect(in geom.Vector, normal geom.Vector) geom.Vector {
	// reflect(in, normal) = in - normal * 2 * dot(in, normal)
	dotProduct := in.DotProduct(normal)
	reflection := geom.SubtractVectors([]geom.Vector{
		in,
		*normal.ScalarMultiply(2 * dotProduct),
	})
	return *reflection
}

type Light struct {
	Intensity world.Color
	Position  geom.Point
}

func NewLight(intensity world.Color, pos geom.Point) Light {
	return Light{intensity, pos}
}

type Material struct {
	Color     world.Color
	Ambient   float64 // ranges between 0 and 1
	Diffuse   float64 // ranges between 0 and 1
	Specular  float64 // ranges between 0 and 1
	Shininess int     // ranges between 10 and 200
}

func DefaultMaterial() Material {
	return Material{
		Color:     *world.NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
