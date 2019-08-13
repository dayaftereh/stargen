package inertia

import "math"

func CalculateMomentOfInertiaCoefficient(mass, radius float64) float64 {
	return CalculateMomentOfInertia(mass, radius) / (mass * math.Pow(radius, 2.0))
}

func CalculateMomentOfInertia(mass, radius float64) float64 {
	return (2.0 / 5.0) * mass * math.Pow(radius, 2.0)
}

func AVE(x, y float64) float64 {
	return (x + y) / 2.0
}

func GetSpinResonanceFactor(eccentricity float64) float64 {
	top := 1.0 - eccentricity
	bottom := 1.0 + eccentricity
	fraction := top / bottom

	if fraction <= AVE(1.0/10.0, 1.0/9.0) {
		return 1.0 / 10.0
	} else if fraction <= AVE(1.0/9.0, 1.0/8.0) {
		return 1.0 / 9.0
	} else if fraction <= AVE(1.0/8.0, 1.0/7.0) {
		return 1.0 / 8.0
	} else if fraction <= AVE(1.0/7.0, 1.0/6.0) {
		return 1.0 / 7.0
	} else if fraction <= AVE(1.0/6.0, 1.0/5.0) {
		return 1.0 / 6.0
	} else if fraction <= AVE(1.0/5.0, 2.0/9.0) {
		return 1.0 / 5.0
	} else if fraction <= AVE(2.0/9.0, 1.0/4.0) {
		return 2.0 / 9.0
	} else if fraction <= AVE(1.0/4.0, 2.0/7.0) {
		return 1.0 / 4.0
	} else if fraction <= AVE(2.0/7.0, 3.0/10.0) {
		return 2.0 / 7.0
	} else if fraction <= AVE(3.0/10.0, 1.0/3.0) {
		return 3.0 / 10.0
	} else if fraction <= AVE(1.0/3.0, 3.0/8.0) {
		return 1.0 / 3.0
	} else if fraction <= AVE(3.0/8.0, 2.0/5.0) {
		return 3.0 / 8.0
	} else if fraction <= AVE(2.0/5.0, 3.0/7.0) {
		return 2.0 / 5.0
	} else if fraction <= AVE(3.0/7.0, 4.0/9.0) {
		return 3.0 / 7.0
	} else if fraction <= AVE(4.0/9.0, 1.0/2.0) {
		return 4.0 / 9.0
	} else if fraction <= AVE(1.0/2.0, 5.0/9.0) {
		return 1.0 / 2.0
	} else if fraction <= AVE(5.0/9.0, 4.0/7.0) {
		return 5.0 / 9.0
	} else if fraction <= AVE(4.0/7.0, 3.0/5.0) {
		return 4.0 / 7.0
	} else if fraction <= AVE(3.0/5.0, 5.0/8.0) {
		return 3.0 / 5.0
	} else if fraction <= AVE(5.0/8.0, 2.0/3.0) {
		return 5.0 / 8.0
	} else if fraction <= AVE(2.0/3.0, 7.0/10.0) {
		return 2.0 / 3.0
	} else if fraction <= AVE(7.0/10.0, 5.0/7.0) {
		return 7.0 / 10.0
	} else if fraction <= AVE(5.0/7.0, 3.0/4.0) {
		return 5.0 / 7.0
	} else if fraction <= AVE(3.0/4.0, 7.0/9.0) {
		return 3.0 / 4.0
	} else if fraction <= AVE(7.0/9.0, 4.0/5.0) {
		return 7.0 / 9.0
	} else if fraction <= AVE(4.0/5.0, 5.0/6.0) {
		return 4.0 / 5.0
	} else if fraction <= AVE(5.0/6.0, 6.0/7.0) {
		return 5.0 / 6.0
	} else if fraction <= AVE(6.0/7.0, 7.0/8.0) {
		return 6.0 / 7.0
	} else if fraction <= AVE(7.0/8.0, 8.0/9.0) {
		return 7.0 / 8.0
	} else if fraction <= AVE(8.0/9.0, 9.0/10.0) {
		return 8.0 / 9.0
	} else {
		return 9.0 / 10.0
	}
}
