package vkm

import "github.com/chewxy/math32"

// NewMatTranslate generates a translation matrix using v.
// Note that the fourth component of v is ignored and that component of the matrix is set to 1.0
func NewMatTranslate(v Vec) Mat {
	rval := Identity()
	rval[3] = v
	rval[3][3] = 1.0
	return rval
}

// Translate applys the translation of v to m and returns the resulting matrix.
func (m Mat) Translate(v Vec) Mat {
	return NewMatTranslate(v).MultM(m)
}

// NewMatScale creates a scaling matrix using v. A negative value on any axis will create a
// reflection across that axis.
func NewMatScale(v Vec) Mat {
	rval := Identity()
	rval[0][0] = v[0]
	rval[1][1] = v[1]
	rval[2][2] = v[2]
	return rval
}

// Scale applies a scale transformation to m and returns the result
func (m Mat) Scale(v Vec) Mat {
	return m.MultM(NewMatScale(v))
}

/*** Rotation Transformations ***/

// NewMatRotateX generates a CCW rotation around the X axis by theta radians
func NewMatRotateX(theta float32) Mat {
	ct := math32.Cos(theta)
	st := math32.Sin(theta)

	return Mat{
		{1, 0, 0, 0},
		{0, ct, st, 0},
		{0, -st, ct, 0},
		{0, 0, 0, 1},
	}
}

// NewMatRotateXDeg generates a CCW rotation around the X axis by deg degrees
func NewMatRotateXDeg(deg float32) Mat {
	return NewMatRotateX(2 * math32.Pi * deg / 360.0)
}

// RotateX applies a rotation around the X axis to m and returns the resulting matrix.
func (m Mat) RotateX(theta float32) Mat {
	return NewMatRotateX(theta).MultM(m)
}

// RotateXDeg applies a rotation around the X axis to m and returns the resulting matrix.
func (m Mat) RotateXDeg(deg float32) Mat {
	return NewMatRotateXDeg(deg).MultM(m)
}

// NewMatRotateY generates a CCW rotation around the Y axis by theta radians
func NewMatRotateY(theta float32) Mat {
	ct := math32.Cos(theta)
	st := math32.Sin(theta)

	return Mat{
		{ct, 0, -st, 0},
		{0, 1, 0, 0},
		{st, 0, ct, 0},
		{0, 0, 0, 1},
	}
}

// NewMatRotateYDeg generates a CCW rotation around the Y axis by deg degrees
func NewMatRotateYDeg(deg float32) Mat {
	return NewMatRotateY(2 * math32.Pi * deg / 360.0)
}

// RotateY applies a rotation around the Y axis to m and returns the resulting matrix.
func (m Mat) RotateY(theta float32) Mat {
	return NewMatRotateY(theta).MultM(m)
}

// RotateYDeg applies a rotation around the Y axis to m and returns the resulting matrix.
func (m Mat) RotateYDeg(deg float32) Mat {
	return NewMatRotateYDeg(deg).MultM(m)
}

// NewMatRotateZ generates a CCW rotation around the Z axis by theta radians
func NewMatRotateZ(theta float32) Mat {
	ct := math32.Cos(theta)
	st := math32.Sin(theta)

	return Mat{
		{ct, st, 0, 0},
		{-st, ct, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// NewMatRotateZDeg generates a CCW rotation around the Z axis by deg degrees
func NewMatRotateZDeg(deg float32) Mat {
	return NewMatRotateZ(2 * math32.Pi * deg / 360.0)
}

// RotateZ applies a rotation around the Z axis to m and returns the resulting matrix.
func (m Mat) RotateZ(theta float32) Mat {
	return NewMatRotateZ(theta).MultM(m)
}

// RotateZDeg applies a rotation around the Z axis to m and returns the resulting matrix.
func (m Mat) RotateZDeg(deg float32) Mat {
	return NewMatRotateZDeg(deg).MultM(m)
}

// NewMatRotate generates a rotation by theta radians around an arbitrary axis.
// Note that axis is assumed to be a unit vector. The rotation angle will be
// scaled by the length of the axis vector.
func NewMatRotate(axis Vec, theta float32) Mat {
	ct := math32.Cos(theta)
	st := math32.Sin(theta)
	omct := 1 - ct

	xx := axis[0] * axis[0] * omct
	xy := axis[0] * axis[1] * omct
	xz := axis[0] * axis[2] * omct
	yy := axis[1] * axis[1] * omct
	yz := axis[1] * axis[2] * omct
	zz := axis[2] * axis[2] * omct

	xst := axis[0] * st
	yst := axis[1] * st
	zst := axis[2] * st

	return Mat{
		{ct + xx, xy + zst, xz - yst, 0},
		{xy - zst, ct + yy, yz + xst, 0},
		{xz + yst, yz - xst, ct + zz, 0},
		{0, 0, 0, 1},
	}
}

// NewMatRotateDeg generates a rotation by deg degrees around an arbitrary axis.
// Note that axis is assumed to be a unit vector. The rotation angle will be
// scaled by the length of the axis vector.
func NewMatRotateDeg(axis Vec, deg float32) Mat {
	return NewMatRotate(axis, 2*math32.Pi*deg/360.0)
}

// Rotate applies an arbitrary rotation measured in radian around the provided axis to m and returns the resulting matrix.
func (m Mat) Rotate(axis Vec, theta float32) Mat {
	return NewMatRotate(axis, theta).MultM(m)
}

// RotateDeg applies an arbitrary rotation measured in degrees around the provided axis to m and returns the resulting matrix.
func (m Mat) RotateDeg(axis Vec, deg float32) Mat {
	return NewMatRotateDeg(axis, deg).MultM(m)
}
