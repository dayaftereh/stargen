package mathf

import (
	"math"

	"github.com/dayaftereh/discover/server/utils"
)

func CloseZero(x float64) bool {
	return math.Abs(x) < Epsilon
}

// about returns a value within a certain variation of the exact value given it in 'value'.
func About(value float64, variation float64) float64 {
	return (value + (value * utils.RandFloat64(-variation, variation)))
}

func QuadFix(x, y, w, z, p, q float64) (float64, float64, float64) {
	a := ((q * (w - x)) - (w * y) + (p * (y - z)) + (x * z)) / ((p - w) * (p - x) * (w - x))
	b := ((q * (math.Pow(x, 2.0) - math.Pow(w, 2.0))) + (math.Pow(w, 2.0) * y) - (math.Pow(x, 2.0) * z) + (math.Pow(p, 2.0) * (z - y))) / ((p - w) * (p - x) * (w - x))
	c := ((q * w * x * (w - x)) + (p * ((p * w * y) - (math.Pow(w, 2) * y) - (p * x * z) + (math.Pow(x, 2) * z)))) / ((p - w) * (p - x) * (w - x))
	return a, b, c
}

func QuadTrend(a, b, c, x float64) float64 {
	return (a * math.Pow(x, 2.0)) + (b * x) + c
}

func LogFix(x, y, w, z float64) (float64, float64) {
	a := ((y * math.Log(w)) - (z * math.Log(x))) / (math.Log(w) - math.Log(x))
	b := (z - y) / (math.Log(w) - math.Log(x))
	return a, b
}

func LnTrend(a, b, x float64) float64 {
	return a + (b * math.Log(x))
}

func ETrend(a, b, x float64) float64 {
	return a + (b * math.Exp(x))
}

func EFix(x, y, w, z float64) (float64, float64) {
	a := ((math.Exp(x) * z) - (math.Exp(w) * y)) / (math.Exp(x) - math.Exp(w))
	b := (y - z) / (math.Exp(x) - math.Exp(w))
	return a, b
}

func Clamp(v float64, min float64, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}

func ToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func ToDegress(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}
