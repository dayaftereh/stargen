package orbit

import (
	"math"

	"github.com/dayaftereh/stargen/mathf"
)

type LaguerreFunction func(x float64) float64

func LaguerreSolve(guess float64, f LaguerreFunction, df LaguerreFunction, d2f LaguerreFunction) float64 {
	x := guess

	for n := 0; n <= 10; n++ {
		f := f(x)
		fPrime := df(x)
		fPrimePrime := d2f(x)

		delta := 2.0 * math.Sqrt(math.Abs(
			(4.0*math.Pow(fPrime, 2.0))-
				(5.0*f*fPrimePrime)))

		dx := (5.0 * f) / (fPrime + ((math.Abs(fPrime) / fPrime) * delta))

		x = x - dx

		if mathf.CloseZero(dx) {
			break
		}
	}

	return x
}
