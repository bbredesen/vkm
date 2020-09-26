// Package vkm provides matrix math functions targeted at computer graphics.
// It was built for use with go-vk (https://github.com/bbredesen/go-vk), but should be
// appropriate for any graphics API expecting column-major matricies.
//
// All matricies are 4x4. The Vec and Pt types are 4-component, homogenous column
// vectors. You can directly instantiate them via:
//  v := Vec{i, j, k, 0}
//  p := Pt{x, y, z, 1}
// but it is recommended instead to call the New... functions to ensure that the fourth
// component is correctly set.
//  v := NewVec(i, j, k)
//  p := NewPt(x, y, z)
//
// Convenience "zero value" functions for the Origin, ZeroVec, and Identity matrix are included.
//
// Additionally, standard transformation matricies are provided. You can either generate
// the matrix directly:
//  NewMatTranslate(transVector)
// or you can apply a transformation to an existing matrix:
//  myMat.Translate(transVector)
//
// Note that applied transformations are left-multiplied, so m.Translate(v) = T(v) x m, whereas
// calling MultM will right-multiply: m.MultM(t) = m x t.
package vkm
