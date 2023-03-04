package vkm

import "github.com/chewxy/math32"

// Perspective generates a perspective projection matrix for Vulkan. Note that in Vulkan, the standard z clipping space is in the range
// [0..1], and x/y in the range [-1..1] (or rather -w to +w).
//
// aspect is the ratio of height divided by width
//
// fov is the vertical field of view measured in radians
//
// near and far are the near and far bounds of the frustum
func Perspective(fov, aspect, near, far float32) Mat {
	f := 1 / math32.Tan(fov/2)
	return Mat{
		{-f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, far / (near - far), -1},
		{0, 0, far * near / (near - far), 0},
	}
}

// GlTFPerspective is the perspective matrix mandated by the glTF specification.
func GlTFPerspective(fov, aspect, near, far float32) Mat {
	f := 1 / math32.Tan(fov/2)
	return Mat{
		{f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (near + far) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

// InvertedDepthPerspective generates a Vulkan perspective matrix with depth "reversed". This means that you must set
// DepthCompareOp to vk.COMPARE_OP_GREATER, and clear the depth buffer to zero, instead of one. The reason you might use
// this perspective over the "standard" perspecvtive matrix is for greater depth precision. Floating point numbers are
// more precise as you approach 1.0, and the persepctive matrix tends to compress verticies towards the near projection
// plane, hence closer to 1.0 in the depth buffer.
//
// Parameters for this function have the same meaning as [Perspective]
func InvertedDepthPerspective(fov, aspect, near, far float32) Mat {
	f := 1 / math32.Tan(fov/2)
	return Mat{ // Clipping Z reversed!
		{-f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, near / (far - near), -1},
		{0, 0, far * near / (far - near), 0},
	}
}

// PerspectiveDeg generates a perspective projection matrix. This version
// accepts the FOV angle in degrees. See Perspective.
func PerspectiveDeg(fovDeg, aspect, near, far float32) Mat {
	return Perspective(2.0*math32.Pi*fovDeg/360.0, aspect, near, far)
}

// Ortho generates an orthogonal projection matrix.
func OrthoProjection(width, height, near, far float32) Mat {
	return Mat{
		{2 / width, 0, 0, 0},
		{0, 2 / height, 0, 0},
		{0, 0, 1 / (far - near), (far + near) / (far - near)},
		{0, 0, 0, 1},
	}
}

// GlTFOrthoProjection is the orthographic projection matrix mandated by the glTF specification.
func GlTFOrthoProjection(xmag, ymag, znear, zfar float32) Mat {
	return Mat{
		{2 / xmag, 0, 0, 0},
		{0, 2 / ymag, 0, 0},
		{0, 0, 2 / (znear - zfar), (znear + zfar) / (znear - zfar)},
		{0, 0, 0, 1},
	}
}

// LookAt creates a view matrix from the provided eye and focus points, and an
// up vector. Note that this function does NOT generate a persepctive matrix. Note that LookAt (and all projection
// helpers in this library) are based on Vulkan's clip space, which differs from OpenGL.
//
// LookAt only orients the view from an arbitrary orientation down the positive Z axis. The up vector does not need to be orthagonal to
// the eye -> focus vector, but it cannot be parallel. Note that Vulkan's clip space is different than OpenGL:
// Z increases into the screen, Positive-Y is down the screen, and the clip boundaries are (-1 <= x <= 1, -1 <= y <= 1,
// 0 <= z <= 1).
func LookAt(eye, focus Pt, up Vec) Mat {
	return Camera(eye, focus.VecFrom(eye), up)
}

// Camera creates a view matrix from the provided eye location, pointed along the direction vector. Camera will continue to look "forward" as you move the eye point, while [LookAt] will change the look direction
// to continue pointing at the focus point as the eye moves around.
func Camera(eye Pt, look Vec, up Vec) Mat {
	// x := Mat{{1, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, 1}}.Inverse()

	w := look.Invert().Normalize()
	// u := w.Cross(up).Normalize()
	u := up.Cross(w).Normalize()
	// v := w.Cross(u)
	v := u.Cross(w)

	t := NewMatTranslate(eye.VecFrom(Origin()))

	m := Mat{u, v, w, {0, 0, 0, 1}} //.Transpose() //.Translate()
	r := t.MultM(m)
	// m = Identity()

	// m[1][1], m[2][2] = -1, -1
	// m[3][3] = -1
	// return x.MultM(m)
	return r.Inverse()
}
