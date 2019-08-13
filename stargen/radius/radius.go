package radius

import (
	"math"

	"github.com/dayaftereh/stargen/mathf"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

// volumeRadius calculates the radius from the volume. The mass is in units of solar masses, and the density is in units of grams/cc.
// The radius returned is in units of km.
func VolumeRadius(mass float64, density float64) float64 {
	mass = mass * constants.SolarMassInGrams
	volume := mass / density
	return math.Pow((3.0*volume)/(4.0*math.Pi), (1.0/3.0)) / constants.CMPerKM
}

func FractionRadius(mass, imf, rmf, cmf float64) float64 {
	mass /= constants.SunMassInEarthMasses
	iceFraction := imf * constants.IceDensity
	carbonFraction := (cmf * rmf) * constants.CarbonDensity
	silicateFraction := (rmf - (cmf * rmf)) * constants.RockDensity
	ironFraction := (1.0 - (rmf + imf)) * constants.IronDensity
	density := iceFraction + silicateFraction + carbonFraction + ironFraction
	radius := VolumeRadius(mass, density) / constants.EarthRadiusInKM
	return radius
}

func RangeAdjust(x, y1, y2, lower, upper float64) float64 {
	r := upper - lower
	upperFraction := (x - lower) / r
	lowerFraction := 1.0 - upperFraction
	return (lowerFraction * y1) + (upperFraction * y2)
}

func FudgedRadius(planet *types.Planet) float64 {
	mass := planet.Mass * constants.SunMassInEarthMasses

	imf := planet.IceMassFraction
	rmf := planet.RockMassFraction
	cmf := planet.CarbonMassFraction
	var iceRockRadius, nonIceRockRadius float64
	if rmf <= 0.5 {
		r := 0.5 - 0.0
		upperFraction := rmf / r
		lowerFraction := 1.0 - upperFraction
		nonIceRockRadius = (upperFraction * halfRockHalfIronRadius(mass, cmf)) + (lowerFraction + ironRadius(mass))
	} else {
		r := 1.0 - 0.5
		rmf += mathf.QuadTrend(-3.0, 4.5, -1.5, rmf)
		upperFraction := rmf / r
		lowerFraction := 1.0 - upperFraction
		nonIceRockRadius = (upperFraction * rockRadius(mass, cmf)) + (lowerFraction + halfRockHalfIronRadius(mass, cmf))
	}

	if imf <= 0.5 {
		r := 0.5 - 0.0
		upperFraction := imf / r
		lowerFraction := 1.0 - upperFraction
		iceRockRadius = (upperFraction * halfRockHalfWaterRadius(mass, cmf)) + (lowerFraction * rockRadius(mass, cmf))
	} else if imf <= 0.75 {
		r := 0.75 - 0.5
		upperFraction := (imf - 0.5) / r
		lowerFraction := 1.0 - upperFraction
		iceRockRadius = (upperFraction * oneQuaterRockThreeFourthsWaterRadius(mass, cmf)) + (lowerFraction * halfRockHalfWaterRadius(mass, cmf))
	} else {
		r := 1.0 - 0.75
		upperFraction := (imf - 0.75) / r
		lowerFraction := 1.0 - upperFraction
		iceRockRadius = (upperFraction * waterRadius(mass)) + (lowerFraction * oneQuaterRockThreeFourthsWaterRadius(mass, cmf))
	}

	radius := (iceRockRadius * imf) + (nonIceRockRadius * (1.0 - imf))
	radius *= constants.EarthRadiusInKM

	return radius
}

func RadiusImproved(mass float64, planet *types.Planet) float64 {
	mass = mass * constants.SunMassInEarthMasses

	imf := planet.IceMassFraction
	rmf := planet.RockMassFraction
	cmf := planet.CarbonMassFraction

	ironRatio := 0.0
	if imf < 1.0 {
		ironRatio = (1.0 - imf - rmf) / (1.0 - rmf)
	}

	if mathf.CloseZero(mass) {
		mass = constants.ProtoPlanetMass
	}

	nonIceRockRadiiZero := ironRadius(mass)
	nonIceRockRadiiHalf := halfRockHalfIronRadius(mass, cmf)
	nonIceRockRadiiFull := rockRadius(mass, cmf)

	if mathf.CloseZero(imf) {
		var radius float64
		if rmf < 0.5 {
			radius = PlanetRadiusHelper(rmf, 0.0, nonIceRockRadiiZero, 0.5, nonIceRockRadiiHalf, 1.0, nonIceRockRadiiFull)
		} else {
			radius1 := PlanetRadiusHelper(rmf, 0.0, nonIceRockRadiiZero, 0.5, nonIceRockRadiiHalf, 1.0, nonIceRockRadiiFull)
			radius2 := PlanetRadiusHelper2(rmf, 0.5, nonIceRockRadiiHalf, 1.0, nonIceRockRadiiFull)
			radius = RangeAdjust(rmf, radius1, radius2, 0.5, 1.0)
		}

		return radius * constants.EarthRadiusInKM
	}

	iceRockRadiiZero := rockRadius(mass, cmf)
	iceRockRadiiHalf := halfRockHalfWaterRadius(mass, cmf)
	iceRockRadiiQuater := oneQuaterRockThreeFourthsWaterRadius(mass, cmf)
	iceRockRadiiFull := waterRadius(mass)

	var iceRockRadius float64
	if imf < 0.5 {
		iceRockRadius = PlanetRadiusHelper(imf, 0.0, iceRockRadiiZero, 0.5, iceRockRadiiHalf, 0.75, iceRockRadiiQuater)
	} else if imf < 0.75 {
		radius1 := PlanetRadiusHelper(imf, 0.0, iceRockRadiiZero, 0.5, iceRockRadiiHalf, 0.75, iceRockRadiiQuater)
		radius2 := PlanetRadiusHelper(imf, 0.5, iceRockRadiiHalf, 0.75, iceRockRadiiQuater, 1.0, iceRockRadiiFull)
		iceRockRadius = RangeAdjust(imf, radius1, radius2, 0.5, 0.75)
	} else {
		radius1 := PlanetRadiusHelper(imf, 0.5, iceRockRadiiHalf, 0.75, iceRockRadiiQuater, 1.0, iceRockRadiiFull)
		radius2 := PlanetRadiusHelper2(imf, 0.75, iceRockRadiiQuater, 1.0, iceRockRadiiFull)
		iceRockRadius = RangeAdjust(imf, radius1, radius2, 0.5, 0.75)
	}

	iceIronRadiiZero := ironRadius(mass)
	iceIronRadii047 := solid0point953Iron0point047WaterRadius(mass)
	iceIronRadii49 := solid0point51Iron0point49WaterRadius(mass)
	iceIronRadii736 := solid0point264Iron0point736WaterRadius(mass)
	iceIronRadiiFull := iceRockRadiiFull

	var iceIronRadius float64
	if imf < 0.047 {
		iceIronRadius = PlanetRadiusHelper(imf, 0.0, iceIronRadiiZero, 0.047, iceIronRadii047, 0.49, iceIronRadii49)
	} else if imf < 0.49 {
		radius1 := PlanetRadiusHelper(imf, 0.0, iceIronRadiiZero, 0.047, iceIronRadii047, 0.49, iceIronRadii49)
		radius2 := PlanetRadiusHelper(imf, 0.047, iceIronRadii047, 0.49, iceIronRadii49, 0.736, iceIronRadii736)
		iceIronRadius = RangeAdjust(imf, radius1, radius2, 0.047, 0.49)
	} else if imf < 0.736 {
		radius1 := PlanetRadiusHelper(imf, 0.047, iceIronRadii047, 0.49, iceIronRadii49, 0.736, iceIronRadii736)
		radius2 := PlanetRadiusHelper(imf, 0.49, iceIronRadii49, 0.736, iceIronRadii736, 1.0, iceIronRadiiFull)
		iceIronRadius = RangeAdjust(imf, radius1, radius2, 0.49, 0.736)
	} else {
		radius1 := PlanetRadiusHelper(imf, 0.49, iceIronRadii49, 0.736, iceIronRadii736, 1.0, iceIronRadiiFull)
		radius2 := PlanetRadiusHelper2(imf, 0.736, iceIronRadii736, 1.0, iceIronRadiiFull)
		iceIronRadius = RangeAdjust(imf, radius1, radius2, 0.736, 1.0)
	}

	halfMassFactor := 1.0
	if mass > 0.0 {
		average := (nonIceRockRadiiZero + nonIceRockRadiiFull) / 2.0
		halfMassFactor = nonIceRockRadiiHalf / average
	}

	averageIceRockRadius := ((iceRockRadius + iceIronRadius) / 2.0) * halfMassFactor
	radius := PlanetRadiusHelper(ironRatio, 0.0, iceRockRadius, 0.5, averageIceRockRadius, 1.0, iceIronRadius)
	return radius * constants.EarthRadiusInKM
}

func PlanetRadiusHelper2(planetMass, mass1, radius1, mass2, radius2 float64) float64 {
	a, b := mathf.LogFix(mass1, radius1, mass2, radius2)
	radius := mathf.LnTrend(a, b, planetMass)
	return radius
}

func PlanetRadiusHelper(planetMass, mass1, radius1, mass2, radius2, mass3, radius3 float64) float64 {
	a, b, c := mathf.QuadFix(mass1, radius1, mass2, radius2, mass3, radius3)
	radius := mathf.QuadTrend(a, b, c, planetMass)
	return radius
}

func PlanetRadiusHelper3(temperature, temperature1, radius1, temperature2, radius2 float64) float64 {
	adjustedTemperature := temperature / 1000.0
	adjustedTemperature1 := temperature1 / 1000.0
	adjustedTemperature2 := temperature2 / 1000.0
	a, b := mathf.EFix(adjustedTemperature1, radius1, adjustedTemperature2, radius2)
	radius := mathf.ETrend(a, b, adjustedTemperature)
	return radius
}

func CalculateLuminosity(planet *types.Planet, sun *types.Sun) float64 {
	starLuminosity := sun.Luminosity
	return math.Pow(1.0/planet.SemiMajorAxis, 2.0) * starLuminosity
}
