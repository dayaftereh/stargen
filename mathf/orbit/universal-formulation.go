package orbit

import (
	"math"

	"github.com/dayaftereh/stargen/mathf"
)

func UniversalFormulationZ(x float64, a float64) float64 {
	return math.Pow(x, 2) / a
}

func UniversalFormulationF(x float64, a float64, r *mathf.Vec3, v *mathf.Vec3, mu float64, dt float64) float64 {
	z := UniversalFormulationZ(x, a)
	s := StumpffS(z)
	c := StumpffC(z)

	return (((1.0 - (r.Length() / a)) * s * math.Pow(x, 3)) +
		((r.Dot(v) / math.Sqrt(mu)) * c * math.Pow(x, 2)) +
		(r.Length() * x) -
		(math.Sqrt(mu) * dt))
}

func UniversalFormulationDFDT(x float64, a float64, r *mathf.Vec3, v *mathf.Vec3, mu float64) float64 {
	z := UniversalFormulationZ(x, a)
	s := StumpffS(z)
	c := StumpffC(z)

	return ((c * math.Pow(x, 2.0)) +
		((r.Dot(v) / math.Sqrt(mu)) * (1.0 - (s * z)) * x) +
		(r.Length() * (1.0 - (c * z))))
}

func UniversalFormulationD2FDT(x float64, a float64, r *mathf.Vec3, v *mathf.Vec3, mu float64) float64 {
	z := UniversalFormulationZ(x, a)
	s := StumpffS(z)
	c := StumpffC(z)

	return (((1.0 - (r.Length() / a)) * (1.0 - (s * z)) * x) +
		((r.Dot(v) / math.Sqrt(mu)) * (1.0 - (c * z))))
}
