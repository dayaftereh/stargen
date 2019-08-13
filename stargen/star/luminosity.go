package star

import (
	"math"
)

func luminosity(stellarMassRatio float64) float64 {
	n := 0.5*(2.0-stellarMassRatio) + 4.4
	if stellarMassRatio < 1.0 {
		n = 1.75*(stellarMassRatio-0.1) + 3.325
	}
	return (math.Pow(stellarMassRatio, n))
}

func massToLuminosity(mass float64) float64 {
	if mass <= 0.6224 {
		return 0.3815 * math.Pow(mass, 2.5185)
	}

	if mass <= 1.0 {
		return math.Pow(mass, 4.551)
	}

	if mass <= 3.1623 {
		return math.Pow(mass, 4.351)
	}

	if mass <= 16.0 {
		return 2.7563 * math.Pow(mass, 3.4704)
	}

	return 42.321 * math.Pow(mass, 2.4853)
}
