package lighting

import (
	"math"

	"github.com/Naveenaidu/gray/src/geom"
	"github.com/Naveenaidu/gray/src/material"
	"github.com/Naveenaidu/gray/src/world"
)

/*
To calculate the normal for a sphere at a point, we do the following:
1. Convert the sphere to object world
2. Calculate the normal in object world
3. Convert the object world normal into world coordinate system
*/
func NormalAt(s material.Sphere, p geom.Point) geom.Vector {
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

func Lighting(material material.Material, light Light, point geom.Point, eyev geom.Vector, normalv geom.Vector) world.Color {
	// combine the surface color with the light's color/intensity
	effectiveColor := world.MultiplyColors([]world.Color{material.Color, light.Intensity})

	// find the direction of light source
	lightV := light.Position.Subtract(point).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.ScalarMultiply(material.Ambient)
	diffuse := world.Black
	specular := world.Black

	// light_dot_normal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.
	lightDotNormal := lightV.DotProduct(normalv)
	if lightDotNormal > 0 {
		diffuse = effectiveColor.ScalarMultiply(material.Diffuse).ScalarMultiply(lightDotNormal)
	}

	// reflect_dot_eye represents the cosine of the angle between the
	// reflection vector and the eye vector. A negative number means the
	// light reflects away from the eye.
	reflectV := Reflect(*lightV.ScalarMultiply(-1), normalv)
	reflectDotEye := reflectV.DotProduct(eyev)
	if reflectDotEye > 0 {
		factor := math.Pow(reflectDotEye, float64(material.Shininess))
		specular = light.Intensity.ScalarMultiply(material.Specular).ScalarMultiply(factor)
	}

	return *world.AddColors([]world.Color{*ambient, *diffuse, *specular})
}
