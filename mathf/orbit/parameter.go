package orbit

import (
	"math"

	"github.com/dayaftereh/stargen/mathf"
)

func OrbitFromParams(params *OrbitParameter) *Orbit {
	params = extendParams(params)

	rP := perifocalPosition(
		*params.AngularMomentum,
		*params.Eccentricity,
		*params.TrueAnomaly,
		*params.MU,
	)

	vP := perifocalVelocity(
		*params.AngularMomentum,
		*params.Eccentricity,
		*params.TrueAnomaly,
		*params.MU,
	)

	q := transformMatrix(
		*params.ArgumentOfPeriapsis,
		*params.Inclination,
		*params.RightAscension,
	)

	r := q.MultiplyVec(rP)
	v := q.MultiplyVec(vP)

	return NewOrbit(r, v, *params.MU, *params.CentralBodyRadius)
}

func perifocalPosition(angularMomentum float64, eccentricity float64, trueAnomaly float64, mu float64) *mathf.Vec3 {
	h := angularMomentum
	e := eccentricity
	theta := mathf.ToRadians(trueAnomaly)

	return mathf.NewVec3(
		math.Cos(theta), math.Sin(theta), 0.0,
	).Multiply(
		(math.Pow(h, 2.0) / mu) *
			(1.0 / (1.0 + (e * math.Cos(theta)))))
}

func perifocalVelocity(angularMomentum float64, eccentricity float64, trueAnomaly float64, mu float64) *mathf.Vec3 {
	h := angularMomentum
	e := eccentricity
	theta := mathf.ToRadians(trueAnomaly)

	return mathf.NewVec3(
		-math.Sin(theta), e+math.Cos(theta), 0.0,
	).Multiply(mu / h)
}

func transformMatrix(argumentOfPeriapsis float64, inclination float64, rightAscension float64) *mathf.Mat3 {
	w := mathf.ToRadians(argumentOfPeriapsis)
	i := mathf.ToRadians(inclination)
	omega := mathf.ToRadians(rightAscension)

	sinOmega := math.Sin(omega)
	cosOmega := math.Cos(omega)
	sinI := math.Sin(i)
	cosI := math.Cos(i)
	sinW := math.Sin(w)
	cosW := math.Cos(w)

	return mathf.NewMat3(
		-sinOmega*cosI*sinW+(cosOmega*cosW),
		-sinOmega*cosI*cosW-(cosOmega*sinW),
		sinOmega*sinI,

		cosOmega*cosI*sinW+(sinOmega*cosW),
		cosOmega*cosI*cosW-(sinOmega*sinW),
		-cosOmega*sinI,

		sinI*sinW,
		sinI*cosW,
		cosI,
	)
}

func angularMomentumFromSemilatusRectum(semilatusRectum float64, mu float64) float64 {
	return math.Sqrt(semilatusRectum * mu)
}

func semilatusRectumFromSemimajorAxisAndEccentricity(semimajorAxis float64, eccentricity float64) float64 {
	return semimajorAxis * (1.0 - (math.Pow(eccentricity, 2.0)))
}

func semimajorAxisFromSemilatusRectumAndEccentricity(semilatusRectum float64, eccentricity float64) float64 {
	return semilatusRectum / (1.0 - math.Pow(eccentricity, 2.0))
}

func semimajorAxisFromApogeeAndPerigee(apogee float64, perigee float64, centralBodyRadius float64) float64 {
	return ((centralBodyRadius * 2.0) + apogee + perigee) / 2.0
}

func eccentricityFromSemimajorAxisAndPerigee(semimajorAxis float64, perigee float64, centralBodyRadius float64) float64 {
	return (semimajorAxis / (centralBodyRadius + perigee)) - 1.0
}
