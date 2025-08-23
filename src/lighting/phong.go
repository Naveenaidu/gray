package lighting

import (
	"math"

	color "github.com/Naveenaidu/gray/src/core/color"
	core "github.com/Naveenaidu/gray/src/core/math"
	"github.com/Naveenaidu/gray/src/material"
)

/*
To calculate the normal for a sphere at a point, we do the following:
1. Convert the sphere to object world
2. Calculate the normal in object world
3. Convert the object world normal into world coordinate system
*/
func NormalAt(s material.Sphere, p core.Point) core.Vector {
	// FIXME: Check if the transform can be invertible
	// get the sphere to be at the origin in object world
	invertTransformM := s.Transform.Inverse()
	objectPoint := invertTransformM.Multiply(*p.ToMatrix()).ToPoint()
	objectNormal := objectPoint.Subtract(*core.ObjectOrigin())
	// use transpose of inverse matrix to convert vector in object space to
	// world space
	// world_normal ← transpose(inverse(sphere.transform)) * object_normal
	worldNormal := invertTransformM.Transpose().Multiply(*objectNormal.ToMatrix()).ToVector()

	return *worldNormal.Normalize()
}

func Reflect(in core.Vector, normal core.Vector) core.Vector {
	// reflect(in, normal) = in - normal * 2 * dot(in, normal)
	dotProduct := in.DotProduct(normal)
	reflection := core.SubtractVectors([]core.Vector{
		in,
		*normal.ScalarMultiply(2 * dotProduct),
	})
	return *reflection
}

type Light struct {
	Intensity color.Color
	Position  core.Point
}

func NewLight(intensity color.Color, pos core.Point) Light {
	return Light{intensity, pos}
}

func Lighting(material material.Material, light Light, point core.Point, eyev core.Vector, normalv core.Vector) color.Color {
	// combine the surface color with the light's color/intensity
	effectiveColor := color.MultiplyColors([]color.Color{material.Color, light.Intensity})

	// find the direction of light source
	lightV := light.Position.Subtract(point).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.ScalarMultiply(material.Ambient)
	diffuse := color.Black
	specular := color.Black

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

	return *color.AddColors([]color.Color{*ambient, *diffuse, *specular})
}
