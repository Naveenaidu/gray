package lighting

import (
	"github.com/Naveenaidu/gray/src/geom"
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
