package orbit

import (
	"math"

	"github.com/dayaftereh/stargen/mathf"
)

func LagrangeF(x float64, z float64, r *mathf.Vec3) float64 {
	c := StumpffC(z)
	return 1.0 - (math.Pow(x, 2.0)/r.Length())*c
}

func LagrangeG(x float64, z float64, mu float64, dt float64) float64 {
	s := StumpffS(z)
	return dt - ((1.0 / math.Sqrt(mu)) * math.Pow(x, 3) * s)
}

func LagrangeDf(x float64, z float64, r *mathf.Vec3, r0 *mathf.Vec3, mu float64) float64 {
	s := StumpffS(z)
	return (math.Sqrt(mu) / (r.Length() * r0.Length())) * (s*z - 1.0) * x
}

func LagrangeDg(x float64, z float64, r *mathf.Vec3) float64 {
	c := StumpffC(z)
	return 1.0 - ((math.Pow(x, 2.0) / r.Length()) * c)
}
