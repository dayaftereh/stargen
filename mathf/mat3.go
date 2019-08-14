package mathf

import "fmt"

// Mat3 is a 3x3 matrix.
type Mat3 struct {
	elements []float64
}

// NewMat3 creates a matrix with the given components
func NewMat3(e1 float64, e2 float64, e3 float64,
	e4 float64, e5 float64, e6 float64,
	e7 float64, e8 float64, e9 float64) *Mat3 {
	return &Mat3{
		elements: []float64{
			e1, e2, e3,
			e4, e5, e6,
			e7, e8, e9,
		},
	}
}

// NewZeroMat3 creats a Matrix with zero components
func NewZeroMat3() *Mat3 {
	return &Mat3{
		elements: []float64{
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
		},
	}
}

// NewIdentityMat3 creates a new Identity Matrix
func NewIdentityMat3() *Mat3 {
	return &Mat3{
		elements: []float64{
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
		},
	}
}

// Mat3FromQuaternion creates a matrix from the given Quaternion
func Mat3FromQuaternion(q *Quaternion) *Mat3 {
	x2, y2, z2 := q.X+q.X, q.Y+q.Y, q.Z+q.Z

	xx, xy, xz := q.X*x2, q.X*y2, q.X*z2
	yy, yz, zz := q.Y*y2, q.Y*z2, q.Z*z2
	wx, wy, wz := q.W*x2, q.W*y2, q.W*z2

	target := NewZeroMat3()

	target.elements[3*0+0] = 1 - (yy + zz)
	target.elements[3*0+1] = xy - wz
	target.elements[3*0+2] = xz + wy

	target.elements[3*1+0] = xy + wz
	target.elements[3*1+1] = 1 - (xx + zz)
	target.elements[3*1+2] = yz - wx

	target.elements[3*2+0] = xz + wy
	target.elements[3*2+1] = yz + wx
	target.elements[3*2+2] = 1 - (xx + yy)

	return target
}

// SetTrace sets the matrix diagonal elements from a Vec3
func (mat *Mat3) SetTrace(v *Vec3) {
	mat.elements[0] = v.X
	mat.elements[4] = v.Y
	mat.elements[8] = v.Z
}

// GetTrace returns the matrix diagonal elements
func (mat *Mat3) GetTrace(v *Vec3) *Vec3 {
	return NewVec3(
		mat.elements[0],
		mat.elements[4],
		mat.elements[8],
	)
}

// MultiplyVec multiplicat given vector with Matrix
func (mat *Mat3) MultiplyVec(v *Vec3) *Vec3 {
	return NewVec3(
		mat.elements[0]*v.X+mat.elements[1]*v.Y+mat.elements[2]*v.Z,
		mat.elements[3]*v.X+mat.elements[4]*v.Y+mat.elements[5]*v.Z,
		mat.elements[6]*v.X+mat.elements[7]*v.Y+mat.elements[8]*v.Z,
	)
}

// ScaleMultiply Matrix-scalar multiplication
func (mat *Mat3) ScaleMultiply(scale float64) *Mat3 {
	other := NewZeroMat3()
	for i, value := range mat.elements {
		other.elements[i] = value * scale
	}
	return other
}

// Multiply Matrix multiplication
func (mat *Mat3) Multiply(other *Mat3) *Mat3 {
	product := NewZeroMat3()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum := 0.0
			for k := 0; k < 3; k++ {
				sum += other.elements[i+k*3] * mat.elements[k+j*3]
			}
			product.elements[i+j*3] = sum
		}
	}
	return product
}

// Scale each column of the matrix
func (mat *Mat3) Scale(v *Vec3) *Mat3 {
	product := NewZeroMat3()
	for i := 0; i != 3; i++ {
		product.elements[3*i+0] = v.X * mat.elements[3*i+0]
		product.elements[3*i+1] = v.Y * mat.elements[3*i+1]
		product.elements[3*i+2] = v.Z * mat.elements[3*i+2]
	}
	return product
}

// Element returns element at i
func (mat *Mat3) Element(i int) float64 {
	return mat.elements[i]
}

// Elements creates the components for the matrix
func (mat *Mat3) Elements() []float64 {
	return mat.elements
}

// Clone this matrix a new object
func (mat *Mat3) Clone() *Mat3 {
	product := NewZeroMat3()
	for i, value := range mat.elements {
		product.elements[i] = value
	}
	return product
}

// Transpose the matrix
func (mat *Mat3) Transpose() *Mat3 {
	product := NewZeroMat3()

	for i := 0; i != 3; i++ {
		for j := 0; j != 3; j++ {
			product.elements[3*i+j] = mat.elements[3*j+i]
		}
	}

	return product
}

func (mat *Mat3) String() string {
	s := fmt.Sprintf("[ %f, %f, %f ]\n", mat.elements[0], mat.elements[1], mat.elements[2])
	s = fmt.Sprintf("%s[ %f, %f, %f ]\n", s, mat.elements[3], mat.elements[4], mat.elements[5])
	s = fmt.Sprintf("%s[ %f, %f, %f ]\n", s, mat.elements[6], mat.elements[7], mat.elements[8])
	return s
}
