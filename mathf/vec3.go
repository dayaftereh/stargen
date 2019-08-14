package mathf

import (
	"fmt"
	"math"
)

// Vec3 is a 3-dimensional vector
type Vec3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// NewVec3 create a vector with theiven values
func NewVec3(x float64, y float64, z float64) *Vec3 {
	return &Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewZeroVec3 creates a zero vector
func NewZeroVec3() *Vec3 {
	return &Vec3{
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
	}
}

// NewUnitX creates the vector (1,0,0)
func NewUnitX() *Vec3 {
	return &Vec3{
		X: 1.0,
		Y: 0.0,
		Z: 0.0,
	}
}

// NewUnitY creates the vector (0,1,0)
func NewUnitY() *Vec3 {
	return &Vec3{
		X: 0.0,
		Y: 1.0,
		Z: 0.0,
	}
}

// NewUnitZ creates the vector (0,0,1)
func NewUnitZ() *Vec3 {
	return &Vec3{
		X: 0.0,
		Y: 0.0,
		Z: 1.0,
	}
}

// Set the given values to this vector
func (vec *Vec3) Set(x float64, y float64, z float64) *Vec3 {
	vec.X = x
	vec.Y = y
	vec.Z = z
	return vec
}

// SetVec set the given values from the given vector
func (vec *Vec3) SetVec(other *Vec3) *Vec3 {
	vec.X = other.X
	vec.Y = other.Y
	vec.Z = other.Z
	return vec
}

func (vec *Vec3) Add(x float64) *Vec3 {
	return &Vec3{
		X: vec.X + x,
		Y: vec.Y + x,
		Z: vec.Z + x,
	}
}

// AddVec the given vector to this vector and retuns a new vector.
func (vec *Vec3) AddVec(other *Vec3) *Vec3 {
	return &Vec3{
		X: vec.X + other.X,
		Y: vec.Y + other.Y,
		Z: vec.Z + other.Z,
	}
}

func (vec *Vec3) Subtract(x float64) *Vec3 {
	return &Vec3{
		X: vec.X - x,
		Y: vec.Y - x,
		Z: vec.Z - x,
	}
}

// SubtractVec the given vector from this vector and retuns a new vector.
func (vec *Vec3) SubtractVec(other *Vec3) *Vec3 {
	return &Vec3{
		X: vec.X - other.X,
		Y: vec.Y - other.Y,
		Z: vec.Z - other.Z,
	}
}

// Normalize the vector. Note that this changes the values in the vector.
func (vec *Vec3) Normalize() *Vec3 {
	length := vec.Length()
	if length > 0.0 {
		invLength := 1.0 / length
		vec.X *= invLength
		vec.Y *= invLength
		vec.Z *= invLength
	} else {
		vec.X = 0
		vec.Y = 0
		vec.Z = 0
	}
	return vec
}

// Unit returns the version of this vector that is of length 1.
func (vec *Vec3) Unit() *Vec3 {
	length := vec.Length()
	if length > 0.0 {
		invLength := 1.0 / length
		return &Vec3{
			X: vec.X * invLength,
			Y: vec.Y * invLength,
			Z: vec.Z * invLength,
		}
	}
	return NewZeroVec3()
}

// Length calculate the length of the vector
func (vec *Vec3) Length() float64 {
	sqrtLength := vec.SqrtLength()
	return math.Sqrt(sqrtLength)
}

// SqrtLength calculate the squared length of the vector.
func (vec *Vec3) SqrtLength() float64 {
	return vec.Dot(vec)
}

// DistanceTo calculate distance from this point to another point
func (vec *Vec3) DistanceTo(other *Vec3) float64 {
	sqrtDistanceTo := vec.SqrtDistanceTo(other)
	return math.Sqrt(sqrtDistanceTo)
}

// SqrtDistanceTo calculate squared distance from this point to another point
func (vec *Vec3) SqrtDistanceTo(other *Vec3) float64 {
	return ((vec.X - other.X) * (vec.X - other.X)) +
		((vec.Y - other.Y) * (vec.Y - other.Y)) +
		((vec.Z - other.Z) * (vec.Z - other.Z))
}

// Multiply all the components of the vector with a scalar.
func (vec *Vec3) Multiply(scale float64) *Vec3 {
	return &Vec3{
		X: vec.X * scale,
		Y: vec.Y * scale,
		Z: vec.Z * scale,
	}
}

// MultiplyVec multiply the vector with an other vector, component-wise
func (vec *Vec3) MultiplyVec(other *Vec3) *Vec3 {
	return &Vec3{
		X: vec.X * other.X,
		Y: vec.Y * other.Y,
		Z: vec.Z * other.Z,
	}
}

// Dot calculate dot product
func (vec *Vec3) Dot(other *Vec3) float64 {
	return vec.X*other.X + vec.Y*other.Y + vec.Z*other.Z
}

// Cross calculate the cross product
func (vec *Vec3) Cross(other *Vec3) *Vec3 {
	return &Vec3{
		X: (vec.Y * other.Z) - (vec.Z * other.Y),
		Y: (vec.Z * other.X) - (vec.X * other.Z),
		Z: (vec.X * other.Y) - (vec.Y * other.X),
	}
}

// Negate make the vector point in the opposite direction.
func (vec *Vec3) Negate() *Vec3 {
	return &Vec3{
		X: -vec.X,
		Y: -vec.Y,
		Z: -vec.Z,
	}
}

// Tangents compute two artificial tangents to the vector
func (vec *Vec3) Tangents(t1 *Vec3, t2 *Vec3) {
	length := vec.Length()

	if length > 0.0 {
		invLength := 1.0 / length
		n := NewVec3(vec.X*invLength, vec.Y*invLength, vec.Z*invLength)

		tmpVec := NewZeroVec3()
		if math.Abs(n.X) < 0.9 {
			tmpVec.Set(1.0, 0.0, 0.0)
		} else {
			tmpVec.Set(0.0, 1.0, 0.0)
		}

		t1.SetVec(n.Cross(tmpVec))
		t2.SetVec(n.Cross(t1))

		return
	}

	t1.Set(1.0, 0.0, 0.0)
	t2.Set(0.0, 1.0, 0.0)
}

// Clone create a copy of this vector
func (vec *Vec3) Clone() *Vec3 {
	return &Vec3{
		X: vec.X,
		Y: vec.Y,
		Z: vec.Z,
	}
}

func (vec *Vec3) Divide(scale float64) *Vec3 {
	return &Vec3{
		X: vec.X / scale,
		Y: vec.Y / scale,
		Z: vec.Z / scale,
	}
}

func (vec *Vec3) DivideVec(other *Vec3) *Vec3 {
	return &Vec3{
		X: vec.X / other.X,
		Y: vec.Y / other.Y,
		Z: vec.Z / other.Z,
	}
}

// Lerp do a linear interpolation between two vectors
func (vec *Vec3) Lerp(other *Vec3, t float64) *Vec3 {
	return &Vec3{
		X: vec.X + (other.X-vec.X)*t,
		Y: vec.Y + (other.Y-vec.Y)*t,
		Z: vec.Z + (other.Z-vec.Z)*t,
	}
}

// AlmostEquals check if a vector equals is almost equal to another one.
func (vec *Vec3) AlmostEquals(other *Vec3, precision float64) bool {
	if math.Abs(vec.X-other.X) > precision || math.Abs(vec.Y-other.Y) > precision || math.Abs(vec.Z-other.Z) > precision {
		return false
	}
	return true
}

func (vec *Vec3) AngleTo(other *Vec3) float64 {
	theta := vec.Dot(other) / (math.Sqrt(vec.SqrtLength() + other.SqrtLength()))
	return math.Acos(Clamp(theta, -1.0, 1.0))
}

// AlmostZero check if a vector is almost zero
func (vec *Vec3) AlmostZero(precision float64) bool {
	if math.Abs(vec.X) > precision || math.Abs(vec.Y) > precision || math.Abs(vec.Z) > precision {
		return false
	}
	return true
}

// IsAntiparallelTo check if the vector is anti-parallel to another vector.
func (vec *Vec3) IsAntiparallelTo(other *Vec3, precision float64) bool {
	n := vec.Negate()
	return n.AlmostEquals(other, precision)
}

func (vec *Vec3) String() string {
	return fmt.Sprintf("vec3 [ x: %f, y: %f, z: %f ]", vec.X, vec.Y, vec.Z)
}
