package vkm

import (
	"github.com/chewxy/math32"
)

// Vec represents a homogenous 3d vector (i, j, k, and l components)
type Vec [4]float32

// Vec3 is a 3d vector (i, j, and k components)
type Vec3 [3]float32

// NewVec initializes a homogenous vector using the provided i, j, and k values
func NewVec(i, j, k float32) Vec {
	return [4]float32{i, j, k, 0}
}

// ZeroVec returns a homogenous zero vector
func ZeroVec() Vec {
	return [4]float32{0, 0, 0, 0}
}

// AsVec converts a slice of float32 to a Vec. If the slice has more than four elements, only the first four will be copied.
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

// Homogenize converts a Vec3 into a Vec, with the l component fixed to 0.0
func (v Vec3) Homogenize() Vec {
	return Vec{v[0], v[1], v[2], 0.0}
}

// Add returns the sum of two vectors
func (v Vec3) Add(w Vec3) Vec3 {
	return [3]float32{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

// Add returns the sum of two vectors. Note that this includes
func (v Vec) Add(w Vec) Vec {
	return [4]float32{v[0] + w[0], v[1] + w[1], v[2] + w[2], 0.0}
}

// Sub returns the difference between two vectors
func (v Vec3) Sub(w Vec3) Vec3 {
	return [3]float32{v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}

// Sub returns the difference between two vectors
func (v Vec) Sub(w Vec) Vec {
	return [4]float32{v[0] - w[0], v[1] - w[1], v[2] - w[2], 0.0}
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

// Cross returns the cross product of v and w
func (v Vec) Cross(w Vec) Vec {
	return NewVec(
		v[1]*w[2]-v[2]*w[1],
		v[2]*w[0]-v[0]*w[2],
		v[0]*w[1]-v[1]*w[0],
	)
}

// Dot returns the dot product of v and w. Note that the w component is ignored, though it
// should be zero for any homogenous vector
func (v Vec) Dot(w Vec) float32 {
	return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
}

// Length is the length in 3D space of v.
func (v Vec) Length() float32 {
	return math32.Sqrt(v.Dot(v))
}

// Normalize returns a unit vector oriented in the same direction as v.
func (v Vec) Normalize() Vec {
	l := v.Length()
	return Vec{v[0] / l, v[1] / l, v[2] / l, 0}
}
