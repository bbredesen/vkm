package vkm

// Pt represents a homogenous 3d point (x, y, z, and w components)
type Pt Vec

// Pt3 is a 3d point (x, y, and z components)
type Pt3 Vec3

// Pt2 is a 2d point (x and y components only), intended for use as texture coordinates.
type Pt2 [2]float32

// NewPt initilizes a homogenous point using the provided x, y, and z values
func NewPt(x, y, z float32) Pt {
	return Pt{x, y, z, 1}
}

// NewPt3 initializes a 3D point using the provided values.
func NewPt3(x, y, z float32) Pt3 {
	return Pt3{x, y, z}
}

// NewPt2 initializes a 2D point using the provided values
func NewPt2(x, y float32) Pt2 {
	return Pt2{x, y}
}

// Origin is the X/Y/Z origin (0,0,0) as a homogenous point
func Origin() Pt {
	return Pt{0, 0, 0, 1}
}

// AsPt onverts a slice of float32 to a Pt. If the slice has more than four elements, only the first four will be copied.
func AsPt(s []float32) Pt {
	rval := Origin()
	sliceCopy(rval[:], s)
	return rval
}

// Clone creates an duplicate copy of a Pt
func (p Pt) Clone() Pt {
	var rval Pt
	copy(rval[:], p[:])
	return rval
}

// Clone creates an duplicate copy of a Pt3
func (p Pt3) Clone() Pt3 {
	var rval Pt3
	copy(rval[:], p[:])
	return rval
}

// Clone creates an duplicate copy of a Pt2
func (p Pt2) Clone() Pt2 {
	var rval Pt2
	copy(rval[:], p[:])
	return rval
}

// VecTo returns a vector directed from this point to the specified point q
func (p Pt) VecTo(q Pt) Vec {
	return [4]float32{q[0] - p[0], q[1] - p[1], q[2] - p[2], q[3] - p[3]}
}

// VecFrom returns a vector directed to this point, from the specified point q
func (p Pt) VecFrom(q Pt) Vec {
	return [4]float32{p[0] - q[0], p[1] - q[1], p[2] - q[2], p[3] - q[3]}
}

// Add returns the point where p is translated by v
func (p Pt) Add(v Vec) Pt {
	return Pt{p[0] + v[0], p[1] + v[1], p[2] + v[2], p[3] + v[3]}
}

// Homogenize converts a Pt3 into a Pt, with the w component fixed to 1.0
func (p Pt3) Homogenize() Pt {
	return Pt{p[0], p[1], p[2], 1.0}
}

// Homogenize on a Pt divides all components by the w coordinate
func (p Pt) Homogenize() Pt {
	return Pt{p[0] / p[3], p[1] / p[3], p[2] / p[3], 1.0}
}

// EqualTo determines if two points are (approximately) equal. If any component
// differs by more than 0.00001, then the two points are not "equal"
func (p Pt) EqualTo(q Pt) bool {
	lamMin := float32(-0.00001)
	lamMax := -lamMin
	dx := p[0] - q[0]
	dy := p[1] - q[1]
	dz := p[2] - q[2]
	dw := p[3] - q[3]

	return !(dx < lamMin || dx > lamMax ||
		dy < lamMin || dy > lamMax ||
		dz < lamMin || dz > lamMax ||
		dw < lamMin || dw > lamMax)

}
