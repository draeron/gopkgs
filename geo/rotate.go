package geo

import (
	"gonum.org/v1/gonum/num/quat"
	"math"
)

type Point3D struct {
	X, Y, Z float64
}

func (p Point3D) to2D() Point {
	return Point{p.X, p.Y}
}

// raise raises the dimensionality of a point to a quaternion.
func raise(p Point3D) quat.Number {
	return quat.Number{Imag: p.X, Jmag: p.Y, Kmag: p.Z}
}

// rotate performs the quaternion rotation of p by the given quaternion
// and scaling by the scale factor.

// taken from https://github.com/gonum/gonum/blob/ec146a97d7075b4d0e88377732605fa7670c83e2/num/quat/quat_example_test.go
func rotate(p Point3D, by quat.Number, scale float64) Point3D {

	// Ensure the modulus of by is correctly scaled.
	if len := quat.Abs(by); len != scale {
		by = quat.Scale(math.Sqrt(scale)/len, by)
	}

	// Perform the rotation/scaling.
	pp := quat.Mul(quat.Mul(by, raise(p)), quat.Conj(by))

	// Extract the point.
	return Point3D{pp.Imag, pp.Jmag, pp.Kmag}
}

func Rotate(p Point, radian float64) Point {
	return RotateAround(p, Zero, radian)
}

func RotateAround(p Point, origin Point, radian float64) Point {

	p.Translate(origin.Invert())

	q := raise(Point3D{0, 0, 1}) // rotate around Up vector
	scale := 1.0

	q = quat.Scale(math.Sin(radian/2)/quat.Abs(q), q)
	q.Real += math.Cos(radian / 2)

	p = rotate(p.to3D(), q, scale).to2D()

	p.Translate(origin)

	return p
}
