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
	if !result.Equals(id0) {
		t.Errorf("Expected identity matrix! Actual: %+v", result)
	}
}

func TestMultTranslate(t *testing.T) {
	id0 := Identity()
	t1 := NewMatTranslate(Vec{1, 2, 3})
	t2 := NewMatTranslate(Vec{-1, -2, -3})
	t3 := NewMatTranslate(Vec{4, 5, 6})

	r1 := id0.MultM(t1)
	if !r1.Equals(t1) {
		t.Errorf("Translate against identity failed! Expected: %+v Actual: %+v", t1, r1)
	}

	r2 := t1.MultM(t2)
	if !r2.Equals(id0) {
		t.Errorf("Translate with reversal failed! Expected identity, actual: %+v", r2)
	}

	r3 := t1.MultM(t3)
	ex := NewMatTranslate(Vec{5, 7, 9})
	if !r3.Equals(ex) {
		t.Errorf("General translate failed! Expected: %+v Actual: %+v", ex, r3)
	}
}
