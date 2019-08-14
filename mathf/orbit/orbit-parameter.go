package orbit

import (
	"fmt"
)

type OrbitParameter struct {
	MU                *float64 // G*M earth MU
	
	CentralBodyRadius *float64 // // earth radius

	// Apogee defines furthest point to the central body
	Apogee *float64 // km

	// Perigee defines closest point to the central body
	Perigee *float64 // km

	// SemimajorAxis is the sum of the periapsis and apoapsis distances divided by two.
	// For circular orbits, the semimajor axis is the distance between the centers of the bodies, not the distance of the bodies from the center of mass.
	SemimajorAxis *float64

	SemilatusRectum *float64

	// Eccentricity determines the amount by which its orbit around another body deviates from a perfect circle.
	// A value of 0 is a circular orbit, values between 0 and 1 form an elliptic orbit, 1 is a parabolic escape orbit, and greater than 1 is a hyperbola.
	Eccentricity *float64

	// Inclination determines the vertical tilt of the ellipse with respect to the reference plane, measured at the ascending node
	Inclination *float64 // deg

	// ArgumentOfPeriapsis defines the orientation of the ellipse in the orbital plane, as an angle measured from the ascending node to the periapsis
	ArgumentOfPeriapsis *float64 // deg

	// RightAscension horizontally orients the ascending node of the ellipse with respect to the reference frame's vernal point
	RightAscension *float64 // deg

	// TrueAnomaly defines the position of the orbiting body along the ellipse at a specific time (epoch)
	TrueAnomaly     *float64 // deg
	AngularMomentum *float64
}

func (p *OrbitParameter) String() string {
	s := fmt.Sprintf(" MU: %f\n", *p.MU)
	s = fmt.Sprintf("%s CentralBodyRadius: %f\n", s, *p.CentralBodyRadius)
	s = fmt.Sprintf("%s Apogee: %v\n", s, p.Apogee)
	s = fmt.Sprintf("%s Perigee: %v\n", s, p.Perigee)
	s = fmt.Sprintf("%s SemimajorAxis: %v\n", s, *p.SemimajorAxis)
	s = fmt.Sprintf("%s SemilatusRectum: %f\n", s, *p.SemilatusRectum)
	s = fmt.Sprintf("%s Eccentricity: %f\n", s, *p.Eccentricity)
	s = fmt.Sprintf("%s Inclination: %f\n", s, *p.Inclination)
	s = fmt.Sprintf("%s ArgumentOfPeriapsis: %f\n", s, *p.ArgumentOfPeriapsis)
	s = fmt.Sprintf("%s RightAscension: %f\n", s, *p.RightAscension)
	s = fmt.Sprintf("%s TrueAnomaly: %f\n", s, *p.TrueAnomaly)
	s = fmt.Sprintf("%s AngularMomentum: %v\n", s, *p.AngularMomentum)
	return s
}

func extendParams(params *OrbitParameter) *OrbitParameter {
	params = baseParams(params)

	if params.Apogee != nil && params.Perigee != nil {
		semimajorAxis := semimajorAxisFromApogeeAndPerigee(*params.Apogee, *params.Perigee, *params.CentralBodyRadius)
		eccentricity := eccentricityFromSemimajorAxisAndPerigee(semimajorAxis, *params.Perigee, *params.CentralBodyRadius)

		params.SemimajorAxis = &semimajorAxis
		params.Eccentricity = &eccentricity
	} else if params.SemilatusRectum != nil {
		semimajorAxis := semimajorAxisFromSemilatusRectumAndEccentricity(*params.SemilatusRectum, *params.Eccentricity)
		params.SemimajorAxis = &semimajorAxis
	}

	if params.SemilatusRectum == nil {
		semilatusRectum := semilatusRectumFromSemimajorAxisAndEccentricity(*params.SemimajorAxis, *params.Eccentricity)
		params.SemilatusRectum = &semilatusRectum
	}

	angularMomentum := angularMomentumFromSemilatusRectum(*params.SemilatusRectum, *params.MU)
	params.AngularMomentum = &angularMomentum

	return params
}

func baseParams(params *OrbitParameter) *OrbitParameter {
	zero := 0.0
	if params.Inclination == nil {
		params.Inclination = &zero
	}
	if params.ArgumentOfPeriapsis == nil {
		params.ArgumentOfPeriapsis = &zero
	}
	if params.RightAscension == nil {
		params.RightAscension = &zero
	}
	if params.TrueAnomaly == nil {
		params.TrueAnomaly = &zero
	}
	if params.Eccentricity == nil {
		params.Eccentricity = &zero
	}
	return params
}
