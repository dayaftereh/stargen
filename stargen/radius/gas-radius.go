package radius

import (
	"math"

	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

func MiniNeptuneRadius(planet *types.Planet, sun *types.Sun) float64 {
	coreRadius := planet.CoreRadius

	if coreRadius <= 0.0 {
		coreRadius = RadiusImproved(planet.DustMass, planet)
		planet.CoreRadius = coreRadius
	}

	coreRadiusEU := coreRadius / constants.EarthRadiusInKM
	flux := CalculateLuminosity(planet, sun)
	envRadiusEU := 2.06 * math.Pow(planet.Mass*constants.SunMassInEarthMasses, -0.21) * math.Pow((planet.GasMass/planet.Mass)/0.05, 0.59) * math.Pow(flux, 0.044) * math.Pow(sun.Age/5.0e9, -0.18)
	totalRadiusEU := coreRadiusEU + envRadiusEU
	return totalRadiusEU * constants.EarthRadiusInKM
}

func GasRadius(planet *types.Planet, sun *types.Sun) float64 {
	ageRadii300e6 := gasRadius300Myr(planet)
	ageRadii1e9 := gasRadius1Gyr(planet)
	ageRadii45e9 := gasRadius4point5Gyr(planet)
	var jupiterRadii float64
	if sun.Age < 1.0e9 {
		jupiterRadii = PlanetRadiusHelper(sun.Age, 300.0e6, ageRadii300e6, 1.0e9, ageRadii1e9, 4.5e9, ageRadii45e9)
	} else if sun.Age < 4.5e9 {
		jupiterRadii1 := PlanetRadiusHelper(sun.Age, 300.0e6, ageRadii300e6, 1.0e9, ageRadii1e9, 4.5e9, ageRadii45e9)
		jupiterRadii2 := PlanetRadiusHelper2(sun.Age, 1.0e9, ageRadii1e9, 4.5e9, ageRadii45e9)
		jupiterRadii = RangeAdjust(sun.Age, jupiterRadii1, jupiterRadii2, 1.0e9, 4.5e9)
	} else {
		jupiterRadii = PlanetRadiusHelper2(sun.Age, 1.0e9, ageRadii1e9, 4.5e9, ageRadii45e9)
	}

	radius := jupiterRadii * constants.JupiterRadius
	totalEarthMasseses := planet.Mass * constants.SunMassInEarthMasses
	if totalEarthMasseses < 10.0 {
		r := 10.0 - 0.0
		upperFraction := math.Pow((10.0-totalEarthMasseses)/r, 2.0)
		lowerFraction := 1.0 - upperFraction
		radius = (lowerFraction * radius) + (upperFraction * gasDwarfRadius(planet))
	}

	return radius
}
