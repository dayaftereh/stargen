package radius

import (
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

func gasRadius4point5Gyr1960K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-0.00117641848, 0.0813324707, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 28.0, 1.355, 46.0, 1.252)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(7.1355424E-5, -0.0110025236, 1.607128008, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 28.0, 1.355, 46.0, 1.252)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.355, 46.0, 1.252, 77.0, 1.183)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(2.8438817E-5, -0.005723781, 1.455117388, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.355, 46.0, 1.252, 77.0, 1.183)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.252, 77.0, 1.183, 129.0, 1.19)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.059734E-6, 3.5292059E-4, 1.162108278, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.252, 77.0, 1.183, 129.0, 1.19)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.183, 129.0, 1.19, 215.0, 1.189)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.521665E-7, 1.4391737E-4, 1.178959162, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.183, 129.0, 1.19, 215.0, 1.189)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.19, 215.0, 1.189, 318.0, 1.179)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(2.523727E-7, -2.31602E-4, 1.227128508, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.19, 215.0, 1.189, 318.0, 1.179)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.189, 318.0, 1.179, 464.0, 1.174)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(4.6805591E-8, -7.084855E-5, 1.19679668, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.189, 318.0, 1.179, 464.0, 1.174)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.179, 464.0, 1.174, 774.0, 1.17)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(3.4235799E-8, -5.528714E-5, 1.192282405, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.179, 464.0, 1.174, 774.0, 1.17)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.174, 774.0, 1.17, 1292.0, 1.178)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.296037E-8, 6.2880146E-5, 1.135085775, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.174, 774.0, 1.17, 1292.0, 1.178)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.17, 1292.0, 1.178, 2154.0, 1.164)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.821523E-9, 7.2656677E-6, 1.179999679, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.17, 1292.0, 1.178, 2154.0, 1.164)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.178, 2154.0, 1.164, 3594.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)

	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.178, 2154.0, 1.164, 3594.0, 1.118)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.164, 3594.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.853640017, -0.0898544185, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.164, 3594.0, 1.118)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1960K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.726, 0.934};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.726, 28.0, 0.934)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-4.892024E-4, 0.0409231975, 0.1716851271, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.726, 28.0, 0.934)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.726, 28.0, 0.934, 46.0, 1.019)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.148051E-5, 0.0092717797, 0.7225908858, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.726, 28.0, 0.934, 46.0, 1.019)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.934, 46.0, 1.019, 77.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-8.782026E-6, 0.0027898667, 0.9092489013, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.934, 46.0, 1.019, 77.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.019, 77.0, 1.072, 129.0, 1.123)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.000519E-6, 0.002010876, 0.9468106187, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.019, 77.0, 1.072, 129.0, 1.123)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.072, 129.0, 1.123, 215.0, 1.148)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.281238E-6, 7.3144355E-4, 1.049964864, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.072, 129.0, 1.123, 215.0, 1.148)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.123, 215.0, 1.148, 318.0, 1.153)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-8.492542E-8, 9.3808937E-5, 1.131756756, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.123, 215.0, 1.148, 318.0, 1.153)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.148, 318.0, 1.153, 464.0, 1.157)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.88593E-8, 5.7785233E-5, 1.138553904, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.148, 318.0, 1.153, 464.0, 1.157)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.153, 464.0, 1.157, 774.0, 1.16)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(1.6290584E-8, -1.049032E-5, 1.158360213, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.153, 464.0, 1.157, 774.0, 1.16)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.157, 774.0, 1.16, 1292.0, 1.172)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.687474E-8, 7.868923E-5, 1.115194546, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.157, 774.0, 1.16, 1292.0, 1.172)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.16, 1292.0, 1.172, 2154.0, 1.16)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.226082E-9, 1.0979967E-5, 1.169876123, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.16, 1292.0, 1.172, 2154.0, 1.16)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.172, 2154.0, 1.16, 3594.0, 1.116)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.172, 2154.0, 1.16, 3594.0, 1.116)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.16, 3594.0, 1.116)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.819655669, -0.0859477047, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.16, 3594.0, 1.116)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1960K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.430, 0.756};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.43, 46.0, 0.756)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.563821E-4, 0.0370833882, -0.4073312852, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.43, 46.0, 0.756)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.43, 46.0, 0.756, 77.0, 0.928)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-4.275165E-5, 0.0108068403, 0.349347843, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.43, 46.0, 0.756, 77.0, 0.928)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.756, 77.0, 0.928, 129.0, 1.032)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.521402E-6, 0.0039614088, 0.679423913, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.756, 77.0, 0.928, 129.0, 1.032)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.928, 129.0, 1.032, 215.0, 1.091)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.345651E-6, 0.0014929505, 0.8784433657, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.928, 129.0, 1.032, 215.0, 1.091)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.032, 215.0, 1.091, 318.0, 1.116)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.621635E-7, 5.4325161E-4, 1.000380413, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.032, 215.0, 1.091, 318.0, 1.116)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.091, 318.0, 1.116, 464.0, 1.131)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.262685E-7, 2.0148169E-4, 1.064697598, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.091, 318.0, 1.116, 464.0, 1.131)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.116, 464.0, 1.131, 774.0, 1.145)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.257519E-8, 6.0729373E-5, 1.105528959, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.116, 464.0, 1.131, 774.0, 1.145)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.131, 774.0, 1.145, 1292.0, 1.163)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.190564E-8, 1.0066608E-4, 1.086198356, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.131, 774.0, 1.145, 1292.0, 1.163)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.145, 1292.0, 1.163, 2154.0, 1.155)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-8.940212E-9, 2.1527229E-5, 1.150110395, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.145, 1292.0, 1.163, 2154.0, 1.155)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.163, 2154.0, 1.155, 3594.0, 1.112)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.163, 2154.0, 1.155, 3594.0, 1.112)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.155, 3594.0, 1.112)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.799663495, -0.0839943477, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.155, 3594.0, 1.112)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1960K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.695, 0.891};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.695, 129.0, 0.891)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.779186E-5, 0.0074343548, 0.2280426421, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.695, 129.0, 0.891)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.695, 129.0, 0.891, 215.0, 1.004)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.280948E-6, 0.0027865995, 0.6027679149, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.695, 129.0, 0.891, 215.0, 1.004)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.891, 215.0, 1.004, 318.0, 1.056)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.064772E-6, 0.0010723781, 0.8226578179, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.891, 215.0, 1.004, 318.0, 1.056)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.004, 318.0, 1.056, 464.0, 1.091)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.134909E-7, 4.8487588E-4, 0.9335109194, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.004, 318.0, 1.056, 464.0, 1.091)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.056, 464.0, 1.091, 774.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-5.39259E-8, 1.6353445E-4, 1.026730044, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.056, 464.0, 1.091, 774.0, 1.121)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.091, 774.0, 1.121, 1292.0, 1.148)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.113328E-8, 1.371049E-4, 1.039522764, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.091, 774.0, 1.121, 1292.0, 1.148)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.121, 1292.0, 1.148, 2154.0, 1.144)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-9.447662E-9, 2.7916272E-5, 1.127702819, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.121, 1292.0, 1.148, 2154.0, 1.144)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.148, 2154.0, 1.144, 3594.0, 1.106)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.148, 2154.0, 1.144, 3594.0, 1.106)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.144, 3594.0, 1.106)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.713702623, -0.0742275631, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.144, 3594.0, 1.106)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1960K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.613, 0.841};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.613, 215.0, 0.841)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.736311E-6, 0.0056564538, 0.0286984127, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.613, 215.0, 0.841)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.613, 215.0, 0.841, 318.0, 0.944)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.173076E-6, 0.0021582494, 0.4774268031, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.613, 215.0, 0.841, 318.0, 0.944)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.841, 318.0, 0.944, 464.0, 1.011)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.465498E-7, 8.8630602E-4, 0.7174239831, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.841, 318.0, 0.944, 464.0, 1.011)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.944, 464.0, 1.011, 774.0, 1.076)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.553096E-7, 4.0195069E-4, 0.8579324135, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.944, 464.0, 1.011, 774.0, 1.076)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.011, 774.0, 1.076, 1292.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.286988E-8, 1.9031025E-4, 0.9603729424, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.011, 774.0, 1.076, 1292.0, 1.118)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.076, 1292.0, 1.118, 2154.0, 1.125)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.257775E-8, 5.146358E-5, 1.072504642, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.076, 1292.0, 1.118, 2154.0, 1.125)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.118, 2154.0, 1.125, 3594.0, 1.095)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.118, 2154.0, 1.125, 3594.0, 1.095)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.125, 3594.0, 1.095)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.574765229, -0.0586007077, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.125, 3594.0, 1.095)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1960K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius4point5Gyr1960K0coreMass(planet)
	coreMassRadii10 := gasRadius4point5Gyr1960K10coreMass(planet)
	coreMassRadii25 := gasRadius4point5Gyr1960K25coreMass(planet)
	coreMassRadii50 := gasRadius4point5Gyr1960K50coreMass(planet)
	coreMassRadii100 := gasRadius4point5Gyr1960K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if coreEarthMasses <= 10.0 {
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1300K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0024406035, 0.1063726127, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.103, 28.0, 1.065)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(6.7398119E-5, -0.0064874608, 1.193808777, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.103, 28.0, 1.065)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.103, 28.0, 1.065, 46.0, 1.038)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(3.7853851E-5, -0.004301185, 1.15575576, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.103, 28.0, 1.065, 46.0, 1.038)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.065, 46.0, 1.038, 77.0, 1.049)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(4.2975874E-6, -1.737645E-4, 1.036899474, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.065, 46.0, 1.038, 77.0, 1.049)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.038, 77.0, 1.049, 129.0, 1.086)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.555132E-6, 0.0024438957, 0.9588984114, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.038, 77.0, 1.049, 129.0, 1.086)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.049, 129.0, 1.086, 215.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.066205E-6, 5.8770477E-4, 1.027928803, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.049, 129.0, 1.086, 215.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.086, 215.0, 1.105, 318.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.047454E-8, 4.6320406E-5, 1.097374298, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.086, 215.0, 1.105, 318.0, 1.107)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.105, 318.0, 1.107, 464.0, 1.108)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(2.0350257E-8, -9.064586E-6, 1.107824639, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.105, 318.0, 1.107, 464.0, 1.108)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.107, 464.0, 1.108, 774.0, 1.113)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-7.821887E-9, 2.5812529E-5, 1.097707008, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.107, 464.0, 1.108, 774.0, 1.113)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.108, 774.0, 1.113, 1292.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.296686E-8, 5.7102052E-5, 1.082561909, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.108, 774.0, 1.113, 1292.0, 1.118)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.113, 1292.0, 1.118, 2154.0, 1.099)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.301773E-9, -7.217854E-6, 1.134506262, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.113, 1292.0, 1.118, 2154.0, 1.099)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.118, 2154.0, 1.099, 3594.0, 1.053)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.118, 2154.0, 1.099, 3594.0, 1.053)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.099, 3594.0, 1.053)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.788640017, -0.0898544185, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.099, 3594.0, 1.053)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1300K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.599, 0.775};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.599, 28.0, 0.775)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-3.544061E-4, 0.0319482759, 0.158302682, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.599, 28.0, 0.775)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.599, 28.0, 0.775, 46.0, 0.878)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.016385E-5, 0.0101743472, 0.5372867384, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.599, 28.0, 0.775, 46.0, 0.878)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.775, 46.0, 0.878, 77.0, 0.964)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.836378E-5, 0.0050329382, 0.6853425962, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.775, 46.0, 0.878, 77.0, 0.964)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.878, 77.0, 0.964, 129.0, 1.029)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.687563E-5, 0.002421638, 0.8112554348, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.878, 77.0, 0.964, 129.0, 1.029)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.964, 129.0, 1.029, 215.0, 1.069)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.741767E-6, 0.0010642841, 0.9206920943, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.964, 129.0, 1.029, 215.0, 1.069)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.029, 215.0, 1.069, 318.0, 1.083)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.983072E-7, 2.9492007E-4, 1.019381435, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.029, 215.0, 1.069, 318.0, 1.083)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.069, 318.0, 1.083, 464.0, 1.092)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.029421E-8, 1.0097391E-4, 1.05597625, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.069, 318.0, 1.083, 464.0, 1.092)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.083, 464.0, 1.092, 774.0, 1.104)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-2.809863E-8, 7.3495776E-5, 1.063947482, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.083, 464.0, 1.092, 774.0, 1.104)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.092, 774.0, 1.104, 1292.0, 1.112)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.548231E-8, 6.8090476E-5, 1.066563814, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.092, 774.0, 1.104, 1292.0, 1.112)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.104, 1292.0, 1.112, 2154.0, 1.095)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.008003E-9, -2.464E-6, 1.123543167, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.104, 1292.0, 1.112, 2154.0, 1.095)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.112, 2154.0, 1.095, 3594.0, 1.05)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.112, 2154.0, 1.095, 3594.0, 1.05)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.095, 3594.0, 1.05)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.769647843, -0.0879010616, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.095, 3594.0, 1.05)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1300K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.403, 0.686};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.403, 46.0, 0.686)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.155292E-4, 0.0316713847, -0.3148238607, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.403, 46.0, 0.686)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.403, 46.0, 0.686, 77.0, 0.846)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.762444E-5, 0.0097890968, 0.3153148674, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.403, 46.0, 0.686, 77.0, 0.846)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.686, 77.0, 0.846, 129.0, 0.952)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.126027E-6, 0.0039184232, 0.5983896321, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.686, 77.0, 0.846, 129.0, 0.952)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.846, 129.0, 0.952, 215.0, 1.019)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.529624E-6, 0.0016492603, 0.7813408846, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.846, 129.0, 0.952, 215.0, 1.019)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.952, 215.0, 1.019, 318.0, 1.05)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-6.860799E-7, 6.6665144E-4, 0.9073839815, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.952, 215.0, 1.019, 318.0, 1.05)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.019, 318.0, 1.05, 464.0, 1.069)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.368313E-7, 2.3713903E-4, 0.9884267135, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.019, 318.0, 1.05, 464.0, 1.069)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.05, 464.0, 1.069, 774.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-4.917259E-8, 1.2861761E-4, 1.019908093, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.05, 464.0, 1.069, 774.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.069, 774.0, 1.09, 1292.0, 1.104)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.135386E-8, 9.1804101E-5, 1.037726971, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.069, 774.0, 1.09, 1292.0, 1.104)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.09, 1292.0, 1.104, 2154.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.916513E-9, 4.1470028E-6, 1.108518294, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.09, 1292.0, 1.104, 2154.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.104, 2154.0, 1.09, 3594.0, 1.047)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.104, 2154.0, 1.09, 3594.0, 1.047)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.09, 3594.0, 1.047)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		///jupiterRadii = ln_trend(1.734663495, -0.0839943477, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.09, 3594.0, 1.047)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1300K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.648, 0.831};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.648, 129.0, 0.831)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.614879E-5, 0.0068458816, 0.2166132943, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.648, 129.0, 0.831)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.648, 129.0, 0.831, 215.0, 0.942)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.055163E-6, 0.0026856738, 0.5520300509, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.648, 129.0, 0.831, 215.0, 0.942)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.831, 215.0, 0.942, 318.0, 0.996)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.08774E-6, 0.0011040371, 0.7549127896, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.831, 215.0, 0.942, 318.0, 0.996)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.942, 318.0, 0.996, 464.0, 1.033)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.08161E-7, 4.9440659E-4, 0.8699411819, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.942, 318.0, 0.996, 464.0, 1.033)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.996, 464.0, 1.033, 774.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-8.506302E-8, 2.1821125E-4, 0.9500637093, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.996, 464.0, 1.033, 774.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.033, 774.0, 1.068, 1292.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.845194E-8, 1.2168549E-4, 0.996785166, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.033, 774.0, 1.068, 1292.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.068, 1292.0, 1.09, 2154.0, 1.081)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.229582E-9, 1.4472305E-5, 1.083369863, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.068, 1292.0, 1.09, 2154.0, 1.081)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.081, 3594.0, 1.042)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.081, 3594.0, 1.042)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.081, 3594.0, 1.042)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.665694797, -0.0761809201, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.081, 3594.0, 1.042)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1300K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.587, 0.798};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.587, 215.0, 0.798)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-7.94726E-6, 0.0051873457, 0.0500827554, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.587, 215.0, 0.798)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.587, 215.0, 0.798, 318.0, 0.896)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.033136E-6, 0.0020351177, 0.4544313939, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.587, 215.0, 0.798, 318.0, 0.896)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.798, 318.0, 0.896, 464.0, 0.961)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.165089E-7, 8.4911544E-4, 0.6782127358, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.798, 318.0, 0.896, 464.0, 0.961)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.896, 464.0, 0.961, 774.0, 1.026)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.692987E-7, 4.1926925E-4, 0.8029084081, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.896, 464.0, 0.961, 774.0, 1.026)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.961, 774.0, 1.026, 1292.0, 1.062)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.952027E-8, 1.7180685E-4, 0.9226878251, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.961, 774.0, 1.026, 1292.0, 1.062)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.026, 1292.0, 1.062, 2154.0, 1.063)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-9.855721E-9, 3.5122909E-5, 1.033073003, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.026, 1292.0, 1.062, 2154.0, 1.063)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.062, 2154.0, 1.063, 3594.0, 1.032)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.062, 2154.0, 1.063, 3594.0, 1.032)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.063, 3594.0, 1.032)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.527757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.063, 3594.0, 1.032)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr1300K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius4point5Gyr1300K0coreMass(planet)
	coreMassRadii10 := gasRadius4point5Gyr1300K10coreMass(planet)
	coreMassRadii25 := gasRadius4point5Gyr1300K25coreMass(planet)
	coreMassRadii50 := gasRadius4point5Gyr1300K50coreMass(planet)
	coreMassRadii100 := gasRadius4point5Gyr1300K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if coreEarthMasses <= 10.0 {
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr875K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0023768144, 0.1032293736, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.068, 28.0, 1.027)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(8.6381052E-5, -0.0076144201, 1.172481017, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.068, 28.0, 1.027)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.068, 28.0, 1.027, 46.0, 1.005)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(3.745154E-5, -0.0039936362, 1.109459805, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.068, 28.0, 1.027, 46.0, 1.005)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.027, 46.0, 1.005, 77.0, 1.024)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(1.4200723E-6, 4.3823433E-4, 0.9818363479, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.027, 46.0, 1.005, 77.0, 1.024)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.005, 77.0, 1.024, 129.0, 1.062)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.357445E-6, 0.0014224028, 0.9343812709, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.005, 77.0, 1.024, 129.0, 1.062)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.024, 129.0, 1.062, 215.0, 1.085)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.158191E-6, 6.658597E-4, 0.995377562, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.024, 129.0, 1.062, 215.0, 1.085)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.062, 215.0, 1.085, 318.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.3994E-7, 1.2313171E-4, 1.064995409, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.062, 215.0, 1.085, 318.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.085, 318.0, 1.09, 464.0, 1.092)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(1.9478103E-8, -1.533247E-6, 1.088517869, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.085, 318.0, 1.09, 464.0, 1.092)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.09, 464.0, 1.092, 774.0, 1.099)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.561369E-8, 4.1910394E-5, 1.075915142, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.09, 464.0, 1.092, 774.0, 1.099)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.092, 774.0, 1.099, 1292.0, 1.104)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.380751E-8, 5.8838828E-5, 1.067721256, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.092, 774.0, 1.099, 1292.0, 1.104)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.099, 1292.0, 1.104, 2154.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-3.797823E-9, -1.011456E-5, 1.123407579, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.099, 1292.0, 1.104, 2154.0, 1.084)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.104, 2154.0, 1.084, 3594.0, 1.038)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.104, 2154.0, 1.084, 3594.0, 1.038)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.084, 3594.0, 1.038)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.773640017, -0.0898544185, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.084, 3594.0, 1.038)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr875K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.592, 0.755};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.592, 28.0, 0.755)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-3.136538E-4, 0.0289326019, 0.1907917102, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.592, 28.0, 0.755)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.592, 28.0, 0.755, 46.0, 0.858)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.148051E-5, 0.0102717797, 0.5155908858, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.592, 28.0, 0.755, 46.0, 0.858)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.755, 46.0, 0.858, 77.0, 0.942)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.735478E-5, 0.0048443152, 0.6718842118, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.755, 46.0, 0.858, 77.0, 0.942)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.858, 77.0, 0.942, 129.0, 1.008)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.574136E-6, 0.0024175028, 0.7889013378, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.858, 77.0, 0.942, 129.0, 1.008)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.942, 129.0, 1.008, 215.0, 1.051)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.823599E-6, 0.001127318, 0.8929224842, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.942, 129.0, 1.008, 215.0, 1.051)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.008, 215.0, 1.051, 318.0, 1.067)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-3.487817E-7, 3.4124048E-4, 0.9937557337, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.008, 215.0, 1.051, 318.0, 1.067)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.051, 318.0, 1.067, 464.0, 1.077)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.82405E-8, 1.1403722E-4, 1.036625676, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.051, 318.0, 1.067, 464.0, 1.077)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.067, 464.0, 1.077, 774.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-3.199453E-8, 8.1544708E-5, 1.046051549, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.067, 464.0, 1.077, 774.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.077, 774.0, 1.09, 1292.0, 1.098)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.632296E-8, 6.9827253E-5, 1.051723161, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.077, 774.0, 1.09, 1292.0, 1.098)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.09, 1292.0, 1.098, 2154.0, 1.08)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.202383E-9, -6.40026E-6, 1.113284022, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.09, 1292.0, 1.098, 2154.0, 1.08)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.098, 2154.0, 1.08, 3594.0, 1.036)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.098, 2154.0, 1.08, 3594.0, 1.036)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.08, 3594.0, 1.036)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.739655669, -0.0859477047, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.08, 3594.0, 1.036)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr875K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.404, 0.675};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.404, 46.0, 0.675)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.058737E-4, 0.0302902129, -0.2827209421, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.404, 46.0, 0.675)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.404, 46.0, 0.675, 77.0, 0.829)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.552423E-5, 0.0093372223, 0.3206570451, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.404, 46.0, 0.675, 77.0, 0.829)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.675, 77.0, 0.829, 129.0, 0.934)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-8.902414E-6, 0.003853128, 0.5850915552, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.675, 77.0, 0.829, 129.0, 0.934)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.829, 129.0, 0.934, 215.0, 1.002)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.539778E-6, 0.0016643813, 0.7615592541, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.829, 129.0, 0.934, 215.0, 1.002)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.934, 215.0, 1.002, 318.0, 1.034)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-6.975635E-7, 6.8248095E-4, 0.8875114673, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.934, 215.0, 1.002, 318.0, 1.034)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.002, 318.0, 1.034, 464.0, 1.054)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.377034E-7, 2.4467036E-4, 0.9701199433, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.002, 318.0, 1.034, 464.0, 1.054)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.034, 464.0, 1.054, 774.0, 1.077)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-5.929592E-8, 1.476019E-4, 0.9982788934, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.034, 464.0, 1.054, 774.0, 1.077)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.054, 774.0, 1.077, 1292.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.079559E-8, 8.8720219E-5, 1.026779451, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.054, 774.0, 1.077, 1292.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.077, 1292.0, 1.09, 2154.0, 1.075)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.110893E-9, 2.1074347E-7, 1.098259148, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.077, 1292.0, 1.09, 2154.0, 1.075)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.075, 3594.0, 1.033)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.075, 3594.0, 1.033)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.075, 3594.0, 1.033)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.70467132, -0.0820409908, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.075, 3594.0, 1.033)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr875K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.639, 0.817};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.639, 129.0, 0.817)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.545202E-5, 0.0066061938, 0.2219381271, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.639, 129.0, 0.817)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.639, 129.0, 0.817, 215.0, 0.928)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.055163E-6, 0.0026856738, 0.5380300509, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.639, 129.0, 0.817, 215.0, 0.928)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.817, 215.0, 0.928, 318.0, 0.982)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.08774E-6, 0.0011040371, 0.7409127896, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.817, 215.0, 0.928, 318.0, 0.982)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.928, 318.0, 0.982, 464.0, 1.019)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.010869E-7, 4.8887461E-4, 0.856984985, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.928, 318.0, 0.982, 464.0, 1.019)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.982, 464.0, 1.019, 774.0, 1.055)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-9.129045E-8, 2.2914661E-4, 0.9323304424, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.982, 464.0, 1.019, 774.0, 1.055)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.019, 774.0, 1.055, 1292.0, 1.076)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.778367E-8, 1.1860161E-4, 0.9858376464, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.019, 774.0, 1.055, 1292.0, 1.076)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.055, 1292.0, 1.076, 2154.0, 1.066)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.725632E-9, 1.15756E-5, 1.07227118, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.055, 1292.0, 1.076, 2154.0, 1.066)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.076, 2154.0, 1.066, 3594.0, 1.027)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.076, 2154.0, 1.066, 3594.0, 1.027)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.066, 3594.0, 1.027)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.650694797, -0.0761809201, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.066, 3594.0, 1.027)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr875K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.582, 0.788};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.582, 215.0, 0.788)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-7.742381E-6, 0.0050587279, 0.058265064, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.582, 215.0, 0.788)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.582, 215.0, 0.788, 318.0, 0.884)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.955154E-6, 0.0019741359, 0.4539377689, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.582, 215.0, 0.788, 318.0, 0.884)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.788, 318.0, 0.884, 464.0, 0.949)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.165089E-7, 8.4911544E-4, 0.6662127358, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.788, 318.0, 0.884, 464.0, 0.949)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.884, 464.0, 0.949, 774.0, 1.014)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.716303E-7, 4.2215568E-4, 0.7900710739, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.884, 464.0, 0.949, 774.0, 1.014)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.949, 774.0, 1.014, 1292.0, 1.049)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.896201E-8, 1.6872307E-4, 0.9127403055, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.949, 774.0, 1.014, 1292.0, 1.049)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.014, 1292.0, 1.049, 2154.0, 1.049)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-9.351771E-9, 3.2226204E-5, 1.022974319, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.014, 1292.0, 1.049, 2154.0, 1.049)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.049, 2154.0, 1.049, 3594.0, 1.018)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.049, 2154.0, 1.049, 3594.0, 1.018)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.049, 3594.0, 1.018)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.513757404, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.049, 3594.0, 1.018)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr875K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius4point5Gyr875K0coreMass(planet)
	coreMassRadii10 := gasRadius4point5Gyr875K10coreMass(planet)
	coreMassRadii25 := gasRadius4point5Gyr875K25coreMass(planet)
	coreMassRadii50 := gasRadius4point5Gyr875K50coreMass(planet)
	coreMassRadii100 := gasRadius4point5Gyr875K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if coreEarthMasses <= 10.0 {
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr260K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0021984339, 0.0970204354, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.014, 28.0, 0.993)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(4.6673633E-5, -0.0040094044, 1.068671195, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.014, 28.0, 0.993)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.014, 28.0, 0.993, 46.0, 0.983)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(2.9771048E-5, -0.0027586131, 1.0469900666, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.014, 28.0, 0.993, 46.0, 0.983)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.993, 46.0, 0.983, 77.0, 1.011)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.846094E-6, 0.0011302954, 0.9349127478, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.993, 46.0, 0.983, 77.0, 1.011)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.983, 77.0, 1.011, 129.0, 1.05)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.412538E-6, 0.0014529828, 0.9193532609, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.983, 77.0, 1.011, 129.0, 1.05)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.011, 129.0, 1.05, 215.0, 1.074)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.116977E-6, 6.6330976E-4, 0.9830206503, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.011, 129.0, 1.05, 215.0, 1.074)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.05, 215.0, 1.074, 318.0, 1.081)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.904145E-7, 1.6945211E-4, 1.046369708, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.05, 215.0, 1.074, 318.0, 1.081)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.074, 318.0, 1.081, 464.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(4.4576753E-9, 1.7062043E-5, 1.075123492, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.074, 318.0, 1.081, 464.0, 1.084)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.081, 464.0, 1.084, 774.0, 1.091)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.561369E-8, 4.1910394E-5, 1.067915142, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.081, 464.0, 1.084, 774.0, 1.091)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.084, 774.0, 1.091, 1292.0, 1.096)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.464816E-8, 6.0575605E-5, 1.058880602, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.084, 774.0, 1.091, 1292.0, 1.096)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.091, 1292.0, 1.096, 2154.0, 1.075)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-2.992203E-9, -1.405082E-5, 1.119148433, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.091, 1292.0, 1.096, 2154.0, 1.075)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.096, 2154.0, 1.075, 3594.0, 1.03)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.096, 2154.0, 1.075, 3594.0, 1.03)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.075, 3594.0, 1.03)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.749647843, -0.0879010616, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.075, 3594.0, 1.03)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr260K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.576, 0.738};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.576, 28.0, 0.738)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-3.028561E-4, 0.0283557994, 0.1814768373, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.576, 28.0, 0.738)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.576, 28.0, 0.738, 46.0, 0.845)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.4699E-5, 0.0107321703, 0.4882232463, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.576, 28.0, 0.738, 46.0, 0.845)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.738, 46.0, 0.845, 77.0, 0.931)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.813208E-5, 0.0050044396, 0.6531632635, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.738, 46.0, 0.845, 77.0, 0.931)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.845, 77.0, 0.931, 129.0, 0.997)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.489876E-6, 0.0024001452, 0.7787382943, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.845, 77.0, 0.931, 129.0, 0.997)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.931, 129.0, 0.997, 215.0, 1.041)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.833753E-6, 0.001142439, 0.8801408538, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.931, 129.0, 0.997, 215.0, 1.041)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.997, 215.0, 1.041, 318.0, 1.058)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-3.877727E-7, 3.7173137E-4, 0.9790025462, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.997, 215.0, 1.041, 318.0, 1.058)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.041, 318.0, 1.058, 464.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.116636E-8, 1.0850524E-4, 1.028669479, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.041, 318.0, 1.058, 464.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.058, 464.0, 1.068, 774.0, 1.082)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-3.589043E-8, 8.9593641E-5, 1.034155616, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.058, 464.0, 1.068, 774.0, 1.082)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.068, 774.0, 1.082, 1292.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.632296E-8, 6.9827253E-5, 1.043723161, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.068, 774.0, 1.082, 1292.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.082, 1292.0, 1.09, 2154.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.202383E-9, -6.40026E-6, 1.105284022, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.082, 1292.0, 1.09, 2154.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.072, 3594.0, 1.028)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.072, 3594.0, 1.028)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.072, 3594.0, 1.028)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.731655669, -0.0859477047, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.072, 3594.0, 1.028)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr260K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.400, 0.666};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.4, 46.0, 0.666)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.002048E-4, 0.0295929339, -0.2716415771, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.4, 46.0, 0.666)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.4, 46.0, 0.666, 77.0, 0.82)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.575593E-5, 0.0093657209, 0.3108363778, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.4, 46.0, 0.666, 77.0, 0.82)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.666, 77.0, 0.82, 129.0, 0.924)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-8.6788E-6, 0.0037878328, 0.5797934783, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.666, 77.0, 0.82, 129.0, 0.924)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.82, 129.0, 0.924, 215.0, 0.993)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.549932E-6, 0.0016795023, 0.7497776237, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.82, 129.0, 0.924, 215.0, 0.993)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.924, 215.0, 0.993, 318.0, 1.026)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-7.365544E-7, 7.1297185E-4, 0.8737582798, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.924, 215.0, 0.993, 318.0, 1.026)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.993, 318.0, 1.026, 464.0, 1.046)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.377034E-7, 2.4467036E-4, 0.9621199433, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.993, 318.0, 1.026, 464.0, 1.046)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.026, 464.0, 1.046, 774.0, 1.069)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-5.929592E-8, 1.476019E-4, 0.9902788934, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.026, 464.0, 1.046, 774.0, 1.069)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.046, 774.0, 1.069, 1292.0, 1.082)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.079559E-8, 8.8720219E-5, 1.018779451, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.046, 774.0, 1.069, 1292.0, 1.082)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.069, 1292.0, 1.082, 2154.0, 1.067)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.110893E-9, 2.1074347E-7, 1.090259148, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.069, 1292.0, 1.082, 2154.0, 1.067)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.082, 2154.0, 1.067, 3594.0, 1.025)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.082, 2154.0, 1.067, 3594.0, 1.025)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.067, 3594.0, 1.025)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.69667132, -0.0820409908, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.067, 3594.0, 1.025)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr260K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.633, 0.810};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.633, 129.0, 0.81)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.539693E-5, 0.0065756138, 0.2179661371, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.633, 129.0, 0.81)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.633, 129.0, 0.81, 215.0, 0.92)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-3.99364E-6, 0.0026528819, 0.5342364001, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.633, 129.0, 0.81, 215.0, 0.92)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.81, 215.0, 0.92, 318.0, 0.974)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.08774E-6, 0.0011040371, 0.7329127896, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.81, 215.0, 0.92, 318.0, 0.974)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.92, 318.0, 0.974, 464.0, 1.011)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.940128E-7, 4.8334264E-4, 0.8500287881, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.92, 318.0, 0.974, 464.0, 1.011)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.974, 464.0, 1.011, 774.0, 1.048)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-9.751787E-8, 2.4008197E-4, 0.9205971755, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.974, 464.0, 1.011, 774.0, 1.048)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.011, 774.0, 1.048, 1292.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.638476E-8, 1.1378095E-4, 0.9817307806, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.011, 774.0, 1.048, 1292.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.048, 1292.0, 1.068, 2154.0, 1.058)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.423962E-9, 1.0536045E-5, 1.065110718, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.048, 1292.0, 1.068, 2154.0, 1.058)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.068, 2154.0, 1.058, 3594.0, 1.02)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.068, 2154.0, 1.058, 3594.0, 1.02)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.058, 3594.0, 1.02)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.627702623, -0.0742275631, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.058, 3594.0, 1.02)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr260K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.578, 0.782};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.578, 215.0, 0.782)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-7.619334E-6, 0.004993144, 0.0606777624, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.578, 215.0, 0.782)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.578, 215.0, 0.782, 318.0, 0.878)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.982661E-6, 0.0019887973, 0.4460570955, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.578, 215.0, 0.782, 318.0, 0.878)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.782, 318.0, 0.878, 464.0, 0.942)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.014885E-7, 8.3052015E-4, 0.6646071121, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.782, 318.0, 0.878, 464.0, 0.942)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.878, 464.0, 0.942, 774.0, 1.007)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.739619E-7, 4.250421E-4, 0.7822337397, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.878, 464.0, 0.942, 774.0, 1.007)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.942, 774.0, 1.007, 1292.0, 1.041)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.756309E-8, 1.6390241E-4, 0.9086334397, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.942, 774.0, 1.007, 1292.0, 1.041)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.007, 1292.0, 1.041, 2154.0, 1.041)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-9.351771E-9, 3.2226204E-5, 1.014974319, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.007, 1292.0, 1.041, 2154.0, 1.041)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.041, 2154.0, 1.041, 3594.0, 1.01)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.041, 2154.0, 1.041, 3594.0, 1.01)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.041, 3594.0, 1.01)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.505757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.041, 3594.0, 1.01)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr260K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius4point5Gyr260K0coreMass(planet)
	coreMassRadii10 := gasRadius4point5Gyr260K10coreMass(planet)
	coreMassRadii25 := gasRadius4point5Gyr260K25coreMass(planet)
	coreMassRadii50 := gasRadius4point5Gyr260K50coreMass(planet)
	coreMassRadii100 := gasRadius4point5Gyr260K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if coreEarthMasses <= 10.0 {
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr78K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0015823147, 0.0738405271, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 0.798, 28.0, 0.827)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-1.619645E-5, 0.0033652038, 0.7454723093, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 0.798, 28.0, 0.827)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.798, 28.0, 0.827, 46.0, 0.866)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-1.327628E-5, 0.0031491113, 0.7492334869, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.798, 28.0, 0.827, 46.0, 0.866)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.827, 46.0, 0.866, 77.0, 0.913)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-8.07199E-6, 0.0025089838, 0.7676670752, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.827, 46.0, 0.866, 77.0, 0.913)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.866, 77.0, 0.913, 129.0, 0.957)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.013922E-6, 0.0014670219, 0.8179088629, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.866, 77.0, 0.913, 129.0, 0.957)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.913, 129.0, 0.957, 215.0, 0.994)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-9.921382E-7, 7.7152808E-4, 0.8739830482, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.913, 129.0, 0.957, 215.0, 0.994)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.957, 215.0, 0.994, 318.0, 1.019)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-4.796417E-7, 4.9836746E-4, 0.9090224331, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.957, 215.0, 0.994, 318.0, 1.019)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.994, 318.0, 1.019, 464.0, 1.037)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.359591E-7, 2.2960769E-4, 0.9597334837, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.994, 318.0, 1.019, 464.0, 1.037)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.019, 464.0, 1.037, 774.0, 1.056)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-6.003298E-8, 1.3561116E-4, 0.9870012845, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.019, 464.0, 1.037, 774.0, 1.056)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.037, 774.0, 1.056, 1292.0, 1.062)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.427802E-8, 4.1081391E-5, 1.032756619, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.037, 774.0, 1.056, 1292.0, 1.062)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.056, 1292.0, 1.062, 2154.0, 1.055)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.125792E-9, 1.2988829E-5, 1.055443997, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.056, 1292.0, 1.062, 2154.0, 1.055)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.062, 2154.0, 1.055, 3594.0, 1.023)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.062, 2154.0, 1.055, 3594.0, 1.023)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.055, 3594.0, 1.023)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.534749577, -0.0625074216, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.055, 3594.0, 1.023)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr78K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.508, 0.653};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.508, 28.0, 0.653)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-2.514803E-4, 0.0244984326, 0.1642044584, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.508, 28.0, 0.653)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.508, 28.0, 0.653, 46.0, 0.759)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.422354E-5, 0.0106414308, 0.405391193, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.508, 28.0, 0.653, 46.0, 0.759)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.653, 46.0, 0.759, 77.0, 0.844)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.751173E-5, 0.0048958788, 0.5708444049, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.653, 46.0, 0.759, 77.0, 0.844)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.759, 77.0, 0.844, 129.0, 0.911)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-4.702367E-6, 0.0022571492, 0.6980798495, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.759, 77.0, 0.844, 129.0, 0.911)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.844, 129.0, 0.911, 215.0, 0.966)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.688606E-6, 0.0012204153, 0.7816665126, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.844, 129.0, 0.911, 215.0, 0.966)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.911, 215.0, 0.966, 318.0, 0.999)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.99018E-7, 6.3966492E-4, 0.8561616467, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.911, 215.0, 0.966, 318.0, 0.999)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.966, 318.0, 0.999, 464.0, 1.024)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.057314E-7, 3.3211484E-4, 0.9141918645, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.966, 318.0, 0.999, 464.0, 1.024)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.999, 464.0, 1.024, 774.0, 1.048)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-7.251792E-8, 1.6719654E-4, 0.9620336238, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.999, 464.0, 1.024, 774.0, 1.048)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.024, 774.0, 1.048, 1292.0, 1.057)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.679346E-8, 5.2069816E-5, 1.017758524, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.024, 774.0, 1.048, 1292.0, 1.057)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.048, 1292.0, 1.057, 2154.0, 1.052)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.832022E-9, 1.7742682E-5, 1.045480902, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.048, 1292.0, 1.057, 2154.0, 1.052)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.057, 2154.0, 1.052, 3594.0, 1.021)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.057, 2154.0, 1.052, 3594.0, 1.021)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.052, 3594.0, 1.021)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.516757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.052, 3594.0, 1.021)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr78K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.378, 0.611};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.378, 46.0, 0.611)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-1.726648E-4, 0.025721637, -0.2068366615, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.378, 46.0, 0.611)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.378, 46.0, 0.611, 77.0, 0.75)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.108464E-5, 0.0083072812, 0.2946401537, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.378, 46.0, 0.611, 77.0, 0.75)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.611, 77.0, 0.75, 129.0, 0.849)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-7.307952E-6, 0.0034092842, 0.5308139632, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.611, 77.0, 0.75, 129.0, 0.849)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.75, 129.0, 0.849, 215.0, 0.926)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.374322E-6, 0.0017121157, 0.6676481738, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.75, 129.0, 0.849, 215.0, 0.926)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.849, 215.0, 0.926, 318.0, 0.972)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.608282E-9, 4.4959116E-4, 0.8295971443, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.849, 215.0, 0.926, 318.0, 0.972)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.926, 318.0, 0.972, 464.0, 1.037)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-9.904761E-7, 0.0012197578, 0.684277931, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.926, 318.0, 0.972, 464.0, 1.037)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.972, 464.0, 1.037, 774.0, 1.035)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(4.2764664E-8, -5.939427E-5, 1.055351879, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.972, 464.0, 1.037, 774.0, 1.035)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.037, 774.0, 1.035, 1292.0, 1.05)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.350566E-8, 7.7520217E-5, 0.989081027, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.037, 774.0, 1.035, 1292.0, 1.05)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.035, 1292.0, 1.05, 2154.0, 1.047)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.236581E-9, 2.1456981E-5, 1.034357345, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.035, 1292.0, 1.05, 2154.0, 1.047)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.05, 2154.0, 1.047, 3594.0, 1.018)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.05, 2154.0, 1.047, 3594.0, 1.018)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.047, 3594.0, 1.018)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.481773054, -0.0566473508, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.047, 3594.0, 1.018)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr78K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.594, 0.754};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.594, 129.0, 0.754)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.294366E-5, 0.0057433175, 0.2285075251, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.594, 129.0, 0.754)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.594, 129.0, 0.754, 215.0, 0.865)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-3.69558E-6, 0.0025619773, 0.4850030821, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.594, 129.0, 0.754, 215.0, 0.865)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.754, 215.0, 0.865, 318.0, 0.926)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.085603E-6, 0.0011708593, 0.6634472108, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.754, 215.0, 0.865, 318.0, 0.926)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.865, 318.0, 0.926, 464.0, 0.973)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.088464E-7, 6.4163566E-4, 0.7633040398, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.865, 318.0, 0.926, 464.0, 0.973)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.926, 464.0, 0.973, 774.0, 1.015)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.123343E-7, 2.7455378E-4, 0.8697921805, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.926, 464.0, 0.973, 774.0, 1.015)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.973, 774.0, 1.015, 1292.0, 1.037)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.909482E-8, 1.0258095E-4, 0.9530323566, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.973, 774.0, 1.015, 1292.0, 1.037)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.015, 1292.0, 1.037, 2154.0, 1.039)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-8.851321E-9, 3.2821838E-5, 1.009369377, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.015, 1292.0, 1.037, 2154.0, 1.039)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.037, 2154.0, 1.039, 3594.0, 1.013)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.037, 2154.0, 1.039, 3594.0, 1.013)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.039, 3594.0, 1.013)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.428796532, -0.05078728, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.039, 3594.0, 1.013)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr78K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.558, 0.746};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.558, 215.0, 0.746)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-6.634961E-6, 0.0044684732, 0.0919793497, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.558, 215.0, 0.746)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.558, 215.0, 0.746, 318.0, 0.842)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.845125E-6, 0.0019154904, 0.4194604624, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.558, 215.0, 0.746, 318.0, 0.842)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.746, 318.0, 0.842, 464.0, 0.911)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.765906E-7, 9.234966E-4, 0.6066352304, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.746, 318.0, 0.842, 464.0, 0.911)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.842, 464.0, 0.911, 774.0, 0.976)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.692987E-7, 4.1926925E-4, 0.7529084081, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.842, 464.0, 0.911, 774.0, 0.976)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.911, 774.0, 0.976, 1292.0, 1.012)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.11138E-8, 1.5443919E-4, 0.881094362, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.911, 774.0, 0.976, 1292.0, 1.012)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 0.976, 1292.0, 1.012, 2154.0, 1.023)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.127518E-8, 5.1615293E-5, 0.9641342947, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 0.976, 1292.0, 1.012, 2154.0, 1.023)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.012, 2154.0, 1.023, 3594.0, 1.004)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.012, 2154.0, 1.023, 3594.0, 1.004)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.023, 3594.0, 1.004)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.307851312, -0.0371137816, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.023, 3594.0, 1.004)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr78K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius4point5Gyr78K0coreMass(planet)
	coreMassRadii10 := gasRadius4point5Gyr78K10coreMass(planet)
	coreMassRadii25 := gasRadius4point5Gyr78K25coreMass(planet)
	coreMassRadii50 := gasRadius4point5Gyr78K50coreMass(planet)
	coreMassRadii100 := gasRadius4point5Gyr78K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if coreEarthMasses <= 10.0 {
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}
	return jupiterRadii
}

func gasRadius4point5Gyr(planet *types.Planet) float64 {
	temperatureRadii1960 := gasRadius4point5Gyr1960K(planet)
	temperatureRadii1300 := gasRadius4point5Gyr1300K(planet)
	temperatureRadii875 := gasRadius4point5Gyr875K(planet)
	temperatureRadii260 := gasRadius4point5Gyr260K(planet)
	temperatureRadii78 := gasRadius4point5Gyr78K(planet)
	temperatureRadii0 := 0.0

	temperature := planet.EstimatedTemperature

	var jupiterRadii float64
	if temperature <= 78.0 {
		jupiterRadii1 := PlanetRadiusHelper(temperature, 0.0, temperatureRadii0, 78.0, temperatureRadii78, 260.0, temperatureRadii260)
		jupiterRadii2 := PlanetRadiusHelper2(temperature, 78.0, temperatureRadii78, 260.0, temperatureRadii260)
		jupiterRadii = RangeAdjust(temperature, jupiterRadii1, jupiterRadii2, 78.0, 260.0)
	} else if temperature <= 260.0 {
		//jupiterRadii1 := PlanetRadiusHelper(temperature, 0.0, temperatureRadii0, 78.0, temperatureRadii78, 260.0, temperatureRadii260, false);
		jupiterRadii1 := PlanetRadiusHelper2(temperature, 78.0, temperatureRadii78, 260.0, temperatureRadii260)
		jupiterRadii2 := PlanetRadiusHelper(temperature, 78.0, temperatureRadii78, 260.0, temperatureRadii260, 875.0, temperatureRadii875)
		jupiterRadii = RangeAdjust(temperature, jupiterRadii1, jupiterRadii2, 78.0, 260.0)
	} else if temperature <= 875.0 {
		jupiterRadii1 := PlanetRadiusHelper(temperature, 78.0, temperatureRadii78, 260.0, temperatureRadii260, 875.0, temperatureRadii875)
		jupiterRadii2 := PlanetRadiusHelper(temperature, 260.0, temperatureRadii260, 875.0, temperatureRadii875, 1300.0, temperatureRadii1300)
		jupiterRadii = RangeAdjust(temperature, jupiterRadii1, jupiterRadii2, 260.0, 875.0)
	} else if temperature <= 1300.0 {
		jupiterRadii1 := PlanetRadiusHelper(temperature, 260.0, temperatureRadii260, 875.0, temperatureRadii875, 1300.0, temperatureRadii1300)
		jupiterRadii2 := PlanetRadiusHelper(temperature, 875.0, temperatureRadii875, 1300.0, temperatureRadii1300, 1960.0, temperatureRadii1960)
		jupiterRadii = RangeAdjust(temperature, jupiterRadii1, jupiterRadii2, 875.0, 1300.0)
	} else if temperature <= 1960.0 {
		jupiterRadii1 := PlanetRadiusHelper(temperature, 875.0, temperatureRadii875, 1300.0, temperatureRadii1300, 1960.0, temperatureRadii1960)
		jupiterRadii2 := PlanetRadiusHelper3(temperature, 1300.0, temperatureRadii1300, 1960.0, temperatureRadii1960)
		jupiterRadii = RangeAdjust(temperature, jupiterRadii1, jupiterRadii2, 1300.0, 1960.0)
	} else {
		jupiterRadii = PlanetRadiusHelper3(temperature, 1300.0, temperatureRadii1300, 1960.0, temperatureRadii1960)
	}
	return jupiterRadii
}
