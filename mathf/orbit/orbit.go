package orbit

import (
	"math"

	"github.com/dayaftereh/stargen/mathf"
)

// https://github.com/jordanstephens/kepler.js

type Orbit struct {
	r                 *mathf.Vec3 // position
	v                 *mathf.Vec3 // velocity
	mu                float64     // mu = G*M
	CentralBodyRadius float64
}

func NewOrbit(r *mathf.Vec3, v *mathf.Vec3, mu float64, centralBodyRadius float64) *Orbit {
	return &Orbit{
		r:                 r,
		v:                 v,
		mu:                mu,
		CentralBodyRadius: centralBodyRadius,
	}
}

func (orbit *Orbit) AngularMomentum() *mathf.Vec3 {
	return orbit.r.Cross(orbit.v)
}

func (orbit *Orbit) RadialVelocity() float64 {
	return orbit.r.Dot(orbit.v) / orbit.r.Length()
}

func (orbit *Orbit) Eccentricity() *mathf.Vec3 {
	return orbit.r.Multiply(
		math.Pow(orbit.v.Length(), 2) - (orbit.mu / orbit.r.Length()),
	).SubtractVec(
		orbit.v.Multiply(orbit.r.Length() * orbit.RadialVelocity()),
	).Multiply(1.0 / orbit.mu)
}

func (orbit *Orbit) SemimajorAxis() float64 {
	h := orbit.AngularMomentum().Length()
	e := orbit.Eccentricity().Length()
	return (math.Pow(h, 2) / orbit.mu) * (1.0 / (1.0 - math.Pow(e, 2)))
}

func (orbit *Orbit) SemiminorAxis() float64 {
	a := orbit.SemimajorAxis()
	e := orbit.Eccentricity().Length()
	return a * math.Sqrt(1.0-math.Pow(e, 2))
}

func (orbit *Orbit) SemilatusRectum() float64 {
	return math.Pow(orbit.AngularMomentum().Length(), 2) / orbit.mu
}

func (orbit *Orbit) Inclination() float64 {
	h := orbit.AngularMomentum()
	K := mathf.NewVec3(0, 0, 1)

	return mathf.ToDegress(math.Acos(K.Dot(h) / h.Length()))
}

func (orbit *Orbit) NodeLine() *mathf.Vec3 {
	K := mathf.NewVec3(0, 0, 1)
	return K.Cross(orbit.AngularMomentum())
}

func (orbit *Orbit) RightAscension() float64 {
	n := orbit.NodeLine()
	if mathf.CloseZero(n.Length()) {
		return 0.0
	}
	omega := mathf.ToDegress(math.Acos(n.X / n.Length()))
	if n.Y < 0 {
		return 360.0 - omega
	}
	return omega
}

func (orbit *Orbit) ArgumentOfPeriapsis() float64 {
	n := orbit.NodeLine()
	if mathf.CloseZero(n.Length()) {
		return 0.0
	}
	e := orbit.Eccentricity()
	w := mathf.ToDegress(math.Acos(n.Dot(e) / (n.Length() * e.Length())))
	if n.Z < 0 {
		return 360.0 - w
	}
	return w
}

func (orbit *Orbit) TrueAnomaly() float64 {
	e := orbit.Eccentricity()
	eNorm := e.Length()
	n := orbit.NodeLine()
	nNorm := n.Length()

	var u float64
	if mathf.CloseZero(eNorm) && mathf.CloseZero(nNorm) {
		u = mathf.ToDegress(math.Acos(math.Min(1.0, orbit.r.X/orbit.r.Length())))
	} else {
		l := e
		if mathf.CloseZero(eNorm) {
			l = n
		}
		u = mathf.ToDegress(math.Acos(math.Min(1.0, l.Dot(orbit.r)/(l.Length()*orbit.r.Length()))))
	}
	if orbit.r.Dot(orbit.v) < 0.0 {
		return 360.0 - u
	}
	return u
}

func (orbit *Orbit) Apoapsis() float64 {
	h := orbit.AngularMomentum()
	e := orbit.Eccentricity()
	return (math.Pow(h.Length(), 2.0) / orbit.mu) * (1.0 / (1.0 + e.Length()*math.Cos(math.Pi)))
}

func (orbit *Orbit) Periapsis() float64 {
	h := orbit.AngularMomentum()
	e := orbit.Eccentricity()
	return (math.Pow(h.Length(), 2.0) / orbit.mu) * (1.0 / (1.0 + e.Length()*math.Cos(0)))
}

func (orbit *Orbit) Period() float64 {
	a := orbit.SemimajorAxis()
	return (2.0 * math.Pi / math.Sqrt(orbit.mu)) * math.Sqrt(math.Pow(a, 3))
}

func (orbit *Orbit) UniversalAnomaly(dt float64) float64 {
	a := orbit.SemimajorAxis()
	// initial guess of x
	x := math.Sqrt(orbit.mu) / (dt / a)

	r := orbit.r
	v := orbit.v
	mu := orbit.mu

	f := func(x2 float64) float64 {
		return UniversalFormulationF(x2, a, r, v, mu, dt)
	}

	df := func(x2 float64) float64 {
		return UniversalFormulationDFDT(x2, a, r, v, mu)
	}

	d2f := func(x2 float64) float64 {
		return UniversalFormulationD2FDT(x2, a, r, v, mu)
	}

	return LaguerreSolve(x, f, df, d2f)
}

func (orbit *Orbit) Update(dt float64) *Orbit {
	x := orbit.UniversalAnomaly(dt)
	a := orbit.SemimajorAxis()
	z := UniversalFormulationZ(x, a)

	r0 := orbit.r
	v0 := orbit.v
	mu := orbit.mu

	r := r0.Multiply(LagrangeF(x, z, orbit.r)).AddVec(v0.Multiply(LagrangeG(x, z, mu, dt)))
	v := r0.Multiply(LagrangeDf(x, z, r, r0, mu)).AddVec(v0.Multiply(LagrangeDg(x, z, r)))

	return &Orbit{
		r:                 r,
		v:                 v,
		mu:                mu,
		CentralBodyRadius: orbit.CentralBodyRadius,
	}
}

func (orbit *Orbit) Position() *mathf.Vec3 {
	return orbit.r.Clone()
}

func (orbit *Orbit) Velocity() *mathf.Vec3 {
	return orbit.v.Clone()
}
