package radius

import (
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

func gasRadius1Gyr1960K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.909, 1.150};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.909, 28.0, 1.15)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-6.194706E-4, 0.0497852665, 0.2416774643, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.909, 28.0, 1.15)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.909, 28.0, 1.15, 46.0, 1.221)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-8.708214E-5, 0.0103885232, 0.973937532, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.909, 28.0, 1.15, 46.0, 1.221)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.15, 46.0, 1.221, 77.0, 1.211)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(7.825346E-6, -0.0012850982, 1.263556085, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.15, 46.0, 1.221, 77.0, 1.211)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.221, 77.0, 1.211, 129.0, 1.228)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.863447E-6, 7.1079308E-4, 1.167317308, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.221, 77.0, 1.211, 129.0, 1.228)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.211, 129.0, 1.228, 215.0, 1.234)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-6.259848E-7, 2.8510622E-4, 1.201638311, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.211, 129.0, 1.228, 215.0, 1.234)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.228, 215.0, 1.234, 318.0, 1.229)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(1.9495458E-7, -1.524545E-4, 1.257765938, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.228, 215.0, 1.234, 318.0, 1.229)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.234, 318.0, 1.229, 464.0, 1.229)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.537068E-8, 2.7659875E-5, 1.223780985, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.234, 318.0, 1.229, 464.0, 1.229)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.229, 464.0, 1.229, 774.0, 1.224)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(4.978943E-8, -7.776821E-5, 1.254365008, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.229, 464.0, 1.229, 774.0, 1.224)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.229, 774.0, 1.224, 1292.0, 1.237)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.986718E-8, 6.6142122E-5, 1.184707949, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.229, 774.0, 1.224, 1292.0, 1.237)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.224, 1292.0, 1.237, 2154.0, 1.235)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.045556E-8, 3.370968E-5, 1.210900186, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.224, 1292.0, 1.237, 2154.0, 1.235)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.237, 2154.0, 1.235, 3594.0, 1.197)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.237, 2154.0, 1.235, 3594.0, 1.197)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.235, 3594.0, 1.197)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.804702623, -0.0742275631, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.235, 3594.0, 1.197)
	}

	return jupiterRadii

}

func gasRadius1Gyr1960K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.461, 0.838};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.461, 46.0, 0.838)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-3.063053E-4, 0.043611038, -0.5199656938, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.461, 46.0, 0.838)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.461, 46.0, 0.838, 77.0, 1.022)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-4.857395E-5, 0.0119100795, 0.3929188167, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.461, 46.0, 0.838, 77.0, 1.022)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.838, 77.0, 1.022, 129.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.751497E-6, 0.003912654, 0.7785422241, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.838, 77.0, 1.022, 129.0, 1.121)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.022, 129.0, 1.121, 215.0, 1.169)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.92574E-6, 0.0012205939, 0.9955896132, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.022, 129.0, 1.121, 215.0, 1.169)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.121, 215.0, 1.169, 318.0, 1.189)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-4.772381E-7, 4.4854267E-4, 1.094623657, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.121, 215.0, 1.169, 318.0, 1.189)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.169, 318.0, 1.189, 464.0, 1.2)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.227799E-7, 1.7135633E-4, 1.146924678, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.169, 318.0, 1.189, 464.0, 1.2)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.189, 464.0, 1.2, 774.0, 1.206)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(2.791812E-8, -1.1520779E-5, 1.201045757, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.189, 464.0, 1.2, 774.0, 1.206)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.2, 774.0, 1.206, 1292.0, 1.228)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.993547E-8, 1.0431773E-4, 1.143191703, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.2, 774.0, 1.206, 1292.0, 1.228)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.206, 1292.0, 1.228, 2154.0, 1.229)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.166574E-8, 4.1360238E-5, 1.194035774, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.206, 1292.0, 1.228, 2154.0, 1.229)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.228, 2154.0, 1.229, 3594.0, 1.192)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.228, 2154.0, 1.229, 3594.0, 1.192)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.229, 3594.0, 1.192)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.783710449, -0.0722742062, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.229, 3594.0, 1.192)
	}
	return jupiterRadii

}

func gasRadius1Gyr1960K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.746, 0.958};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.746, 129.0, 0.958)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.993726E-5, 0.008183998, 0.2340401338, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.746, 129.0, 0.958)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.746, 129.0, 0.958, 215.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.445209E-6, 0.002854733, 0.6637121282, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.746, 129.0, 0.958, 215.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.958, 215.0, 1.072, 318.0, 1.122)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.014298E-6, 0.0010260577, 0.8982835195, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.958, 215.0, 1.072, 318.0, 1.122)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.072, 318.0, 1.122, 464.0, 1.156)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.409153E-7, 4.9947244E-4, 0.9976424774, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.072, 318.0, 1.122, 464.0, 1.156)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.122, 464.0, 1.156, 774.0, 1.18)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-2.122439E-8, 1.0369515E-4, 1.112454977, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.122, 464.0, 1.156, 774.0, 1.18)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.156, 774.0, 1.18, 1292.0, 1.211)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.748182E-8, 1.37283E-4, 1.096197418, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.156, 774.0, 1.18, 1292.0, 1.211)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.18, 1292.0, 1.211, 2154.0, 1.218)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.318109E-8, 5.354269E-5, 1.163825566, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.18, 1292.0, 1.211, 2154.0, 1.218)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.211, 2154.0, 1.218, 3594.0, 1.186)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.211, 2154.0, 1.218, 3594.0, 1.186)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.218, 3594.0, 1.186)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.697749577, -0.0625074216, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.218, 3594.0, 1.186)
	}
	return jupiterRadii

}

func gasRadius1Gyr1960K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.640, 0.888};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.64, 215.0, 0.888)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-9.658564E-6, 0.0062062668, 1.197411E-4, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.64, 215.0, 0.888)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.64, 215.0, 0.888, 318.0, 0.997)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.296993E-6, 0.0022825493, 0.5034303716, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.64, 215.0, 0.888, 318.0, 0.997)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.888, 318.0, 0.997, 464.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-6.278539E-7, 9.7728311E-4, 0.7497150685, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.888, 318.0, 0.997, 464.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.997, 464.0, 1.068, 774.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.273012E-7, 3.575989E-4, 0.9294815511, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.997, 464.0, 1.068, 774.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.068, 774.0, 1.13, 1292.0, 1.179)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.257452E-8, 2.0321354E-4, 1.004208648, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.068, 774.0, 1.13, 1292.0, 1.179)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.13, 1292.0, 1.179, 2154.0, 1.198)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.71168E-8, 8.1026257E-5, 1.102886534, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.13, 1292.0, 1.179, 2154.0, 1.198)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.179, 2154.0, 1.198, 3594.0, 1.173)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.179, 2154.0, 1.198, 3594.0, 1.173)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.198, 3594.0, 1.173)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.572804357, -0.0488339231, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.198, 3594.0, 1.173)
	}

	return jupiterRadii
}

func gasRadius1Gyr1960K(planet *types.Planet) float64 {
	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	coreMassRadii0 := gasRadius1Gyr1960K0coreMass(planet)
	coreMassRadii10 := gasRadius1Gyr1960K10coreMass(planet)
	coreMassRadii25 := gasRadius1Gyr1960K25coreMass(planet)
	coreMassRadii50 := gasRadius1Gyr1960K50coreMass(planet)
	coreMassRadii100 := gasRadius1Gyr1960K100coreMass(planet)

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

func gasRadius1Gyr1300K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0038412911, 0.1529490069, 0.0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.49, 28.0, 1.271)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(5.17938E-4, -0.0432163009, 2.074993034, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.49, 28.0, 1.271)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.49, 28.0, 1.271, 46.0, 1.183)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(7.4098457E-5, -0.0103721747, 1.503327701, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.49, 28.0, 1.271, 46.0, 1.183)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.271, 46.0, 1.183, 77.0, 1.144)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(1.9559628E-5, -0.0036638988, 1.31015117, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.271, 46.0, 1.183, 77.0, 1.144)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.183, 77.0, 1.144, 129.0, 1.163)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-2.310674E-6, 8.4138342E-4, 1.092913462, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.183, 77.0, 1.144, 129.0, 1.163)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.144, 129.0, 1.163, 215.0, 1.167)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-6.056762E-7, 2.5486422E-4, 1.140201572, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.144, 129.0, 1.163, 215.0, 1.167)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.163, 215.0, 1.167, 318.0, 1.16)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(1.9041454E-7, -1.694521E-4, 1.194630292, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.163, 215.0, 1.167, 318.0, 1.16)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.167, 318.0, 1.16, 464.0, 1.157)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(3.7987146E-8, -5.025389E-5, 1.172139326, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.167, 318.0, 1.16, 464.0, 1.157)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.16, 464.0, 1.157, 774.0, 1.156)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(2.2548094E-8, -3.114035E-5, 1.166594607, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.16, 464.0, 1.157, 774.0, 1.156)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.157, 774.0, 1.156, 1292.0, 1.164)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.380102E-8, -6.4616923E-5, 1.120245122, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.157, 774.0, 1.156, 1292.0, 1.164)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.156, 1292.0, 1.164, 2154.0, 1.149)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.110893E-9, 2.1074347E-7, 1.172259148, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.156, 1292.0, 1.164, 2154.0, 1.149)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.164, 2154.0, 1.149, 3594.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.164, 2154.0, 1.149, 3594.0, 1.107)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.149, 3594.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.77867132, -0.0820409908, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.149, 3594.0, 1.107)
	}

	return jupiterRadii
}

func gasRadius1Gyr1300K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.698, 0.888};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.698, 28.0, 0.888)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-4.289446E-4, 0.0365752351, 0.2001859979, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.698, 28.0, 0.888)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.698, 28.0, 0.888, 46.0, 0.975)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-5.387316E-5, 0.0088199473, 0.6832780338, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.698, 28.0, 0.888, 46.0, 0.975)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.888, 46.0, 0.975, 77.0, 1.043)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.345332E-5, 0.0038483064, 0.8264451254, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.888, 46.0, 0.975, 77.0, 1.043)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.975, 77.0, 1.043, 129.0, 1.099)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.444505E-6, 0.0021984911, 0.9059966555, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.975, 77.0, 1.043, 129.0, 1.099)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.043, 129.0, 1.099, 215.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.36307E-6, 7.99447748E-4, 1.019195254, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.043, 129.0, 1.099, 215.0, 1.127)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.099, 215.0, 1.127, 318.0, 1.134)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.078927E-7, 1.2546796E-4, 1.105011728, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.099, 215.0, 1.127, 318.0, 1.134)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.127, 318.0, 1.134, 464.0, 1.14)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.060361E-8, 7.2847912E-5, 1.11490363, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.127, 318.0, 1.134, 464.0, 1.14)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.134, 464.0, 1.14, 774.0, 1.147)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.624546E-9, 2.4591833E-5, 1.128939148, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.134, 464.0, 1.14, 774.0, 1.147)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.14, 774.0, 1.147, 1292.0, 1.158)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.631647E-8, 7.5605347E-5, 1.104247027, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.14, 774.0, 1.147, 1292.0, 1.158)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.147, 1292.0, 1.158, 2154.0, 1.145)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.515452E-9, 3.9250423E-6, 1.162135591, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.147, 1292.0, 1.158, 2154.0, 1.145)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.158, 2154.0, 1.145, 3594.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.158, 2154.0, 1.145, 3594.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.145, 3594.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.744686973, -0.078134277, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.145, 3594.0, 1.105)
	}
	return jupiterRadii
}

func gasRadius1Gyr1300K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.426, 0.739};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.436, 46.0, 0.739)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.436179E-4, 0.0354166118, -0.3746687148, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.436, 46.0, 0.739)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.436, 46.0, 0.739, 77.0, 0.908)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-4.15857E-5, 0.0105666537, 0.3409292654, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.436, 46.0, 0.739, 77.0, 0.908)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.739, 77.0, 0.908, 129.0, 1.012)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.437142E-6, 0.0039440512, 0.660260896, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.739, 77.0, 0.908, 129.0, 1.012)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.908, 129.0, 1.012, 215.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.304436E-6, 0.0014904006, 0.858086454, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.908, 129.0, 1.012, 215.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.012, 215.0, 1.072, 318.0, 1.099)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-6.126381E-7, 5.8867202E-4, 0.9737547114, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.012, 215.0, 1.072, 318.0, 1.099)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.072, 318.0, 1.099, 464.0, 1.115)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.200665E-7, 2.0348106E-4, 1.04643463, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.072, 318.0, 1.099, 464.0, 1.115)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.099, 464.0, 1.115, 774.0, 1.132)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-2.659442E-8, 8.7762597E-5, 1.080003826, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.099, 464.0, 1.115, 774.0, 1.132)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.115, 774.0, 1.132, 1292.0, 1.149)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.134737E-8, 9.7582195E-5, 1.075250837, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.115, 774.0, 1.132, 1292.0, 1.149)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.132, 1292.0, 1.149, 2154.0, 1.14)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.229582E-9, 1.4472305E-5, 1.142369863, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.132, 1292.0, 1.149, 2154.0, 1.14)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.149, 2154.0, 1.14, 3594.0, 1.101)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.149, 2154.0, 1.14, 3594.0, 1.101)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.14, 3594.0, 1.101)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.724684797, -0.0761809201, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.14, 3594.0, 1.101)
	}
	return jupiterRadii
}

func gasRadius1Gyr1300K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.684, 0.877};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.684, 129.0, 0.877)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.754232E-5, 0.0073252573, 0.2239636288, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.684, 129.0, 0.877)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.684, 129.0, 0.877, 215.0, 0.988)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.106532E-6, 0.0027033447, 0.5966053321, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.684, 129.0, 0.877, 215.0, 0.988)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.877, 215.0, 0.988, 318.0, 1.041)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.076256E-6, 0.0010882076, 0.8037853037, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.877, 215.0, 0.988, 318.0, 1.041)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.988, 318.0, 1.041, 464.0, 1.077)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.14363E-7, 4.9240722E-4, 0.9162041491, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.988, 318.0, 1.041, 464.0, 1.077)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.041, 464.0, 1.077, 774.0, 1.109)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-6.638075E-8, 1.8540517E-4, 1.00526351, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.041, 464.0, 1.077, 774.0, 1.109)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.077, 774.0, 1.109, 1292.0, 1.134)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.833545E-8, 1.2746359E-4, 1.033309032, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.077, 774.0, 1.109, 1292.0, 1.134)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.109, 1292.0, 1.134, 2154.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-8.542652E-9, 2.4797607E-5, 1.116221433, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.109, 1292.0, 1.134, 2154.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.134, 2154.0, 1.13, 3594.0, 1.095)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.134, 2154.0, 1.13, 3594.0, 1.095)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.13, 3594.0, 1.095)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.6547261, -0.0683674924, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.13, 3594.0, 1.095)
	}
	return jupiterRadii
}

func gasRadius1Gyr1300K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.607, 0.831};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.607, 215.0, 0.831)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.592956E-6, 0.0055606279, 0.032674372, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.607, 215.0, 0.831)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.607, 215.0, 0.831, 318.0, 0.932)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.095094E-6, 0.0020972676, 0.4769331781, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.607, 215.0, 0.831, 318.0, 0.932)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.831, 318.0, 0.932, 464.0, 0.999)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.394756E-7, 8.8077405E-4, 0.7064677861, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.831, 318.0, 0.932, 464.0, 0.999)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.932, 464.0, 0.999, 774.0, 1.065)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.638685E-7, 4.1577247E-4, 0.8413618123, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.932, 464.0, 0.999, 774.0, 1.065)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.999, 774.0, 1.065, 1292.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.09127E-8, 1.8240572E-4, 0.954318557, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.999, 774.0, 1.065, 1292.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.065, 1292.0, 1.105, 2154.0, 1.111)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.116879E-8, 4.5448211E-5, 1.064924573, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.065, 1292.0, 1.105, 2154.0, 1.111)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.105, 2154.0, 1.111, 3594.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.105, 2154.0, 1.111, 3594.0, 1.084)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.111, 3594.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.515788706, -0.52740637, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.111, 3594.0, 1.084)
	}
	return jupiterRadii
}

func gasRadius1Gyr1300K(planet *types.Planet) float64 {
	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	coreMassRadii0 := gasRadius1Gyr1300K0coreMass(planet)
	coreMassRadii10 := gasRadius1Gyr1300K10coreMass(planet)
	coreMassRadii25 := gasRadius1Gyr1300K25coreMass(planet)
	coreMassRadii50 := gasRadius1Gyr1300K50coreMass(planet)
	coreMassRadii100 := gasRadius1Gyr1300K100coreMass(planet)

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

func gasRadius1Gyr875K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0030548128, 0.1282847594, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0, 17.0, 1.298, 28.0, 1.197)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(1.825148E-4, -0.0173949843, 1.540967955, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0, 17.0, 1.298, 28.0, 1.197)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.298, 28.0, 1.197, 46.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(6.4881867E-5, -0.008690147, 1.389456733, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.298, 28.0, 1.197, 46.0, 1.127)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.197, 46.0, 1.127, 77.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(1.5037819E-5, -0.0025593291, 1.212909115, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.197, 46.0, 1.127, 77.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.127, 77.0, 1.105, 129.0, 1.133)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.059293E-6, 0.0011686759, 1.033150502, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.127, 77.0, 1.105, 129.0, 1.133)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.105, 129.0, 1.133, 215.0, 1.143)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.207091E-7, 3.98603E-4, 1.095237633, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.105, 129.0, 1.133, 215.0, 1.143)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.133, 215.0, 1.143, 318.0, 1.139)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(1.2845637E-7, -1.073022E-4, 1.160132077, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.133, 215.0, 1.143, 318.0, 1.139)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.143, 318.0, 1.139, 464.0, 1.138)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(2.2094565E-8, -2.412726E-5, 1.144438179, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.143, 318.0, 1.139, 464.0, 1.138)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.139, 464.0, 1.138, 774.0, 1.139)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(1.4756291E-8, -1.504248E-5, 1.141802741, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.139, 464.0, 1.138, 774.0, 1.139)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.138, 774.0, 1.139, 1292.0, 1.147)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.548231E-8, 6.8090476E-5, 1.101563814, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.138, 774.0, 1.139, 1292.0, 1.147)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.139, 1292.0, 1.147, 2154.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.404663E-9, -4.54311E-6, 1.160222243, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.139, 1292.0, 1.147, 2154.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.147, 2154.0, 1.13, 3594.0, 1.087)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.147, 2154.0, 1.13, 3594.0, 1.087)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.13, 3594.0, 1.087)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.774663495, -0.0839943477, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.13, 3594.0, 1.087)
	}
	return jupiterRadii
}

func gasRadius1Gyr875K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.665, 0.847};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.665, 28.0, 0.847)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-4.038662E-4, 0.0347194357, 0.1914869383, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.665, 28.0, 0.847)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.665, 28.0, 0.847, 46.0, 0.934)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-4.728988E-5, 0.0083327847, 0.6507572965, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.665, 28.0, 0.847, 46.0, 0.934)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.847, 46.0, 0.934, 77.0, 1.012)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.641305E-5, 0.004539338, 0.760123053, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.847, 46.0, 0.934, 77.0, 1.012)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.934, 77.0, 1.012, 129.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.580618E-6, 0.0023034543, 0.8677215719, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.934, 77.0, 1.012, 129.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.012, 129.0, 1.072, 215.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.567949E-6, 9.2309526E-4, 0.979012945, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.012, 129.0, 1.072, 215.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.072, 215.0, 1.105, 318.0, 1.114)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.308599E-7, 1.5712698E-4, 1.077266699, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.072, 215.0, 1.105, 318.0, 1.114)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.105, 318.0, 1.114, 464.0, 1.122)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-6.357033E-8, 1.0450652E-4, 1.087195414, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.105, 318.0, 1.114, 464.0, 1.122)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.114, 464.0, 1.122, 774.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-5.520447E-9, 3.2640765E-5, 1.108042315, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.114, 464.0, 1.122, 774.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.122, 774.0, 1.13, 1292.0, 1.141)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.799776E-8, 7.90789E-5, 1.085565719, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.122, 774.0, 1.13, 1292.0, 1.141)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.13, 1292.0, 1.141, 2154.0, 1.126)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.809222E-9, -8.288115E-7, 1.150098686, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.13, 1292.0, 1.141, 2154.0, 1.126)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.141, 2154.0, 1.126, 3594.0, 1.085)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.141, 2154.0, 1.126, 3594.0, 1.085)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.126, 3594.0, 1.085)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.740679146, -0.0800876339, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.126, 3594.0, 1.085)
	}
	return jupiterRadii
}

func gasRadius1Gyr875K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.420, 0.719};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.42, 46.0, 0.719)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.310365E-4, 0.0377078122, -0.3426861239, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.42, 46.0, 0.719)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.42, 46.0, 0.719, 77.0, 0.883)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.917905E-5, 0.0101093456, 0.3368729708, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.42, 46.0, 0.719, 77.0, 0.883)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.719, 77.0, 0.883, 129.0, 0.989)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.547328E-6, 0.0040052112, 0.6312048495, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.719, 77.0, 0.883, 129.0, 0.989)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.883, 129.0, 0.989, 215.0, 1.051)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.324745E-6, 0.0015206426, 0.8315231931, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.883, 129.0, 0.989, 215.0, 1.051)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.989, 215.0, 1.051, 318.0, 1.08)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-6.631126E-7, 6.3499242E-4, 0.9451290097, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.989, 215.0, 1.051, 318.0, 1.08)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.051, 318.0, 1.08, 464.0, 1.097)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.209387E-7, 2.110124E-4, 1.02512786, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.051, 318.0, 1.08, 464.0, 1.097)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.08, 464.0, 1.097, 774.0, 1.116)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-3.671774E-8, 1.0674689E-4, 1.055374627, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.08, 464.0, 1.097, 774.0, 1.116)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.097, 774.0, 1.116, 1292.0, 1.132)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.162975E-8, 9.623509E-5, 1.060462663, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.097, 774.0, 1.116, 1292.0, 1.132)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.116, 1292.0, 1.132, 2154.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.523352E-9, 9.718451E-6, 1.130332958, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.116, 1292.0, 1.132, 2154.0, 1.121)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.132, 2154.0, 1.121, 3594.0, 1.081)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.132, 2154.0, 1.121, 3594.0, 1.081)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.121, 3594.0, 1.081)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.720686972, -0.078134277, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.121, 3594.0, 1.081)
	}
	return jupiterRadii
}

func gasRadius1Gyr875K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.670, 0.859};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.67, 129.0, 0.859)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.698491E-5, 0.007133507, 0.221423495, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.67, 129.0, 0.859)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.67, 129.0, 0.859, 215.0, 0.97)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.106532E-6, 0.0027033447, 0.5786053321, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.67, 129.0, 0.859, 215.0, 0.97)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.859, 215.0, 0.97, 318.0, 1.023)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.076256E-6, 0.0010882076, 0.7857853037, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.859, 215.0, 0.97, 318.0, 1.023)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.97, 318.0, 1.023, 464.0, 1.059)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.931406E-7, 4.758113E-4, 0.9013355583, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.97, 318.0, 1.023, 464.0, 1.059)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.023, 464.0, 1.059, 774.0, 1.094)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-8.27315E-8, 2.1532482E-4, 0.9769010435, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.023, 464.0, 1.059, 774.0, 1.094)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.059, 774.0, 1.094, 1292.0, 1.117)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292 {
		//jupiterRadii = quad_trend(-3.721891E-8, 1.2129582E-4, 1.022413993, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.059, 774.0, 1.094, 1292.0, 1.117)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.094, 1292.0, 1.117, 2154.0, 1.111)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.534752E-9, 1.9004198E-5, 1.105024066, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.094, 1292.0, 1.117, 2154.0, 1.111)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.117, 2154.0, 1.111, 3594.0, 1.076)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.117, 2154.0, 1.111, 3594.0, 1.076)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.111, 3594.0, 1.076)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.6357261, -0.0683674924, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.111, 3594.0, 1.076)
	}
	return jupiterRadii
}

func gasRadius1Gyr875K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.600, 0.818};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.6, 215.0, 0.818)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.275185E-6, 0.0053815473, 0.0434877485, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.6, 215.0, 0.818)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.6, 215.0, 0.818, 318.0, 0.918)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.08361E-6, 0.0020814381, 0.4668056922, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.6, 215.0, 0.818, 318.0, 0.918)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.818, 318.0, 0.918, 464.0, 0.984)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.244552E-7, 8.6217876E-4, 0.6968621625, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.818, 318.0, 0.918, 464.0, 0.984)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.918, 464.0, 0.984, 774.0, 1.05)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.685316E-7, 4.2154533E-4, 0.824871439, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.918, 464.0, 0.984, 774.0, 1.05)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.984, 774.0, 1.05, 1292.0, 1.088)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.895551E-8, 1.7450117E-4, 0.9442641716, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.984, 774.0, 1.05, 1292.0, 1.088)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.05, 1292.0, 1.088, 2154.0, 1.093)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.096651E-8, 4.3591061E-5, 1.049986351, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.05, 1292.0, 1.088, 2154.0, 1.093)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.088, 2154.0, 1.093, 3594.0, 1.065)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.088, 2154.0, 1.093, 3594.0, 1.065)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.093, 3594.0, 1.065)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.51278088, -0.0546939939, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.093, 3594.0, 1.065)
	}
	return jupiterRadii
}

func gasRadius1Gyr875K(planet *types.Planet) float64 {
	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	coreMassRadii0 := gasRadius1Gyr875K0coreMass(planet)
	coreMassRadii10 := gasRadius1Gyr875K10coreMass(planet)
	coreMassRadii25 := gasRadius1Gyr875K25coreMass(planet)
	coreMassRadii50 := gasRadius1Gyr875K50coreMass(planet)
	coreMassRadii100 := gasRadius1Gyr875K100coreMass(planet)
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

func gasRadius1Gyr260K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0028449198, 0.120657754, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.229, 28.0, 1.148)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(1.5238593E-4, -0.0142210031, 1.42671752, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.229, 28.0, 1.148)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.229, 28.0, 1.148, 46.0, 1.095)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(5.4165752E-5, -0.0069527101, 1.300209933, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.229, 28.0, 1.148, 46.0, 1.095)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.148, 46.0, 1.095, 77.0, 1.086)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(1.0912135E-5, 0.0016325152, 1.14700562, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.148, 46.0, 1.095, 77.0, 1.086)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.095, 77.0, 1.086, 129.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.448186E-6, 0.001325711, 1.0043645448, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.095, 77.0, 1.086, 129.0, 1.118)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.086, 129.0, 1.118, 215.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.410178E-7, 4.28845E-4, 1.076674372, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.086, 129.0, 1.118, 215.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.118, 215.0, 1.13, 318.0, 1.128)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(5.0474541E-8, -4.632041E-5, 1.137625702, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.118, 215.0, 1.13, 318.0, 1.128)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.13, 318.0, 1.128, 464.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(3.6242839E-8, -3.519121E-5, 1.135525786, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.13, 318.0, 1.128, 464.0, 1.127)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.128, 464.0, 1.127, 774.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(4.632964E-9, 3.94181E-6, 1.124173542, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.128, 464.0, 1.127, 774.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.127, 774.0, 1.13, 1292.0, 1.137)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.324275E-8, 6.153304E-5, 1.096297602, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.127, 774.0, 1.13, 1292.0, 1.137)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.13, 1292.0, 1.137, 2154.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.606943E-9, -2.685961E-6, 1.148160465, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.13, 1292.0, 1.137, 2154.0, 1.121)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.137, 2154.0, 1.121, 3594.0, 1.079)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.137, 2154.0, 1.121, 3594.0, 1.079)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.121, 3594.0, 1.079)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.75067132, -0.0820409908, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.121, 3594.0, 1.079)
	}
	return jupiterRadii
}

func gasRadius1Gyr260K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.646, 0.823};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.646, 28.0, 0.823)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-3.786137E-4, 0.0331285266, 0.1922344131, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.646, 28.0, 0.823)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.646, 28.0, 0.823, 46.0, 0.915)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-5.098383E-5, 0.0088839149, 0.6142217102, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.646, 28.0, 0.823, 46.0, 0.915)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.823, 46.0, 0.915, 77.0, 0.996)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.711561E-5, 0.0047181231, 0.7341829651, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.823, 46.0, 0.915, 77.0, 0.996)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.915, 77.0, 0.996, 129.0, 1.058)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.775064E-6, 0.0023819709, 0.8468285953, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.915, 77.0, 0.996, 129.0, 1.058)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.996, 129.0, 1.058, 215.0, 1.092)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.526734E-6, 9.2054532E-4, 0.9646560333, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.996, 129.0, 1.058, 215.0, 1.092)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.058, 215.0, 1.092, 318.0, 1.103)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.088418E-7, 2.1810877E-4, 1.054760324, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.058, 215.0, 1.092, 318.0, 1.103)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.092, 318.0, 1.103, 464.0, 1.111)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.942205E-8, 9.3442566E-5, 1.07828302, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.092, 318.0, 1.103, 464.0, 1.111)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.103, 464.0, 1.111, 774.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.564377E-8, 5.1625057E-5, 1.090414015, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.103, 464.0, 1.111, 774.0, 1.121)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.111, 774.0, 1.121, 1292.0, 1.131)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.57582E-8, 7.2521464E-5, 1.080299507, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.111, 774.0, 1.121, 1292.0, 1.131)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.121, 1292.0, 1.131, 2154.0, 1.117)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.011502E-9, 1.0283379E-6, 1.138036908, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.121, 1292.0, 1.131, 2154.0, 1.117)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.131, 2154.0, 1.117, 3594.0, 1.077)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.131, 2154.0, 1.117, 3594.0, 1.077)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.117, 3594.0, 1.077)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.716686972, -0.078134277, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.117, 3594.0, 1.077)
	}
	return jupiterRadii
}

func gasRadius1Gyr260K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.416, 0.709};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.416, 46.0, 0.709)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.255504E-4, 0.03296851, -0.3302867384, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.416, 46.0, 0.709)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.416, 46.0, 0.709, 77.0, 0.871)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.840175E-5, 0.0099492212, 0.3325939191, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.416, 46.0, 0.709, 77.0, 0.871)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.709, 77.0, 0.871, 129.0, 0.977)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.463068E-6, 0.0039878536, 0.620041806, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.709, 77.0, 0.871, 129.0, 0.977)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.871, 129.0, 0.977, 215.0, 1.04)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.386268E-6, 0.0015534345, 0.8163168439, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.871, 129.0, 0.977, 215.0, 1.04)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.977, 215.0, 1.04, 318.0, 1.069)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-6.356053E-7, 6.2033104E-4, 0.9360096831, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.977, 215.0, 1.04, 318.0, 1.069)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.04, 318.0, 1.069, 464.0, 1.087)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.28885E-7, 2.2407571E-4, 1.010777287, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.04, 318.0, 1.069, 464.0, 1.087)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.069, 464.0, 1.087, 774.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-4.061365E-8, 1.1479582E-4, 1.042478694, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.069, 464.0, 1.087, 774.0, 1.107)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.087, 774.0, 1.107, 1292.0, 1.123)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.162975E-8, 9.623509E-5, 1.051462663, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.087, 774.0, 1.107, 1292.0, 1.123)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.107, 1292.0, 1.123, 2154.0, 1.112)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.221682E-9, 8.678896E-6, 1.122172496, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.107, 1292.0, 1.123, 2154.0, 1.112)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.123, 2154.0, 1.112, 3594.0, 1.073)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.123, 2154.0, 1.112, 3594.0, 1.073)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.112, 3594.0, 1.073)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.696694797, -0.0761809201, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.112, 3594.0, 1.073)
	}
	return jupiterRadii
}

func gasRadius1Gyr260K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 663, 0.850};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.663, 129.0, 0.85)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.67062E-5, 0.0070376319, 0.2201534281, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.663, 129.0, 0.85)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.663, 129.0, 0.85, 215.0, 0.961)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.106532E-6, 0.0027033447, 0.5696053321, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.663, 129.0, 0.85, 215.0, 0.961)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.85, 215.0, 0.961, 318.0, 1.014)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.076256E-6, 0.0010882076, 0.7767853037, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.85, 215.0, 0.961, 318.0, 1.014)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.961, 318.0, 1.014, 464.0, 1.050)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.931406E-7, 4.758113E-4, 0.8923355583, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.961, 318.0, 1.014, 464.0, 1.050)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.014, 464.0, 1.050, 774.0, 1.085)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-8.27315E-8, 2.1532482E-4, 0.9679010435, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.014, 464.0, 1.050, 774.0, 1.085)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.050, 774.0, 1.085, 1292.0, 1.108)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.721891E-8, 1.2129582E-4, 1.013413993, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.050, 774.0, 1.085, 1292.0, 1.108)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.085, 1292.0, 1.108, 2154.0, 1.102)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.233082E-9, 1.7964643E-5, 1.096863604, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.085, 1292.0, 1.108, 2154.0, 1.102)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.108, 2154.0, 1.102, 3594.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.108, 2154.0, 1.102, 3594.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.102, 3594.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.611733926, -0.0664141354, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.102, 3594.0, 1.068)
	}
	return jupiterRadii
}

func gasRadius1Gyr260K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.595, 0.811};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.595, 215.0, 0.811)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.203507E-6, 0.0053336344, 0.0434757282, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.595, 215.0, 0.811)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.595, 215.0, 0.811, 318.0, 0.91)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.044619E-6, 0.0020509472, 0.4645588798, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.595, 215.0, 0.811, 318.0, 0.91)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.811, 318.0, 0.91, 464.0, 0.976)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.244552E-7, 8.6217876E-4, 0.6888621625, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.811, 318.0, 0.91, 464.0, 0.976)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.91, 464.0, 0.976, 774.0, 1.042)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.685316E-7, 4.2154533E-4, 0.8166871439, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.91, 464.0, 0.976, 774.0, 1.042)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.976, 774.0, 1.042, 1292.0, 1.08)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.895551E-8, 1.7450117E-4, 0.9362641716, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.976, 774.0, 1.042, 1292.0, 1.08)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.042, 1292.0, 1.08, 2154.0, 1.085)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.096651E-8, 4.3591061E-5, 1.041986351, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.042, 1292.0, 1.08, 2154.0, 1.085)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.08, 2154.0, 1.085, 3594.0, 1.057)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.08, 2154.0, 1.085, 3594.0, 1.057)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.085, 3594.0, 1.057)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.50478088, -0.0546939939, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.085, 3594.0, 1.057)
	}
	return jupiterRadii
}

func gasRadius1Gyr260K(planet *types.Planet) float64 {
	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	coreMassRadii0 := gasRadius1Gyr260K0coreMass(planet)
	coreMassRadii10 := gasRadius1Gyr260K10coreMass(planet)
	coreMassRadii25 := gasRadius1Gyr260K25coreMass(planet)
	coreMassRadii50 := gasRadius1Gyr260K50coreMass(planet)
	coreMassRadii100 := gasRadius1Gyr260K100coreMass(planet)

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

func gasRadius1Gyr78K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0017354851, 0.0799150115, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 0.857, 28.0, 0.877)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(5.2246604E-7, 0.0017946708, 0.8263396029, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 0.857, 28.0, 0.877)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.857, 28.0, 0.877, 46.0, 0.91)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-7.790213E-6, 0.0024098091, 0.8156328725, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.857, 28.0, 0.877, 46.0, 0.91)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.877, 46.0, 0.91, 77.0, 0.955)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-6.367903E-6, 0.002234865, 0.8206706927, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.877, 46.0, 0.91, 77.0, 0.955)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.91, 77.0, 0.955, 129.0, 1.003)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.234295E-6, 0.0015893417, 0.8517968227, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.91, 77.0, 0.955, 129.0, 1.003)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.955, 129.0, 1.003, 215.0, 1.044)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.2896E-6, 9.2036673E-4, 0.9057329327, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.955, 129.0, 1.003, 215.0, 1.044)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.003, 215.0, 1.044, 318.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-3.581289E-7, 4.238924E-4, 0.9694176408, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.003, 215.0, 1.044, 318.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.044, 318.0, 1.068, 464.0, 1.089)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.456497E-7, 2.5773368E-4, 1.00076937, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.044, 318.0, 1.068, 464.0, 1.089)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.068, 464.0, 1.089, 774.0, 1.113)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-7.951249E-8, 1.7585582E-4, 1.024521621, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.068, 464.0, 1.089, 774.0, 1.113)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.089, 774.0, 1.113, 1292.0, 1.119)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.679996E-8, 4.6291721E-5, 1.087234658, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.089, 774.0, 1.113, 1292.0, 1.119)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.113, 1292.0, 1.119, 2154.0, 1.109)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.518952E-9, 7.4173805E-6, 1.118629332, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.113, 1292.0, 1.119, 2154.0, 1.109)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.119, 2154.0, 1.109, 3594.0, 1.074)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.119, 2154.0, 1.109, 3594.0, 1.074)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.109, 3594.0, 1.074)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.6337261, -0.0683674924, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.109, 3594.0, 1.074)
	}
	return jupiterRadii
}

func gasRadius1Gyr78K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.532, 0.683};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.532, 28.0, 0.683)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-2.664577E-4, 0.0257178683, 0.1718025078, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.532, 28.0, 0.683)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.532, 28.0, 0.683, 46.0, 0.791)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.254115E-5, 0.0106280448, 0.4344470046, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.532, 28.0, 0.683, 46.0, 0.791)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.683, 46.0, 0.791, 77.0, 0.882)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.845347E-5, 0.0052052602, 0.5906055637, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.683, 46.0, 0.791, 77.0, 0.882)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.791, 77.0, 0.882, 129.0, 0.955)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.285707E-6, 0.0024927018, 0.7214009197, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.791, 77.0, 0.882, 129.0, 0.955)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.882, 129.0, 0.955, 215.0, 1.013)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.821807E-6, 0.0013011202, 0.8174721837, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.882, 129.0, 0.955, 215.0, 1.013)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.955, 215.0, 1.013, 318.0, 1.047)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.55487E-7, 6.2617166E-4, 0.9040504793, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.955, 215.0, 1.013, 318.0, 1.047)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.013, 318.0, 1.047, 464.0, 1.075)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.15422E-7, 3.6024083E-4, 0.9542277508, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.013, 318.0, 1.047, 464.0, 1.075)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.047, 464.0, 1.075, 774.0, 1.104)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-9.199743E-8, 2.074412E-4, 0.9985539604, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.047, 464.0, 1.075, 774.0, 1.104)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.075, 774.0, 1.104, 1292.0, 1.113)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.931541E-8, 5.7280146E-5, 1.071236565, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.075, 774.0, 1.104, 1292.0, 1.113)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.104, 1292.0, 1.113, 2154.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.923512E-9, 1.1131679E-5, 1.108505775, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.104, 1292.0, 1.113, 2154.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.113, 2154.0, 1.105, 3594.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.113, 2154.0, 1.105, 3594.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.105, 3594.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.599741752, -0.0644607785, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.105, 3594.0, 1.072)
	}
	return jupiterRadii
}

func gasRadius1Gyr78K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.386, 0.631};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.386, 46.0, 0.631)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-1.796869E-4, 0.0269079438, -0.2265478751, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.386, 46.0, 0.631)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.386, 46.0, 0.631, 77.0, 0.78)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.288589E-5, 0.0088514156, 0.2934214177, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.386, 46.0, 0.631, 77.0, 0.78)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.631, 77.0, 0.78, 129.0, 0.888)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-8.140831E-6, 0.0037539343, 0.5392140468, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.631, 77.0, 0.78, 129.0, 0.888)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.78, 129.0, 0.888, 215.0, 0.97)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.579201E-6, 0.0018407335, 0.6934658653, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.78, 129.0, 0.888, 215.0, 0.97)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.888, 215.0, 0.97, 318.0, 1.018)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(8.1453624E-8, 4.2260464E-4, 0.8753748095, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.888, 215.0, 0.97, 318.0, 1.018)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.97, 318.0, 1.018, 464.0, 1.089)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.059376E-6, 0.0013147336, 0.7070430821, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.97, 318.0, 1.018, 464.0, 1.089)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.018, 464.0, 1.089, 774.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(3.107696E-8, -3.524747E-5, 1.098664081, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.018, 464.0, 1.089, 774.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.089, 774.0, 1.09, 1292.0, 1.105)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.518695E-8, 8.0993771E-5, 1.04239972, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.089, 774.0, 1.09, 1292.0, 1.105)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.09, 1292.0, 1.105, 2154.0, 1.1)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.832022E-9, 1.7742682E-5, 1.093480902, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.09, 1292.0, 1.105, 2154.0, 1.1)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.105, 2154.0, 1.1, 3594.0, 1.069)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.105, 2154.0, 1.1, 3594.0, 1.069)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.1, 3594.0, 1.069)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.564757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.1, 3594.0, 1.069)
	}
	return jupiterRadii
}

func gasRadius1Gyr78K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.610, 0.784};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.61, 129.0, 0.784)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.413627E-5, 0.0062582251, 0.211930602, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.61, 129.0, 0.784)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.61, 129.0, 0.784, 215.0, 0.904)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-3.992445E-6, 0.00276875, 0.4932695331, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.61, 129.0, 0.784, 215.0, 0.904)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.784, 215.0, 0.904, 318.0, 0.97)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.170529E-6, 0.0012646685, 0.6862039668, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.784, 215.0, 0.904, 318.0, 0.97)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.904, 318.0, 0.97, 464.0, 1.021)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.335574E-7, 6.8835694E-4, 0.7949455497, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.904, 318.0, 0.97, 464.0, 1.021)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.97, 464.0, 1.021, 774.0, 1.068)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.318138E-7, 3.1479844E-4, 0.9033125171, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.97, 464.0, 1.021, 774.0, 1.068)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.021, 774.0, 1.068, 1292.0, 1.09)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.993547E-8, 1.0431773E-4, 1.005191703, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.021, 774.0, 1.068, 1292.0, 1.09)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.068, 1292.0, 1.09, 2154.0, 1.091)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-8.950711E-9, 3.2004244E-5, 1.063591617, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.068, 1292.0, 1.09, 2154.0, 1.091)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.091, 3594.0, 1.063)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.09, 2154.0, 1.091, 3594.0, 1.063)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.091, 3594.0, 1.063)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.51078088, -0.0546939939, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.091, 3594.0, 1.063)
	}
	return jupiterRadii
}

func gasRadius1Gyr78K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.570, 0.775};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.57, 215.0, 0.775)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-7.321275E-6, 0.0049022394, 0.0594444444, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.57, 215.0, 0.775)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.57, 215.0, 0.775, 318.0, 0.878)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.92551E-6, 0.002026297, 0.4283528635, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.57, 215.0, 0.775, 318.0, 0.878)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.775, 318.0, 0.878, 464.0, 0.954)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 315.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-6.322147E-7, 0.0010149398, 0.6191812173, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.775, 318.0, 0.878, 464.0, 0.954)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.878, 464.0, 0.954, 774.0, 1.026)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 315.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.942385E-7, 4.7272535E-4, 0.7764742136, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.878, 464.0, 0.954, 774.0, 1.026)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.954, 774.0, 1.026, 1292.0, 1.063)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.251272E-8, 1.5925985E-4, 0.9282012278, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.954, 774.0, 1.026, 1292.0, 1.063)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.026, 1292.0, 1.063, 2154.0, 1.074)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.187852E-8, 5.3694403E-5, 1.013455219, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.026, 1292.0, 1.063, 2154.0, 1.074)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.063, 2154.0, 1.074, 3594.0, 1.053)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.063, 2154.0, 1.074, 3594.0, 1.053)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.074, 3594.0, 1.053)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.38883566, -0.0410204954, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.074, 3594.0, 1.053)
	}
	return jupiterRadii
}

func gasRadius1Gyr78K(planet *types.Planet) float64 {
	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	coreMassRadii0 := gasRadius1Gyr78K0coreMass(planet)
	coreMassRadii10 := gasRadius1Gyr78K10coreMass(planet)
	coreMassRadii25 := gasRadius1Gyr78K25coreMass(planet)
	coreMassRadii50 := gasRadius1Gyr78K50coreMass(planet)
	coreMassRadii100 := gasRadius1Gyr78K100coreMass(planet)

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

func gasRadius1Gyr(planet *types.Planet) float64 {
	temperatureRadii1960 := gasRadius1Gyr1960K(planet)
	temperatureRadii1300 := gasRadius1Gyr1300K(planet)
	temperatureRadii875 := gasRadius1Gyr875K(planet)
	temperatureRadii260 := gasRadius1Gyr260K(planet)
	temperatureRadii78 := gasRadius1Gyr78K(planet)
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
