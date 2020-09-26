package vkm

import "testing"

func TestAsVec(t *testing.T) {
	sl := []float32{0, 1, 2, 3, 4, 5}
	v := AsVec(sl)

	for i, val := range v {
		if val != float32(i) {
			t.Errorf("AsVec did not copy correctly at index %d", i)
		}
	}

	v = AsVec(sl[:2])
	if v[0] != 0.0 || v[1] != 1.0 {
		t.Errorf("AsVec did not copy correctly.")
	}
	if v[2] != 0.0 || v[3] != 0 {
		t.Errorf("AsVec overwrote beyond the end of the slice. Result: %+v", v)
	}
}

func TestAsPt(t *testing.T) {
	sl := []float32{0, 1, 2, 3, 4, 5}
	v := AsPt(sl)

	for i, val := range v {
		if val != float32(i) {
			t.Errorf("AsPt did not copy correctly at index %d", i)
		}
	}
}
