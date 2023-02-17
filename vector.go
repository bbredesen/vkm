package vkm

import (
	"github.com/chewxy/math32"
)

// Vec represents a homogenous 3d vector (i, j, k, and l components)
type Vec [4]float32

// Vec3 is a 3d vector (i, j, and k components)
type Vec3 [3]float32

// Vec2 is a 2d vector (i and j components)
type Vec2 [2]float32

// NewVec initializes a homogenous vector using the provided i, j, and k values
func NewVec(i, j, k float32) Vec {
	return [4]float32{i, j, k, 0}
}

// ZeroVec returns a homogenous zero vector
func ZeroVec() Vec {
	return [4]float32{0, 0, 0, 0}
}

// UnitVecX returns a unit vector in the positive X direction
func UnitVecX() Vec {
	return [4]float32{1, 0, 0, 0}
}

// UnitVecY returns a unit vector in the positive Y direction
func UnitVecY() Vec {
	return [4]float32{0, 1, 0, 0}
}

// UnitVecZ returns a unit vector in the positive Z direction
func UnitVecZ() Vec {
	return [4]float32{0, 0, 1, 0}
}

// AsVec converts a slice of float32 to a Vec. This function does not validate the input in any way, so if the slice has
// more than four elements, only the first four will be copied. If the slice has less than four elements, the remaining
// elements in the return value will be zeroes.
func AsVec(s []float32) Vec {
	rval := ZeroVec()
	sliceCopy(rval[:], s)
	return rval
}

func sliceCopy(target, source []float32) {
	len := len(source)
	if len > 4 {
		len = 4
	}
	copy(target, source[:len])
}

// Homogenize converts a Vec2 in to a Vec, with the k and l components fixed to 0.0
func (v Vec2) Homogenize() Vec {
	return Vec{v[0], v[1], 0.0, 0.0}
}

// Add returns the sum of two vectors
func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{v[0] + u[0], v[1] + u[1]}
}

// Sub returns the difference between two vectors
func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{v[0] - u[0], v[1] - u[1]}
}

// Invert returns a Vec2 of the same magnitude, pointed in the opposite direction
func (v Vec2) Invert() Vec2 {
	return Vec2{-v[0], -v[1]}
}

// Homogenize converts a Vec3 into a Vec, with the l component fixed to 0.0
func (v Vec3) Homogenize() Vec {
	return Vec{v[0], v[1], v[2], 0.0}
}

// Add returns the sum of two vectors
func (v Vec3) Add(u Vec3) Vec3 {
	return [3]float32{v[0] + u[0], v[1] + u[1], v[2] + u[2]}
}

// Add returns the sum of two vectors. Note that the w component from the input is ignored, and is clamped to zero on the output.
func (v Vec) Add(u Vec) Vec {
	return [4]float32{v[0] + u[0], v[1] + u[1], v[2] + u[2], 0.0}
}

// Sub returns the difference between two vectors
func (v Vec3) Sub(u Vec3) Vec3 {
	return [3]float32{v[0] - u[0], v[1] - u[1], v[2] - u[2]}
}

// Sub returns the difference between two vectors. Note that the w component from the input is ignored, and is clamped to zero on the output.
func (v Vec) Sub(u Vec) Vec {
	return [4]float32{v[0] - u[0], v[1] - u[1], v[2] - u[2], 0.0}
}

// Invert returns a Vec3 of the same magnitude, pointed in the opposite direction
func (v Vec3) Invert() Vec3 {
	return [3]float32{-v[0], -v[1], -v[2]}
}

// Invert returns a Vec of the same magnitude, pointed in the opposite direction
func (v Vec) Invert() Vec {
	return [4]float32{-v[0], -v[1], -v[2], 0.0}
}

// Clone creates an duplicate copy of a Vec
func (v Vec) Clone() Vec {
	var rval Vec
	copy(rval[:], v[:])
	return rval
}

// Cross returns the cross product of v and u Note that the w component is ignored, though it
// should be zero for any homogenous vector
func (v Vec) Cross(u Vec) Vec {
	return NewVec(
		v[1]*u[2]-v[2]*u[1],
		v[2]*u[0]-v[0]*u[2],
		v[0]*u[1]-v[1]*u[0],
	)
}

// Dot returns the dot product of v and u. Note that the w component is ignored, though it
// should be zero for any homogenous vector
func (v Vec) Dot(u Vec) float32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}
func (v Vec3) Dot(u Vec3) float32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}
func (v Vec2) Dot(u Vec2) float32 {
	return v[0]*u[0] + v[1]*u[1]
}

// Length is the length of v. This function calculates a square root, making it relatively expensive. Use
// [Vec2.SquareLength] instead if you are comparing two vectors, or comparing against a constant length.
func (v Vec2) Length() float32 {
	return math32.Sqrt(v.SquareLength())
}
func (v Vec2) SquareLength() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

// Length is the length in 3D space of v. This function calculates a square root, making it relatively expensive. Use
// [Vec3.SquareLength] instead if you are comparing two vectors, or comparing against a constant length.
func (v Vec3) Length() float32 {
	return math32.Sqrt(v.SquareLength())
}
func (v Vec3) SquareLength() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Length is the length in 3D space of v. This function calculates a square root, making it relatively expensive. Use
// [Vec.SquareLength] instead if you are comparing two vectors, or comparing against a constant length.
func (v Vec) Length() float32 {
	return math32.Sqrt(v.Dot(v))
}

// Calculates the dot product of v with itself, equal to the length of the vector squared. This is much more efficient
// than calling Length() if you are comparing the length of two vectors.
func (v Vec) SquareLength() float32 {
	return v.Dot(v)
}

// Normalize returns a unit vector oriented in the same direction as v.
func (v Vec) Normalize() Vec {
	l := v.Length()
	return Vec{v[0] / l, v[1] / l, v[2] / l, 0}
}

func (v Vec3) Normalize() Vec3 {
	l := v.Length()
	return Vec3{v[0] / l, v[1] / l, v[2] / l}
}
func (v Vec2) Normalize() Vec2 {
	l := v.Length()
	return Vec2{v[0] / l, v[1] / l}
}

func (v Vec) RotateZ(theta float32) Vec {
	return NewMatRotate(UnitVecZ(), theta).MultV(v)
}

func (v Vec) Scale(factor float32) Vec {
	return Vec{v[0] * factor, v[1] * factor, v[2] * factor, v[3] * factor}
}
func (v Vec3) Scale(factor float32) Vec3 {
	return Vec3{v[0] * factor, v[1] * factor, v[2] * factor}
}
func (v Vec2) Scale(factor float32) Vec2 {
	return Vec2{v[0] * factor, v[1] * factor}
}
