package geom

import (
	"math"
	"math/rand"
)

// Vec2 is a vector in 2D space with cartesian coordinates. Holds 2 components:
// x and y in this order.
type Vec2 [2]int

// Vec3 is a vector in 3D space with cartesian coordinates. Holds 3 components:
// x, y and z in this order.
type Vec3 [3]float64

// Norm normalizes a vector to length 1 keeping its direction.
func (v *Vec3) Norm() {
	abs := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	if abs != 0 {
		v[0] /= abs
		v[1] /= abs
		v[2] /= abs
	}
}

// Neg negates the vector's components.
func (v *Vec3) Neg() {
	v[0] = -v[0]
	v[1] = -v[1]
	v[2] = -v[2]
}

// Sub subtracts another vector.
func (v *Vec3) Sub(w *Vec3) {
	v[0] -= w[0]
	v[1] -= w[1]
	v[2] -= w[2]
}

// Add adds another vector.
func (v *Vec3) Add(w *Vec3) {
	v[0] += w[0]
	v[1] += w[1]
	v[2] += w[2]
}

// Scale scales the vector.
func (v *Vec3) Scale(s float64) {
	v[0] *= s
	v[1] *= s
	v[2] *= s
}

// Cross returns a new vector that is the cross product of the two vectors.
func Cross(v, w *Vec3) *Vec3 {
	return &Vec3{
		v[1]*w[2] - v[2]*w[1],
		v[2]*w[0] - v[0]*w[2],
		v[0]*w[1] - v[1]*w[0],
	}
}

// Vec4 is a vector in 3D space with homogeneous coordinates. Holds 4
// components: x, y, z and w in this order.
type Vec4 [4]float64

// Norm normalizes a homogeneous vector by dividing x, y and z by w so that w
// will be 1.
func (v *Vec4) Norm() {
	v[0] /= v[3]
	v[1] /= v[3]
	v[2] /= v[3]
	v[3] = 1.0
}

// NewVec4 returns a new vector with homogeneous coordinates corresponding to
// the given cartesian coordinates (w will be 1).
func NewVec4(x, y, z float64) *Vec4 {
	return &Vec4{x, y, z, 1}
}

// Mat4 is a matrix with homogeneous coordinates used to transform homogeneous
// vectors.  Holds 16 components, the 4 first elements make up the first row
// from left to right, and so on.
type Mat4 [16]float64

// ZeroMat returns a new matrix with all values set to zero.
func ZeroMat() *Mat4 {
	return &Mat4{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}

// RandMat returns a new matrix random values.
func RandMat(r *rand.Rand) *Mat4 {
	m := Mat4{}
	for i := 0; i < len(m); i++ {
		m[i] = r.Float64()
	}
	return &m
}

// Mul multiplies the matrix with another one, modifying the former one.
func (m *Mat4) Mul(n *Mat4) {
	t := ZeroMat()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				t[i*4+j] += m[i*4+k] * n[j+k*4]
			}
		}
	}
	*m = *t
}

// Transf returns a new transformed vector by multiplying the matrix with the
// given vector.
func (m *Mat4) Transf(v *Vec4) *Vec4 {
	p := Vec4{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			p[i] += m[i*4+j] * v[j]
		}
	}
	return &p
}
