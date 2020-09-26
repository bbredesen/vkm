package vkm

import (
	"testing"
)

func TestPerspective(t *testing.T) {
	pMat := PerspectiveDeg(90.0, 1.0, 2.0, 10.0)

	p0 := Origin()        // Should be clipped out, projected to negative Inf
	p1 := NewPt(0, 0, 1)  // Should be clipped on Z
	p2 := NewPt(0, 0, -2) // Should be projected to (0, 0, 0.0)
	p3 := NewPt(5, 5, -5) // Should be projected to (1, 1, <1.0)

	r0 := pMat.MultP(p0)
	r1 := pMat.MultP(p1)
	r2 := pMat.MultP(p2)
	r3 := pMat.MultP(p3)

	if !testClipped(r0) {
		t.Errorf("Origin was not clipped! Result: %+v", r0)
	}
	if !testClipped(r1) {
		t.Errorf("Point behind view (z = 1) was not clipped! Result: %+v", r1)
	}
	if testClipped(r2) {
		t.Errorf("Point on front plane was clipped! Result: %+v", r2)
	}
	if testClipped(r3) {
		t.Errorf("Point on back plane was clipped! Result: %+v", r3)
	}
}

func testClipped(pt Pt) bool {
	return !(pt[0] >= -pt[3] && pt[0] <= pt[3] &&
		pt[1] >= -pt[3] && pt[1] <= pt[3] &&
		pt[2] >= 0.0 && pt[2] <= pt[3])

}

func TestLookAt(t *testing.T) {
	mat0 := LookAt(Origin(), NewPt(0, 0, -1), NewVec(0, 1, 0))
	exp0 := Mat{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, -1, 1},
	}

	if !mat0.Equals(exp0) {
		t.Errorf("Look from origin matrix was not expected! Expected identity, returned: %+v", mat0)
	}
	mat1 := LookAt(NewPt(5, 0, 0), Origin(), NewVec(0, 1, 0))
	exp1 := Mat{
		{0, 0, -1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 5, 1},
	}

	if !mat1.Equals(exp1) {
		t.Errorf("Look from (1,1,1) matrix was not expected! Returned: %+v", mat1)
	}

}
