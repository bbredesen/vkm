package vkm

import "github.com/chewxy/math32"

// Perspective generates a perspective projection matrix. Note that the "camera" is looking down
// negative Z, by convention.
//
// aspect is the aspect ratio of height divided by width
//
// fov is the vertical field of view in radians
func Perspective(fov, aspect, near, far float32) Mat {
	f := 1 / math32.Tan(fov/2)

	return Mat{
		{f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, -far / (far - near), -1},
		{0, 0, -(far * near) / (far - near), 0},
	}
}

// PerspectiveDeg generates a perspective projection matrix. This version
// accepts the FOV angle in degrees. See Perspective.
func PerspectiveDeg(fovDeg, aspect, near, far float32) Mat {
	return Perspective(2.0*math32.Pi*fovDeg/360.0, aspect, near, far)
}

// LookAt creates a view matrix from the provided eye and focus points, and an up vector.
// Note that this function does NOT generate a persepctive matrix. LookAt only orients the
// view from an arbitrary orientation into "standardized" coordinates down the negative Z
// axis.
func LookAt(eye, focus Pt, up Vec) Mat {
	vBack := eye.VecFrom(focus).Normalize()
	vRight := up.Cross(vBack).Normalize()
	vUp := vBack.Cross(vRight)

	m := Mat{vRight, vUp, vBack, Vec(Origin())}
	t := NewMatTranslate(eye.VecTo(focus))
	return m.MultM(t)
}
