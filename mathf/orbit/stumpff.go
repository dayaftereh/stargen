package orbit

import "math"

func StumpffC(z float64) float64 {
	if z > 0.0 {
		return (1.0 - math.Cos(math.Sqrt(z))) / z
	}
	if z < 0.0 {
		return (math.Cosh(math.Sqrt(-z)) - 1) / (-z)
	}
	return 0.5
}

func StumpffS(z float64) float64 {
	if z > 0.0 {
		return (math.Sqrt(z) - math.Sin(math.Sqrt(z))) / math.Pow(math.Sqrt(z), 3.0)
	}
	if z < 0.0 {
		return (math.Sinh(math.Sqrt(-z)) - math.Sqrt(-z)) / math.Pow(math.Sqrt(-z), 3.0)
	}
	return 1.0 / 6.0
}
