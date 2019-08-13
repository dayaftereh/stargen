package eccentricity

import (
	"math"

	"github.com/dayaftereh/stargen/mathf/random"
	"github.com/dayaftereh/stargen/stargen/constants"
)

func Rand(random random.Random) float64 {
	E := 1.0 - math.Pow(random.RandFloat64(0.0, 1.0), constants.EccentricityCoefficient)
	if E > 0.9999 {
		E = 0.999
	}
	return E
}
