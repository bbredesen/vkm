package vkm

import (
	"testing"

	"github.com/chewxy/math32"
)

func TestRotateX(t *testing.T) {
	m0 := NewMatRotateX(math32.Pi)
	pt := NewPt(0, 1, 1)

	exp0 := NewPt(0, -1, -1)
	res0 := m0.MultP(pt)

	if !res0.EqualTo(exp0) {
		t.Errorf("NewMatRotateX by pi radians failed! Expected: %+v Actual: %+v", exp0, res0)
	}

	m1 := NewMatRotateXDeg(90)
	exp1 := NewPt(0, -1, 1)
	res1 := m1.MultP(pt)

	if !res1.EqualTo(exp1) {
		t.Errorf("NewMatRotateXDeg by 90 degrees failed! Expected: %+v Actual: %+v", exp1, res1)
	}

	m2 := Identity()
	exp2 := NewPt(0, 1, -1)
	res2 := m2.RotateXDeg(-90).MultP(pt)
	if !res2.EqualTo(exp2) {
		t.Errorf("RotateXDeg by -90 degrees failed! Expected: %+v Actual: %+v", exp2, res2)
	}
}

func TestRotateY(t *testing.T) {
	m0 := NewMatRotateY(math32.Pi)
	pt := NewPt(1, 1, 1)

	exp0 := NewPt(-1, 1, -1)
	res0 := m0.MultP(pt)

	if !res0.EqualTo(exp0) {
		t.Errorf("NewMatRotateY by pi radians failed! Expected: %+v Actual: %+v", exp0, res0)
	}

	m1 := NewMatRotateYDeg(90)
	exp1 := NewPt(1, 1, -1)
	res1 := m1.MultP(pt)

	if !res1.EqualTo(exp1) {
		t.Errorf("NewMatRotateYDeg by 90 degrees failed! Expected: %+v Actual: %+v", exp1, res1)
	}

	m2 := Identity()
	exp2 := NewPt(-1, 1, 1)
	res2 := m2.RotateYDeg(-90).MultP(pt)
	if !res2.EqualTo(exp2) {
		t.Errorf("RotateYDeg by -90 degrees failed! Expected: %+v Actual: %+v", exp2, res2)
	}
}

func TestRotateZ(t *testing.T) {
	m0 := NewMatRotateZ(math32.Pi)
	pt := NewPt(1, 1, 1)

	exp0 := NewPt(-1, -1, 1)
	res0 := m0.MultP(pt)

	if !res0.EqualTo(exp0) {
		t.Errorf("NewMatRotateZ by pi radians failed! Expected: %+v Actual: %+v", exp0, res0)
	}

	m1 := NewMatRotateZDeg(90)
	exp1 := NewPt(-1, 1, 1)
	res1 := m1.MultP(pt)

	if !res1.EqualTo(exp1) {
		t.Errorf("NewMatRotateZDeg by 90 degrees failed! Expected: %+v Actual: %+v", exp1, res1)
	}

	m2 := Identity()
	exp2 := NewPt(1, -1, 1)
	res2 := m2.RotateZDeg(-90).MultP(pt)
	if !res2.EqualTo(exp2) {
		t.Errorf("RotateZDeg by -90 degrees failed! Expected: %+v Actual: %+v", exp2, res2)
	}
}

func TestRotate(t *testing.T) {
	// Rotation around Z
	m0 := NewMatRotate(NewVec(0, 0, 1), math32.Pi)
	pt := NewPt(1, 1, 1)

	exp0 := NewPt(-1, -1, 1)
	res0 := m0.MultP(pt)

	if !res0.EqualTo(exp0) {
		t.Errorf("NewMatRotate by pi radians around Z axis failed! Expected: %+v Actual: %+v", exp0, res0)
	}

	// Rotation around axis that the test point lives on, by an arbitrary angle of 5 radians.
	// Should not change the point at all.
	m1 := NewMatRotate(NewVec(1, 1, 1).Normalize(), 5)
	exp1 := pt.Clone()
	res1 := m1.MultP(pt)
	if !res1.EqualTo(exp1) {
		t.Errorf("NewMatRotate applied to point on rotation axis failed! Expected: %+v Actual: %+v", exp1, res1)
	}

	m2 := Identity()
	exp2 := NewPt(1, 1, -1)
	axis2 := NewVec(1, 0, 0)
	res2 := m2.RotateDeg(axis2, -90).MultP(pt)
	if !res2.EqualTo(exp2) {
		t.Errorf("RotateDeg by -90 degrees around X axis failed! Expected: %+v Actual: %+v", exp2, res2)
	}

}
