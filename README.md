# VKM

VKM is a simple vector and matrix math package written in Go, inspired by glm,
and targeting 3D graphics programing. 

VKM is specifically targeted at 3D graphics and is NOT inteded to be an
all-purpose vector math or linear algebra package. It includes some variant types like `Pt2` and
`Vec3` for convenience, but all matrix math functions rely on homogenous 3D
vectors and points.

While this module can be used in any Go program wanting to work with 3D transformations, it was built to use with
[go-vk](https://github.com/bbredesen/go-vk) (and [go-vk-samples](https://github.com/bbredesen/go-vk-samples)), a Go
language binding for the [Vulkan](https://khronos.org/vulkan/) graphics API.

## Usage

See the GoDoc for the full API.

VKM defines three primary types: `Pt`, `Vec`, and `Matrix`. All three are stored
in column-major order and are fundamentally arrays of `float32`. (Note: this explicitly means static arrays and not
slices.) 

We use `float32`, and not Go's default 65-bit floating point type, because Vulkan (generally) uses 32-bit
floats on the GPU. As a consequence, and rather than constantly forcing float32 type casts, VKM uses the [chewxy/math32 package](https://github.com/chewxy/math32) for float32 operations and constants. 

## Examples

### Create a basic transformation
You can generate a simpole tranformation with the `NewMat__` variants:

```go
import (
    "github.com/bbredesen/vkm"
    "github.com/chewxy/math32"
)
// ...
transVec := vkm.NewVec(-1, -2, -3)
transMat := vkm.NewMatTranslate(transVec)
```

Rotations are measured in radians by default, but each rotation function has a
"Deg" variant to use degrees. Rotation angles will scale with the length of the axis vector, so be sure to
normalize the vector first:

```go
axisVec := vkm.NewVec(1, 1, 1).Normalize()
aRotationMat := vkm.NewMatRotate(axisVec, math32.Pi)
```

### Multiply arbitrary matricies and points or vectors

Matrix multiplication happens in the order written when you directly call
`MultM`, `MultV`, or `MultP`. The following code has the apparent
effect of rotating the world before translating:
```go
transformMat := myTranslationMat.MultM(myRotationMat)
```

### Create a "readable" transformation
Transformations can be chained, as long as you have a matrix to start from. In
this example, we apply a translation of -5 in the X direction, rotate around the
Y axis by 45 degrees counterclockwise, and then translate by +5 in the X
direction:
```go
transVec := NewVec(-5, 0, 0)

m := Identity().
    Translate(transVec).
    RotateYDeg(-45).
    Translate(transVec.Invert())
```
When written in this form, we apply transformations in the order they are
written. In practice, the code is actually right-multiplying each subsequent
transformation, as in the previous example.

## Performance Optimization TODO

All math in this library is currently writing in pure Go. Performance could benefit from using SIMD extensions on
Intel/AMD CPUs, which will require writing those functions in assembly.