package vkm

// Mat is a column-major 4x4 matrix of float32s. Because it is fundamentally an array of arrays,
// elements can be directly addressed via double brackets: m[col][row]
type Mat [4]Vec

// MultV computes a matrix multiplication on the provided vector
func (m Mat) MultV(v Vec) Vec {
	return Vec{
		m[0][0]*v[0] + m[1][0]*v[1] + m[2][0]*v[2] + m[3][0]*v[3],
		m[0][1]*v[0] + m[1][1]*v[1] + m[2][1]*v[2] + m[3][1]*v[3],
		m[0][2]*v[0] + m[1][2]*v[1] + m[2][2]*v[2] + m[3][2]*v[3],
		m[0][3]*v[0] + m[1][3]*v[1] + m[2][3]*v[2] + m[3][3]*v[3],
	}
}

// MultP computes a matrix multiplication on the provided point
func (m Mat) MultP(v Pt) Pt {
	return Pt(m.MultV(Vec(v)))
}

// MultM performs a matrix multiplication.
func (m Mat) MultM(n Mat) Mat {
	return Mat{
		{
			m[0][0]*n[0][0] + m[1][0]*n[0][1] + m[2][0]*n[0][2] + m[3][0]*n[0][3],
			m[0][1]*n[0][0] + m[1][1]*n[0][1] + m[2][1]*n[0][2] + m[3][1]*n[0][3],
			m[0][2]*n[0][0] + m[1][2]*n[0][1] + m[2][2]*n[0][2] + m[3][2]*n[0][3],
			m[0][3]*n[0][0] + m[1][3]*n[0][1] + m[2][3]*n[0][2] + m[3][3]*n[0][3],
		},
		{
			m[0][0]*n[1][0] + m[1][0]*n[1][1] + m[2][0]*n[1][2] + m[3][0]*n[1][3],
			m[0][1]*n[1][0] + m[1][1]*n[1][1] + m[2][1]*n[1][2] + m[3][1]*n[1][3],
			m[0][2]*n[1][0] + m[1][2]*n[1][1] + m[2][2]*n[1][2] + m[3][2]*n[1][3],
			m[0][3]*n[1][0] + m[1][3]*n[1][1] + m[2][3]*n[1][2] + m[3][3]*n[1][3],
		},
		{
			m[0][0]*n[2][0] + m[1][0]*n[2][1] + m[2][0]*n[2][2] + m[3][0]*n[2][3],
			m[0][1]*n[2][0] + m[1][1]*n[2][1] + m[2][1]*n[2][2] + m[3][1]*n[2][3],
			m[0][2]*n[2][0] + m[1][2]*n[2][1] + m[2][2]*n[2][2] + m[3][2]*n[2][3],
			m[0][3]*n[2][0] + m[1][3]*n[2][1] + m[2][3]*n[2][2] + m[3][3]*n[2][3],
		},
		{
			m[0][0]*n[3][0] + m[1][0]*n[3][1] + m[2][0]*n[3][2] + m[3][0]*n[3][3],
			m[0][1]*n[3][0] + m[1][1]*n[3][1] + m[2][1]*n[3][2] + m[3][1]*n[3][3],
			m[0][2]*n[3][0] + m[1][2]*n[3][1] + m[2][2]*n[3][2] + m[3][2]*n[3][3],
			m[0][3]*n[3][0] + m[1][3]*n[3][1] + m[2][3]*n[3][2] + m[3][3]*n[3][3],
		},
	}
}

// Identity returns a 4x4 identity matrix
func Identity() Mat {
	return Mat{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// Equals returns true if m and n are exactly equal. Note that this
// function does not allow for "approximately equal" conditions, like many
// floating point comparison functions.
func (m Mat) Equals(n Mat) bool {
	for i := range m {
		for j := range m[i] {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}

// Transpose returns the transpose matrix of m.
func (m Mat) Transpose() Mat {
	return Mat{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}
