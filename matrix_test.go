package vkm

import "testing"

func TestMultV(t *testing.T) {
	identity := Identity()
	v := Vec{1, 2, 3, 1}

	r := identity.MultV(v)
	for i, val := range r {
		if val != v[i] {
			t.Errorf("Identity multiply returned a non-equal vector! Expected %+v, found %+v", v, r)
		}
	}

	translate := Mat{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{1, 1, 1, 1},
	}

	r = translate.MultV(v)
	expected := Vec{2, 3, 4, 1}
	for i := range r {
		if r[i] != expected[i] {
			t.Errorf("Translate multiply did not return the expected result! Expected %+v, found %+v", expected, r)
		}
	}
}

func TestMultIdentity(t *testing.T) {
	id0 := Identity()

	result := id0.MultM(id0)
	if !result.ApproximatelyEquals(id0, 0.00001) {
		t.Errorf("Expected identity matrix! Actual: %+v", result)
	}
}

func TestMultTranslate(t *testing.T) {
	id0 := Identity()
	t1 := NewMatTranslate(Vec{1, 2, 3})
	t2 := NewMatTranslate(Vec{-1, -2, -3})
	t3 := NewMatTranslate(Vec{4, 5, 6})

	r1 := id0.MultM(t1)
	if !r1.ApproximatelyEquals(t1, 0.00001) {
		t.Errorf("Translate against identity failed! Expected: %+v Actual: %+v", t1, r1)
	}

	r2 := t1.MultM(t2)
	if !r2.ApproximatelyEquals(id0, 0.00001) {
		t.Errorf("Translate with reversal failed! Expected identity, actual: %+v", r2)
	}

	r3 := t1.MultM(t3)
	ex := NewMatTranslate(Vec{5, 7, 9})
	if !r3.ApproximatelyEquals(ex, 0.00001) {
		t.Errorf("General translate failed! Expected: %+v Actual: %+v", ex, r3)
	}
}

func TestDeterminant(t *testing.T) {
	m0 := Mat{
		{1, 1, 1, -1},
		{1, 1, -1, 1},
		{1, -1, 1, 1},
		{-1, 1, 1, 1},
	}
	if d := m0.Determinant(); d != -16 {
		t.Errorf("Determinant failed! Expected: %v Actual: %v", -16, d)
	}
}

func TestInverse(t *testing.T) {
	mId := Identity()
	mIdInv := mId.Inverse()
	exId := Identity()

	if !mIdInv.ApproximatelyEquals(exId, 0.00001) {
		t.Errorf("Inverse on identity failed! Expected: %+v Actual: %+v", exId, mIdInv)
	}

	m0 := Mat{
		{1, 1, 1, -1},
		{1, 1, -1, 1},
		{1, -1, 1, 1},
		{-1, 1, 1, 1},
	}
	m0Inv := m0.Inverse()
	ex0 := Mat{
		{0.25, 0.25, 0.25, -0.25},
		{0.25, 0.25, -0.25, 0.25},
		{0.25, -0.25, 0.25, 0.25},
		{-0.25, 0.25, 0.25, 0.25},
	}

	if !m0Inv.ApproximatelyEquals(ex0, 0.00001) {
		t.Errorf("Inverse on test 1 failed! Expected: %+v Actual: %+v", ex0, m0Inv)
	}

	m1 := Mat{
		{1, 2, 0, 0},
		{-2, 4, 0, 0},
		{0, 0, 1, 0},
		{3, 6, 9, 1},
	}
	m1Inv := m1.Inverse()
	ex1 := Mat{
		{0.5, -0.25, 0, 0},
		{0.25, 0.125, 0, 0},
		{0, 0, 1, 0},
		{-3, 0, -9, 1},
	}
	if !m1Inv.ApproximatelyEquals(ex1, 0.00001) {
		t.Errorf("Inverse on test 2 failed! Expected: %+v Actual: %+v", ex1, m1Inv)
	}

}
