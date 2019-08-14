package mathf

import (
	"fmt"
	"math"
)

// A Quaternion describes a rotation in 3D space.
// The Quaternion is mathematically defined as Q = x*i + y*j + z*k + w, where (i,j,k) are imaginary basis vectors.
// (x,y,z) can be seen as a vector related to the axis of rotation, while the real multiplier, w, is related to the amount of rotation.
type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewQuaternion creates a new quaternion with given values
func NewQuaternion(x float64, y float64, z float64, w float64) *Quaternion {
	return &Quaternion{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// NewZeroQuaternion creates a new zero quaternion [0,0,0,1]
func NewZeroQuaternion() *Quaternion {
	return &Quaternion{
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
		W: 1.0,
	}
}

// Set the value of the quaternion
func (quaternion *Quaternion) Set(x float64, y float64, z float64, w float64) *Quaternion {
	quaternion.X = x
	quaternion.Y = y
	quaternion.Z = z
	quaternion.W = w
	return quaternion
}

func (quaternion *Quaternion) Add(other *Quaternion) *Quaternion {
	return &Quaternion{
		X: quaternion.X + other.X,
		Y: quaternion.Y + other.Y,
		Z: quaternion.Z + other.Z,
		W: quaternion.W + other.W,
	}
}

// QuaternionFromAxisAngle creates the quaternion components from an axis and an angle
func QuaternionFromAxisAngle(axis *Vec3, angle float64) *Quaternion {
	s := math.Sin(angle * 0.5)

	return &Quaternion{
		X: axis.X * s,
		Y: axis.Y * s,
		Z: axis.Z * s,
		W: math.Cos(angle * 0.5),
	}
}

// ToAxisAngle converts the quaternion to axis/angle representation.
func (quaternion *Quaternion) ToAxisAngle() (*Vec3, float64) {
	quaternion.Normalize()

	angle := 2.0 * math.Acos(quaternion.W)
	s := math.Sqrt(1.0 - quaternion.W*quaternion.W)

	if s < 0.0001 {
		return &Vec3{
			X: quaternion.X,
			Y: quaternion.Y,
			Z: quaternion.Z,
		}, angle
	}

	return &Vec3{
		X: quaternion.X / s,
		Y: quaternion.Y / s,
		Z: quaternion.Z / s,
	}, angle
}

// QuaternionFromVectors creates a quaternion from the given two vectors.The resulting rotation will be the needed rotation to rotate u to v.
func QuaternionFromVectors(u *Vec3, v *Vec3) *Quaternion {
	if u.IsAntiparallelTo(v, Epsilon) {
		t1 := Vec3{}
		t2 := Vec3{}

		u.Tangents(&t1, &t2)
		return QuaternionFromAxisAngle(&t1, math.Pi)
	}

	a := u.Cross(v)

	quaternion := &Quaternion{
		X: a.X,
		Y: a.Y,
		Z: a.Z,
		W: math.Sqrt(u.SqrtLength()*v.SqrtLength()) + u.Dot(v),
	}

	quaternion.Normalize()

	return quaternion
}

func (quaternion *Quaternion) Length() float64 {
	return math.Sqrt(
		quaternion.X*quaternion.X +
			quaternion.Y*quaternion.Y +
			quaternion.Z*quaternion.Z +
			quaternion.W*quaternion.W)
}

// Normalize the quaternion
func (quaternion *Quaternion) Normalize() *Quaternion {
	l := quaternion.Length()

	if l <= 0.0 {
		return NewZeroQuaternion()
	}

	invLength := 1.0 / l

	norm := &Quaternion{
		X: quaternion.X * invLength,
		Y: quaternion.Y * invLength,
		Z: quaternion.Z * invLength,
		W: quaternion.W * invLength,
	}

	return norm
}

// Multiply this quaternion by the given
func (quaternion *Quaternion) Multiply(other *Quaternion) *Quaternion {
	return &Quaternion{
		X: quaternion.X*other.W + quaternion.W*other.X + quaternion.Y*other.Z - quaternion.Z*other.Y,
		Y: quaternion.Y*other.W + quaternion.W*other.Y + quaternion.Z*other.X - quaternion.X*other.Z,
		Z: quaternion.Z*other.W + quaternion.W*other.Z + quaternion.X*other.Y - quaternion.Y*other.X,
		W: quaternion.W*other.W - quaternion.X*other.X - quaternion.Y*other.Y - quaternion.Z*other.Z,
	}
}

// Inverse calculates the inverse quaternion rotation.
func (quaternion *Quaternion) Inverse() *Quaternion {
	c := quaternion.Conjugate()

	inorm2 := 1.0 / (quaternion.X*quaternion.X + quaternion.Y*quaternion.Y + quaternion.Z*quaternion.Z + quaternion.W*quaternion.W)

	c.X = c.X * inorm2
	c.Y = c.Y * inorm2
	c.Z = c.Z * inorm2
	c.W = c.W * inorm2

	return c
}

// Conjugate calculates the quaternion conjugate
func (quaternion *Quaternion) Conjugate() *Quaternion {
	return &Quaternion{
		X: -quaternion.X,
		Y: -quaternion.Y,
		Z: -quaternion.Z,
		W: quaternion.W,
	}
}

// MultiplyVec multiply the quaternion by a vector
func (quaternion *Quaternion) MultiplyVec(vec *Vec3) *Vec3 {
	ix := quaternion.W*vec.X + quaternion.Y*vec.Z - quaternion.Z*vec.Y
	iy := quaternion.W*vec.Y + quaternion.Z*vec.X - quaternion.X*vec.Z
	iz := quaternion.W*vec.Z + quaternion.X*vec.Y - quaternion.Y*vec.X
	iw := -quaternion.X*vec.X - quaternion.Y*vec.Y - quaternion.Z*vec.Z

	return NewVec3(
		(ix*quaternion.W)+(iw*-quaternion.X)+(iy*-quaternion.Z)-(iz*-quaternion.Y),
		(iy*quaternion.W)+(iw*-quaternion.Y)+(iz*-quaternion.X)-(ix*-quaternion.Z),
		(iz*quaternion.W)+(iw*-quaternion.Z)+(ix*-quaternion.Y)-(iy*-quaternion.X))
}

// Clone this quaternion to new instance
func (quaternion *Quaternion) Clone() *Quaternion {
	return &Quaternion{
		X: quaternion.X,
		Y: quaternion.Y,
		Z: quaternion.Z,
		W: quaternion.W,
	}
}

// ToEuler convert the quaternion to euler angle representation, Order: YZX
func (quaternion *Quaternion) ToEuler() *Vec3 {
	sqx := quaternion.X * quaternion.X
	sqy := quaternion.Y * quaternion.Y
	sqz := quaternion.Z * quaternion.Z

	euler := NewZeroVec3()

	// roll (x-axis rotation)
	sinrCosp := -2.0 * (quaternion.Y*quaternion.Z - quaternion.X*quaternion.W)
	cosrCosp := 1.0 - 2.0*(sqx+sqy)
	euler.X = math.Atan2(sinrCosp, cosrCosp)

	// pitch (y-axis rotation)
	sinp := 2.0 * (quaternion.W*quaternion.Y + quaternion.Z*quaternion.X)
	if math.Abs(sinp) >= 1.0 {
		euler.Y = math.Copysign(math.Pi/2.0, sinp)
	} else {
		euler.Y = math.Asin(sinp)
	}

	// yaw (z-axis rotation)
	sinyCosp := -2.0 * (quaternion.X*quaternion.Y - quaternion.W*quaternion.Z)
	cosyCosp := 1.0 - 2.0*(sqy+sqz)
	euler.Z = math.Atan2(sinyCosp, cosyCosp)

	return euler
}

// QuaternionFromEuler creates the quaternion from the given euler angels
func QuaternionFromEuler(vec *Vec3) *Quaternion {
	c1 := math.Cos(vec.X / 2.0)
	c2 := math.Cos(vec.Y / 2.0)
	c3 := math.Cos(vec.Z / 2.0)
	s1 := math.Sin(vec.X / 2.0)
	s2 := math.Sin(vec.Y / 2.0)
	s3 := math.Sin(vec.Z / 2.0)

	return &Quaternion{
		X: s1*c2*c3 + c1*s2*s3,
		Y: c1*s2*c3 - s1*c2*s3,
		Z: c1*c2*s3 + s1*s2*c3,
		W: c1*c2*c3 - s1*s2*s3,
	}
}

// Integrate rotate an absolute orientation quaternion given an angular velocity and a time step.
func (quaternion *Quaternion) Integrate(angularVelocity *Vec3, dt float64, angularFactor *Vec3) *Quaternion {
	ax := angularVelocity.X * angularFactor.X
	ay := angularVelocity.Y * angularFactor.Y
	az := angularVelocity.Z * angularFactor.Z

	halfDT := dt * 0.5

	return &Quaternion{
		X: halfDT * (ax*quaternion.W + ay*quaternion.Z - az*quaternion.Y),
		Y: halfDT * (ay*quaternion.W + az*quaternion.X - ax*quaternion.Z),
		Z: halfDT * (az*quaternion.W + ax*quaternion.Y - ay*quaternion.X),
		W: halfDT * (-ax*quaternion.X - ay*quaternion.Y - az*quaternion.Z),
	}
}

// Slerp performs a spherical linear interpolation between two quat
func (quaternion *Quaternion) Slerp(toQuaternion *Quaternion, t float64) *Quaternion {
	cosom := quaternion.X*toQuaternion.X + quaternion.Y*toQuaternion.Y + quaternion.Z*toQuaternion.Z + quaternion.W*toQuaternion.W

	if cosom < 0.0 {
		cosom = -cosom
		toQuaternion = &Quaternion{
			X: toQuaternion.X,
			Y: toQuaternion.Y,
			Z: toQuaternion.Z,
			W: toQuaternion.W,
		}
	}

	scale0 := 1.0 - t
	scale1 := t

	if (1.0 - cosom) > 0.000001 {
		omega := math.Acos(cosom)
		sinom := math.Sin(omega)

		scale0 = math.Sin((1.0-t)*omega) / sinom
		scale1 = math.Sin(t*omega) / sinom
	}

	return &Quaternion{
		X: scale0*quaternion.X + scale1*toQuaternion.X,
		Y: scale0*quaternion.X + scale1*toQuaternion.Y,
		Z: scale0*quaternion.Z + scale1*toQuaternion.Z,
		W: scale0*quaternion.W + scale1*toQuaternion.W,
	}
}

func (quaternion *Quaternion) String() string {
	return fmt.Sprintf("Quaternion [ x: %f, y: %f, z: %f, w: %f ]", quaternion.X, quaternion.Y, quaternion.Z, quaternion.W)
}
