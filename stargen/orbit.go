package stargen

import (
	"math"

	"github.com/dayaftereh/stargen/mathf/orbit"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

func Orbit(planet *types.Planet, sun *types.Sun) *orbit.Orbit {
	MU := sun.Mass * (constants.SolarMassInGrams / 1000.0) * 6.67430e-11
	sunRadiusMeter := math.Sqrt((sun.Luminosity * 3.828e26) / (4.0 * math.Pi * (5.7e-8) * (math.Pow(sun.EffectiveTemperature, 4.0))))

	semimajorAxis := planet.SemiMajorAxis * constants.KMPerAU * 1000.0
	eccentricity := planet.Eccentricity
	//inclination := 30.0
	//argumentOfPeriapsis := planet.AxialTilt

	param := &orbit.OrbitParameter{
		CentralBodyRadius: &sunRadiusMeter,
		MU:                &MU,
		SemimajorAxis:     &semimajorAxis,
		Eccentricity:      &eccentricity,
		//ArgumentOfPeriapsis: &argumentOfPeriapsis,
	}

	o := orbit.OrbitFromParams(param)
	return o
}

func Orbits(planets []*types.Planet, sun *types.Sun) []*orbit.Orbit {
	orbits := make([]*orbit.Orbit, 0)
	for _, planet := range planets {
		orbits = append(orbits, Orbit(planet, sun))
	}
	return orbits
}
