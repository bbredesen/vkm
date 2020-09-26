# VKM

VKM is a simple vector and matrix math package written in Go, inspired by glm, and targeting 3D graphics programing.

While this module can be used in any Go program, it was built to use with [go-vk](https://github.com/bbredesen/go-vk), a Go language binding for the [Vulkan](https://khronos.org/vulkan/) API.

## Usage

See the GoDoc for the full API.

VKM defines three primary types: `Pt`, `Vec`, and `Matrix`. All three are stored in column-major order and are fundamentally arrays of `float32`.

VKM is specifically targeted at 3D graphics and is not inteded to be an all-purpose vector math package. It includes some variant types like `Pt2` and `Vec3` for convenience, but all matrix math functions rely on homogenous 3D vectors and points with four elements.

## Examples:

### Create a basic transformation
You can generate a single tranformation with the `NewMat__` variants.
```
transVec := NewVec(-1, -2, -3)
aTranslationMat := NewMatTranslate(transVec)
```

Rotation angles will scale with the length of the axis vector, so be sure to normalize the length.

Rotations are measured in radians by default, but each rotation function has a "Deg" variant to use degrees.
```
axisVec := NewVec(1, 1, 1).Normalize()
aRotationMat := NewMatRotate(axisVec, 3.1415)
```

### Multiply matricies and points or vectors

Matrix multiplication happens in the order written, when you directly call `MultM`, `MultV`, or `MultP`. The following code rotates the world before translating:
```
transformMat := aTranslationMat.MultM(aRotationMat)
```

### Create a "readable" transformation matrix
Transformations can be chained, as long as you have a matrix to start from. In this example, we apply a translation of -5 in the X direction, rotate around the Y axis by 45 degrees counterclockwise, and then translate by +5 in the X direction.
```
transVec := NewVec(-5, 0, 0)

m := Identity().
    Translate(transVec).
    RotateYDeg(-45).
    Translate(transVec.Invert())
```
When written in this form, we apply transformations in the order they are written. In practice, the code is actually right-multiplying each subsequent transformation, as in the previous example.
