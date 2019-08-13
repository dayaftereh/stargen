package radius

import (
	"github.com/dayaftereh/stargen/mathf"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

func gasRadius300Myr(planet *types.Planet) float64 {
	temperatureRadii1960 := gasRadius300Myr1960K(planet)
	temperatureRadii1300 := gasRadius300Myr1300K(planet)
	temperatureRadii875 := gasRadius300Myr875K(planet)
	temperatureRadii260 := gasRadius300Myr260K(planet)
	temperatureRadii78 := gasRadius300Myr78K(planet)
	temperatureRadii0 := 0.0

	temperature := planet.EstimatedTemperature

	var jupiterRadii float64

	if temperature <= 78.0 {
		jupiterRadii1 := PlanetRadiusHelper(temperature, 0.0, temperatureRadii0, 78.0, temperatureRadii78, 260.0, temperatureRadii260)
		jupiterRadii2 := PlanetRadiusHelper2(temperature, 78.0, temperatureRadii78, 260.0, temperatureRadii260)
		jupiterRadii = RangeAdjust(temperature, jupiterRadii1, jupiterRadii2, 0.0, 78.0)
	} else if temperature <= 260.0 {
		//jupiterRadii1 := PlanetRadiusHelper(temperature, 0.0, temperatureRadii0, 78.0,temperatureRadii78, 260.0, temperatureRadii260, false);
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

func gasRadius300Myr1960K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-0.0023409248, 0.1486173223, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 28.0, 2.326, 46.0, 1.883)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(3.5282715E-4, -0.0507203204, 3.469552483, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 28.0, 2.326, 46.0, 1.883)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 2.326, 46.0, 1.883, 77.0, 1.656)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(4.1652964E-5, -0.0124458952, 2.367373509, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 2.326, 46.0, 1.883, 77.0, 1.656)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.883, 77.0, 1.656, 129.0, 1.455)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(2.1521998E-5, -0.0082989163, 2.167412625, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.883, 77.0, 1.656, 129.0, 1.455)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.656, 129.0, 1.455, 215.0, 1.378)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(2.888012E-6, -0.001888825, 1.650599014, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.656, 129.0, 1.455, 215, 1.378)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.455, 215.0, 1.378, 318.0, 1.342)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(9.910636E-7, -8.777515E-4, 1.52090465, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.455, 215.0, 1.378, 318.0, 1.342)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.378, 318.0, 1.342, 464.0, 1.327)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(9.0897815E-8, -1.738218E-4, 1.3880083387, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.378, 318.0, 1.342, 464.0, 1.327)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.342, 464.0, 1.327, 774.0, 1.308)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(8.1016701E-8, -1.61589E-4, 1.384534724, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.342, 464.0, 1.327, 774.0, 1.308)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.327, 774.0, 1.308, 1292.0, 1.311)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-8.34155E-10, 7.514871E-6, 1.302683212, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.327, 774.0, 1.308, 1292.0, 1.311)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.308, 1292.0, 1.311, 2154.0, 1.315)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.136757E-8, 4.3813022E-5, 1.273369053, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.308, 1292.0, 1.311, 2154.0, 1.315)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.311, 2154.0, 1.315, 3594.0, 1.284)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.311, 2154.0, 1.315, 3594.0, 1.284)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.315, 3594.0, 1.284)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.315, 3594.0, 1.284)
	}

	return jupiterRadii
}

func gasRadius300Myr1960K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 1.102, 28.0, 1.388)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-7.490421E-4, 0.0597068966, 0.3034559387, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 1.102, 28.0, 1.388)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.102, 28.0, 1.388, 46.0, 1.465)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-1.156097E-4, 0.0128328944, 1.119316948, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.102, 28.0, 1.388, 46.0, 1.465)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.388, 46.0, 1.465, 77.0, 1.422)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-2.017998E-7, -0.0013622754, 1.528091677, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.388, 46.0, 1.465, 77.0, 1.422)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.465, 77.0, 1.422, 129.0, 1.349)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(8.1505535E-6, -0.0030828602, 1.611055602, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.465, 77.0, 1.422, 129.0, 1.349)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.422, 129.0, 1.349, 215.0, 1.325)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(7.5739385E-7, -5.5396133E-4, 1.406006318, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.422, 129.0, 1.349, 215.0, 1.325)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.349, 215.0, 1.325, 318.0, 1.311)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(4.0833636E-7, -3.535656E-4, 1.382141258, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.349, 215.0, 1.325, 318.0, 1.311)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.325, 318.0, 1.311, 464.0, 1.306)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.713368E-9, -3.212472E-5, 1.321490048, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.325, 318.0, 1.311, 464.0, 1.306)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.311, 464.0, 1.306, 774.0, 1.295)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(6.3838633E-8, -1.145161E-4, 1.345391268, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.311, 464.0, 1.306, 774.0, 1.295)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.306, 774.0, 1.295, 1292.0, 1.304)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-7.546348E-9, 3.2965273E-5, 1.274005715, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.306, 774.0, 1.295, 1292.0, 1.304)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.295, 1292.0, 1.304, 2154.0, 1.31)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.20738E-8, 4.8566876E-5, 1.261405958, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.295, 1292.0, 1.304, 2154.0, 1.31)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.304, 2154.0, 1.31, 3594.0, 1.281)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.304, 2154.0, 1.31, 3594.0, 1.281)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.31, 3594.0, 1.281)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.744773054, -0.0566473508, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.31, 3594.0, 1.281)
	}

	return jupiterRadii
}

func gasRadius300Myr1960K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.493, 46.0, 0.945)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-3.88706E-4, 0.0538753566, -0.7107644649, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.493, 46.0, 0.945)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.493, 46.0, 0.945, 77.0, 1.133)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-5.29089E-5, 0.0125723116, 0.4786289127, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.493, 46.0, 0.945, 77.0, 1.133)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.945, 77.0, 1.133, 129.0, 1.22)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.343159E-6, 0.0035977678, 0.9113674749, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.945, 77.0, 1.133, 129.0, 1.22)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.133, 129.0, 1.22, 215.0, 1.253)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.311104E-6, 8.347406E-4, 1.134136539, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.133, 129.0, 1.22, 215.0, 1.253)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.22, 215.0, 1.253, 318.0, 1.267)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-3.258145E-7, 3.0958146E-4, 1.201500762, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.22, 215.0, 1.253, 318.0, 1.267)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.253, 318.0, 1.267, 464.0, 1.275)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.130893E-7, 1.4323034E-4, 1.232888792, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.253, 318.0, 1.267, 464.0, 1.275)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.267, 464.0, 1.275, 774.0, 1.276)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(3.8071532E-8, -4.390675E-5, 1.287176083, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.267, 464.0, 1.275, 774.0, 1.276)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.275, 774.0, 1.276, 1292.0, 1.294)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.677399E-8, 6.9404099E-5, 1.232330123, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.275, 774.0, 1.276, 1292.0, 1.294)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.276, 1292.0, 1.294, 2154.0, 1.304)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.318459E-8, 5.7035028E-5, 1.242319307, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.276, 1292.0, 1.294, 2154.0, 1.304)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.294, 2154.0, 1.304, 3594.0, 1.277)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.294, 2154.0, 1.304, 3594.0, 1.277)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.304, 3594.0, 1.277)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.708788706, -0.052740637, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.304, 3594.0, 1.277)
	}

	return jupiterRadii
}

func gasRadius300Myr1960K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	var jupiterRadii float64

	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.801, 129.0, 1.03)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-2.230627E-5, 0.008998937, 0.2403357023, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.801, 129.0, 1.03)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.801, 129.0, 1.03, 215.0, 1.144)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.496578E-6, 0.0028724042, 0.7342874095, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.801, 129.0, 1.03, 215.0, 1.144)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.03, 215.0, 1.144, 318.0, 1.193)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.002814E-6, 0.0010102282, 0.9731560336, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.03, 215.0, 1.144, 318.0, 1.193)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.144, 318.0, 1.193, 464.0, 1.226)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.612655E-7, 5.0853703E-4, 1.067817838, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.144, 318.0, 1.193, 464.0, 1.226)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.193, 464.0, 1.226, 774.0, 1.245)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.744883E-9, 6.3450487E-5, 1.19693464, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.193, 464.0, 1.226, 774.0, 1.245)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.226, 774.0, 1.245, 1292.0, 1.276)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.9916E-8, 1.2165201E-4, 1.168763301, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.226, 774.0, 1.245, 1292.0, 1.276)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.245, 1292.0, 1.276, 2154.0, 1.292)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.469994E-8, 6.9217479E-5, 1.211109098, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.245, 1292.0, 1.276, 2154.0, 1.292)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.276, 2154.0, 1.292, 3594.0, 1.27)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.276, 2154.0, 1.292, 3594.0, 1.27)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.292, 3594.0, 1.27)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.621827834, -0.0429738523, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.292, 3594.0, 1.27)
	}

	return jupiterRadii
}

func gasRadius300Myr1960K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	var jupiterRadii float64

	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {

		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.669, 215.0, 0.939)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.065249E-5, 0.0068039927, -0.0314469102, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.669, 215.0, 0.939)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.669, 215.0, 0.939, 318.0, 1.055)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.514914E-6, 0.0024666628, 0.5249194058, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.669, 215.0, 0.939, 318.0, 1.055)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.939, 318.0, 1.055, 464.0, 1.128)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-6.791171E-7, 0.0010310696, 0.7957949066, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.939, 318.0, 1.055, 464.0, 1.128)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.055, 464.0, 1.128, 774.0, 1.187)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.016244E-7, 3.1613354E-4, 1.003193355, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.055, 464.0, 1.128, 774.0, 1.187)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.128, 774.0, 1.187, 1292.0, 1.242)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.340218E-8, 2.1650651E-4, 1.051415926, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.128, 774.0, 1.187, 1292.0, 1.242)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.187, 1292.0, 1.242, 2154.0, 1.27)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.833398E-8, 9.5661492E-5, 1.149009604, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.187, 1292.0, 1.242, 2154.0, 1.27)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.242, 2154.0, 1.27, 3594.0, 1.256)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.242, 2154.0, 1.27, 3594.0, 1.256)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.27, 3594.0, 1.256)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.47989044, -0.0273469969, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.27, 3594.0, 1.256)
	}

	return jupiterRadii
}

func gasRadius300Myr1960K(planet *types.Planet) float64 {
	//totalEarthMasses := := planet.Mass*constants.SunMassInEarthMasses

	coreMassRadii0 := gasRadius300Myr1960K0coreMass(planet)
	coreMassRadii10 := gasRadius300Myr1960K10coreMass(planet)
	coreMassRadii25 := gasRadius300Myr1960K25coreMass(planet)
	coreMassRadii50 := gasRadius300Myr1960K50coreMass(planet)
	coreMassRadii100 := gasRadius300Myr1960K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if coreEarthMasses <= 10.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1960K_0core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1960K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 0.0, 10.0);*/
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1960K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1960K_25core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1960K_25core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1960K_50core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1960K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1960K_100core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		/*upper_fraction = coreEarthMasses / 100.0;
		  jupiterRadii = gas_radius_300Myr_1960K_100core_mass(totalEarthMasses, the_planet) * pow1_4(upper_fraction);*/
		/*jupiterRadii1 := gas_radius_300Myr_1960K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1960K_100core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, jupiterRadii1, 100.0, jupiterRadii2);*/
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}

	return jupiterRadii
}

func gasRadius300Myr1300K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0100049656, 0.3344961803, 0.0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 2.795, 28.0, 1.522)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(0.0036515152, -0.2800454545, 6.500484848, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 2.795, 28.0, 1.522)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 2.795, 28.0, 1.522, 46.0, 1.345)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(1.4143077E-4, -0.02029921, 1.97949616, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 2.795, 28.0, 1.522, 46.0, 1.345)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.522, 46.0, 1.345, 77.0, 1.255)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(3.1503184E-5, -0.0067781174, 1.590132665, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.522, 46.0, 1.345, 77.0, 1.255)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.345, 77.0, 1.255, 129.0, 1.24)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(1.0791787E-6, -5.107723E-4, 1.28793102, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.345, 77.0, 1.255, 129.0, 1.24)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.255, 129.0, 1.24, 215.0, 1.228)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.362393E-8, -1.107693E-4, 1.25568069, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.255, 129.0, 1.24, 215.0, 1.228)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.24, 215.0, 1.228, 318.0, 1.212)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(4.5881091E-7, -3.99886E-4, 1.29276696, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.24, 215.0, 1.228, 318.0, 1.212)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.228, 318.0, 1.212, 464.0, 1.206)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(4.0603608E-8, -7.284791E-5, 1.231059637, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.228, 318.0, 1.212, 464.0, 1.206)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.212, 464.0, 1.206, 774.0, 1.199)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(5.2918075E-8, -8.809322E-5, 1.235482205, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.212, 464.0, 1.206, 774.0, 1.199)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.206, 774.0, 1.199, 1292.0, 1.21)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.127259E-8, 6.5184687E-5, 1.161290949, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.206, 774.0, 1.199, 1292.0, 1.21)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.199, 1292.0, 1.21, 2154.0, 1.203)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.427462E-9, 1.4028384E-5, 1.202604459, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.199, 1292.0, 1.21, 2154.0, 1.203)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.21, 2154.0, 1.203, 3594.0, 1.17)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.21, 2154.0, 1.203, 3594.0, 1.17)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.203, 3594.0, 1.17)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.697741752, -0.0644607785, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.203, 3594.0, 1.17)
	}

	return jupiterRadii
}

func gasRadius300Myr1300K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {

		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.801, 28.0, 1.012)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-5.10101E-4, 0.0421363636, 0.2321010101, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.801, 28.0, 1.012)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.801, 28.0, 1.012, 46.0, 1.091)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.784434E-5, 0.0094093702, 0.8017275986, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.801, 28.0, 1.012, 46.0, 1.091)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.012, 46.0, 1.091, 77.0, 1.124)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-2.630871E-6, 0.0013881132, 1.032713713, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.012, 46.0, 1.091, 77.0, 1.124)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.091, 77.0, 1.124, 129.0, 1.168)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-4.699126E-6, 0.0018141739, 1.012169732, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.091, 77.0, 1.124, 129.0, 1.168)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.124, 129.0, 1.168, 215.0, 1.185)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.045896E-6, 5.5746278E-4, 1.113492063, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.124, 129.0, 1.168, 215.0, 1.185)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.168, 215.0, 1.185, 318.0, 1.185)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(8.2521868E-8, -4.398416E-5, 1.19064202, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.168, 215.0, 1.185, 318.0, 1.185)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.185, 318.0, 1.185, 464.0, 1.188)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.506128E-8, 5.5785869E-5, 1.171816871, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.185, 318.0, 1.185, 464.0, 1.188)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.185, 464.0, 1.188, 774.0, 1.188)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(3.7304385E-8, -4.618283E-5, 1.201397348, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.185, 464.0, 1.188, 774.0, 1.188)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.188, 774.0, 1.188, 1292.0, 1.204)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.658587E-8, 8.581443E-5, 1.137506586, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.188, 774.0, 1.188, 1292.0, 1.204)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.188, 1292.0, 1.204, 2154.0, 1.199)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.832022E-9, 1.7742682E-5, 1.192480902, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.188, 1292.0, 1.204, 2154.0, 1.199)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.204, 2154.0, 1.199, 3594.0, 1.168)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.204, 2154.0, 1.199, 3594.0, 1.168)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.199, 3594.0, 1.168)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.663757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.199, 3594.0, 1.168)
	}

	return jupiterRadii
}

func gasRadius300Myr1300K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.447, 46.0, 0.793)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.770829E-4, 0.0397263551, -0.4481049667, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.447, 46.0, 0.793)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.447, 46.0, 0.793, 77.0, 0.968)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-4.41493E-5, 0.0110755254, 0.3769457532, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.447, 46.0, 0.793, 77.0, 0.968)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.793, 77.0, 0.968, 129.0, 1.071)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.88761E-6, 0.0040176169, 0.7172671405, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.793, 77.0, 0.968, 129.0, 1.071)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.968, 129.0, 1.071, 215.0, 1.124)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.079249E-6, 0.0013315408, 0.9338320234, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.968, 129.0, 1.071, 215.0, 1.124)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.071, 215.0, 1.124, 318.0, 1.147)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.11689E-7, 4.9603121E-4, 1.041006115, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.071, 215.0, 1.124, 318.0, 1.147)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.124, 318.0, 1.147, 464.0, 1.161)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.253963E-7, 1.9395035E-4, 1.098004368, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.124, 318.0, 1.147, 464.0, 1.161)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.147, 464.0, 1.161, 774.0, 1.173)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(4.5427114E-9, 3.3085801E-5, 1.144670161, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.147, 464.0, 1.161, 774.0, 1.173)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.161, 774.0, 1.173, 1292.0, 1.195)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.245741E-8, 1.0952805E-4, 1.107669742, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.161, 774.0, 1.173, 1292.0, 1.195)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.173, 1292.0, 1.195, 2154.0, 1.193)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.740531E-9, 2.4353686E-5, 1.176456029, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.173, 1292.0, 1.195, 2154.0, 1.193)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.195, 2154.0, 1.193, 3594.0, 1.164)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.195, 2154.0, 1.193, 3594.0, 1.164)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.193, 3594.0, 1.164)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.627773054, -0.0566473508, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.193, 3594.0, 1.164)
	}

	return jupiterRadii
}

func gasRadius300Myr1300K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses
	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.719, 129.0, 0.921)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.871224E-5, 0.0077393378, 0.2340158863, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.719, 129.0, 0.921)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.719, 129.0, 0.921, 215.0, 1.033)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.270793E-6, 0.0027714785, 0.6345495454, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.719, 129.0, 0.921, 215.0, 1.033)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.921, 215.0, 1.033, 318.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.025782E-6, 0.0010418872, 0.8564110054, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.921, 215.0, 1.033, 318.0, 1.084)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.033, 318.0, 1.084, 464.0, 1.119)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.20565E-7, 4.9040786E-4, 0.9604671163, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.033, 318.0, 1.084, 464.0, 1.119)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.084, 464.0, 1.119, 774.0, 1.148)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-4.07039E-8, 1.4393981E-4, 1.060975313, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.084, 464.0, 1.119, 774.0, 1.148)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.119, 774.0, 1.148, 1292.0, 1.179)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.000376E-8, 1.4249333E-4, 1.061675457, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.119, 774.0, 1.148, 1292.0, 1.179)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.148, 1292.0, 1.179, 2154.0, 1.183)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-9.859221E-9, 3.8615247E-5, 1.145566744, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.148, 1292.0, 1.179, 2154.0, 1.183)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.179, 2154.0, 1.183, 3594.0, 1.157)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.179, 2154.0, 1.183, 3594.0, 1.157)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.183, 3594.0, 1.157)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.572796532, -0.05078728, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.183, 3594.0, 1.157)
	}

	return jupiterRadii
}

func gasRadius300Myr1300K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.627, 215.0, 0.863)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-9.12576E-6, 0.0058834474, 0.0198970566, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.627, 215.0, 0.863)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.627, 215.0, 0.863, 318.0, 0.968)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.22355E-6, 0.0022045698, 0.4198011015, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.627, 215.0, 0.863, 318.0, 0.968)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.863, 318.0, 0.968, 464.0, 1.036)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.615702E-7, 9.0490131E-4, 0.7370296067, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.863, 318.0, 0.968, 464.0, 1.036)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.968, 464.0, 1.036, 774.0, 1.101)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.43652E-7, 3.8751855E-4, 0.8871190846, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.968, 464.0, 1.036, 774.0, 1.101)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.036, 774.0, 1.101, 1292.0, 1.148)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.313927E-8, 2.0051933E-4, 0.9976325013, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.036, 774.0, 1.101, 1292.0, 1.148)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.101, 1292.0, 1.148, 2154.0, 1.163)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.268764E-8, 6.1123E-5, 1.090208105, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.101, 1292.0, 1.148, 2154.0, 1.163)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.148, 2154.0, 1.163, 3594.0, 1.146)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.148, 2154.0, 1.163, 3594.0, 1.146)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.163, 3594.0, 1.146)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.417866963, -0.0332070677, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.163, 3594.0, 1.146)
	}

	return jupiterRadii
}

func gasRadius300Myr1300K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius300Myr1300K0coreMass(planet)
	coreMassRadii10 := gasRadius300Myr1300K10coreMass(planet)
	coreMassRadii25 := gasRadius300Myr1300K25coreMass(planet)
	coreMassRadii50 := gasRadius300Myr1300K50coreMass(planet)
	coreMassRadii100 := gasRadius300Myr1300K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if coreEarthMasses <= 10.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1300K_0core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1300K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 0.0, 10.0);*/
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1300K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1300K_25core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1300K_25core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1300K_50core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		/*jupiterRadii1 := gas_radius_300Myr_1300K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_1300K_100core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		//jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, jupiterRadii1, 100.0, jupiterRadii2);
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}

	return jupiterRadii
}

func gasRadius300Myr875K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.004000191, 0.1618267762, 0.0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.595, 28.0, 1.395)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(3.8749565E-4, -0.0356191223, 2.088538837, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.595, 28.0, 1.395)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.595, 28.0, 1.395, 46.0, 1.27)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(9.3665423E-5, -0.0138756858, 1.710085509, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.595, 28.0, 1.395, 46.0, 1.27)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.395, 46.0, 1.27, 77.0, 1.197)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(2.9530031E-5, -0.0059870325, 1.48291795, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.395, 46.0, 1.27, 77.0, 1.197)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.27, 77.0, 1.197, 129.0, 1.202)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.033808E-6, 3.0911825E-4, 1.179327341, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.27, 77.0, 1.197, 129.0, 1.202)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.197, 129.0, 1.202, 215.0, 1.198)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-3.189655E-7, 6.3212521E-5, 1.199153491, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.197, 129.0, 1.202, 215.0, 1.198)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.202, 215.0, 1.198, 318.0, 1.187)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(2.9136362E-7, -2.620929E-4, 1.240881696, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.202, 215.0, 1.198, 318.0, 1.187)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.198, 318.0, 1.187, 464.0, 1.182)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(4.6805591E-8, -7.084855E-5, 1.20479667, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.198, 318.0, 1.187, 464.0, 1.182)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.187, 464.0, 1.182, 774.0, 1.178)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(4.1230371E-8, -6.394642E-5, 1.202794407, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.187, 464.0, 1.182, 774.0, 1.178)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.182, 774.0, 1.178, 1292.0, 1.189)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.463518E-8, 7.2131794E-5, 1.136928334, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.182, 774.0, 1.178, 1292.0, 1.189)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.178, 1292.0, 1.189, 2154.0, 1.178)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.713332E-9, 3.4811211E-6, 1.192370187, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.178, 1292.0, 1.189, 2154.0, 1.178)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.189, 2154.0, 1.178, 3594.0, 1.144)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.189, 2154.0, 1.178, 3594.0, 1.144)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.178, 3594.0, 1.144)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.687733926, -0.0664141354, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.178, 3594.0, 1.144)
	}

	return jupiterRadii
}

func gasRadius300Myr875K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.755, 28.0, 0.956)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-4.78753E-4, 0.0398166144, 0.2164771856, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.755, 28.0, 0.956)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.755, 28.0, 0.956, 46.0, 1.035)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-5.73111E-5, 0.00862991, 0.7592944188, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.755, 28.0, 0.956, 46.0, 1.035)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.956, 46.0, 1.035, 77.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-7.459117E-6, 0.0024981165, 0.9358701306, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.956, 46.0, 1.035, 77.0, 1.084)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.035, 77.0, 1.084, 129.0, 1.134)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.029686E-6, 0.0019976537, 0.9600016722, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.035, 77.0, 1.084, 129.0, 1.134)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.084, 129.0, 1.134, 215.0, 1.157)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.260929E-6, 7.0120156E-4, 1.064528125, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.084, 129.0, 1.134, 215.0, 1.157)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.134, 215.0, 1.157, 318.0, 1.160)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-6.943588E-9, 3.2827146E-5, 1.150263131, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.134, 215.0, 1.157, 318.0, 1.160)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.157, 318.0, 1.160, 464.0, 1.164)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.178516E-8, 5.2253258E-5, 1.146597707, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.157, 318.0, 1.160, 464.0, 1.164)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.160, 464.0, 1.164, 774.0, 1.168)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(1.9389255E-8, -1.110067E-5, 1.164976283, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.160, 464.0, 1.164, 774.0, 1.168)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.164, 774.0, 1.168, 1292.0, 1.183)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.854954E-8, 8.7940877E-5, 1.117037105, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.164, 774.0, 1.168, 1292.0, 1.183)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.168, 1292.0, 1.183, 2154.0, 1.174)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.117892E-9, 7.1954199E-6, 1.18224663, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.168, 1292.0, 1.183, 2154.0, 1.174)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.183, 2154.0, 1.174, 3594.0, 1.142)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.183, 2154.0, 1.174, 3594.0, 1.142)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.174, 3594.0, 1.142)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.653749577, -0.0625074216, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.174, 3594.0, 1.142)
	}
	return jupiterRadii
}

func gasRadius300Myr875K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)
	var jupiterRadii float64
	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.438, 46.0, 0.767)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.604418E-4, 0.0375504718, -0.4092268305, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.438, 46.0, 0.767)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.438, 46.0, 0.767, 77.0, 0.938)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-4.2363E-5, 0.0107267781, 0.3632083171, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.438, 46.0, 0.767, 77.0, 0.938)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.767, 77.0, 0.938, 129.0, 1.042)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.689922E-6, 0.003996124, 0.68775, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.767, 77.0, 0.938, 129.0, 1.042)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.938, 129.0, 1.042, 215.0, 1.099)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.273973E-6, 0.0014450376, 0.8934313454, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.938, 129.0, 1.042, 215.0, 1.099)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.042, 215.0, 1.099, 318.0, 1.123)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.231726E-7, 5.1186072E-4, 1.013133601, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.042, 215.0, 1.099, 318.0, 1.123)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.099, 318.0, 1.123, 464.0, 1.138)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.191944E-7, 1.9594972E-4, 1.072741401, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.099, 318.0, 1.123, 464.0, 1.138)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.123, 464.0, 1.138, 774.0, 1.153)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-9.476517E-9, 6.0119025E-5, 1.112145029, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.123, 464.0, 1.138, 774.0, 1.153)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.138, 774.0, 1.153, 1292.0, 1.174)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.358044E-8, 1.0991773E-4, 1.088040915, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.138, 774.0, 1.153, 1292.0, 1.174)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.153, 1292.0, 1.174, 2154.0, 1.169)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.832022E-9, 1.7742682E-5, 1.162480902, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.153, 1292.0, 1.174, 2154.0, 1.169)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.174, 2154.0, 1.169, 3594.0, 1.138)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.174, 2154.0, 1.169, 3594.0, 1.138)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.169, 3594.0, 1.138)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.633757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.169, 3594.0, 1.138)
	}

	return jupiterRadii
}

func gasRadius300Myr875K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.702, 129.0, 0.899)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.801548E-5, 0.00749965, 0.2313407191, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.702, 129.0, 0.899)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.702, 129.0, 0.899, 215.0, 1.011)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.219424E-6, 0.0027538076, 0.6139742641, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.702, 129.0, 0.899, 215.0, 1.011)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.899, 215.0, 1.011, 318.0, 1.063)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.064772E-6, 0.0010723781, 0.8296578179, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.899, 215.0, 1.011, 318.0, 1.063)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.011, 318.0, 1.063, 464.0, 1.098)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.064167E-7, 4.7934391E-4, 0.9415547224, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.011, 318.0, 1.063, 464.0, 1.098)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.063, 464.0, 1.098, 774.0, 1.129)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-5.315875E-8, 1.6581053E-4, 1.03250878, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.063, 464.0, 1.098, 774.0, 1.129)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.098, 774.0, 1.129, 1292.0, 1.158)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.972787E-8, 1.3806234E-4, 1.045939764, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.098, 774.0, 1.129, 1292.0, 1.158)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.129, 1292.0, 1.158, 2154.0, 1.159)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-8.649041E-9, 3.0964689E-5, 1.132431155, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.129, 1292.0, 1.158, 2154.0, 1.159)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.158, 2154.0, 1.159, 3594.0, 1.132)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3494.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.158, 2154.0, 1.159, 3594.0, 1.132)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.159, 3594.0, 1.132)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3494.0)
	} else {
		jupiterRadii = mathf.LnTrend(1.563788706, -0.052740637, totalEarthMasses)
	}
	return jupiterRadii
}

func gasRadius300Myr875K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.618, 215.0, 0.847)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.797834E-6, 0.0056892457, 0.0304920635, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.618, 215.0, 0.847)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.618, 215.0, 0.847, 318.0, 0.95)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.145569E-6, 0.0021435881, 0.4853074765, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.618, 215.0, 0.847, 318.0, 0.95)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.847, 318.0, 0.95, 464.0, 1.018)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.544961E-7, 8.9936934E-4, 0.7200734098, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.847, 318.0, 0.95, 464.0, 1.018)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.95, 464.0, 1.018, 774.0, 1.084)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.545424E-7, 4.0422677E-4, 0.8637111493, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.95, 464.0, 1.018, 774.0, 1.084)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.018, 774.0, 1.084, 1292.0, 1.128)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.146447E-8, 1.9126768E-4, 0.9667899426, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.018, 774.0, 1.084, 1292.0, 1.128)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.084, 1292.0, 1.128, 2154.0, 1.14)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.177913E-8, 5.451197E-5, 1.077232978, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.084, 1292.0, 1.128, 2154.0, 1.14)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.128, 2154.0, 1.14, 3594.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.128, 2154.0, 1.14, 3594.0, 1.121)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.14, 3594.0, 1.121)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		jupiterRadii = mathf.LnTrend(1.424851312, -0.0371137816, totalEarthMasses)
	}
	return jupiterRadii
}

func gasRadius300Myr875K(planet *types.Planet) float64 {
	coreMassRadii0 := gasRadius300Myr875K0coreMass(planet)
	coreMassRadii10 := gasRadius300Myr875K10coreMass(planet)
	coreMassRadii25 := gasRadius300Myr875K25coreMass(planet)
	coreMassRadii50 := gasRadius300Myr875K50coreMass(planet)
	coreMassRadii100 := gasRadius300Myr875K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if coreEarthMasses <= 10.0 {
		/*jupiterRadii1 := gas_radius_300Myr_875K_0core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_875K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 0.0, 10.0);*/
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		/*jupiterRadii1 := gas_radius_300Myr_875K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_875K_25core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		/*jupiterRadii1 := gas_radius_300Myr_875K_25core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_875K_50core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		/*jupiterRadii1 := gas_radius_300Myr_875K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_875K_100core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		/*upper_fraction = coreEarthMasses / 100.0;
		  jupiterRadii = gas_radius_300Myr_875K_100core_mass(totalEarthMasses, the_planet) * pow1_4(upper_fraction);*/
		/*jupiterRadii1 := gas_radius_300Myr_875K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_875K_100core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, jupiterRadii1, 100.0, jupiterRadii2);*/
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}

	return jupiterRadii
}

func gasRadius300Myr260K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0037408327, 0.1520647441, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.504, 28.0, 1.325)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(3.6381052E-4, -0.0326442006, 1.953810171, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 1.504, 28.0, 1.325)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.504, 28.0, 1.325, 46.0, 1.222)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(8.1888669E-5, -0.0117819838, 1.590694828, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 1.504, 28.0, 1.325, 46.0, 1.222)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.325, 46.0, 1.222, 77.0, 1.169)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(2.3610571E-5, -0.0046137777, 1.384273805, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.325, 46.0, 1.222, 77.0, 1.169)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.222, 77.0, 1.169, 129.0, 1.182)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.811594E-6, 6.2318841E-4, 1.131755435, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.222, 77.0, 1.169, 129.0, 1.182)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.169, 129.0, 1.182, 215.0, 1.182)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.623209E-7, 1.5903837E-4, 1.169177531, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.169, 129.0, 1.182, 215.0, 1.182)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.182, 215.0, 1.182, 318.0, 1.173)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(2.4088908E-7, -2.157725E-4, 1.217255994, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.182, 215.0, 1.182, 318.0, 1.173)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.182, 318.0, 1.173, 464.0, 1.169)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(5.3007574E-8, -6.884918E-5, 1.189533702, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.182, 318.0, 1.173, 464.0, 1.169)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.173, 464.0, 1.169, 774.0, 1.168)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(2.9542666E-8, -3.979963E-5, 1.181106609, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.173, 464.0, 1.169, 774.0, 1.168)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.169, 774.0, 1.168, 1292.0, 1.179)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.379453E-8, 7.0395017E-5, 1.127768988, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.169, 774.0, 1.168, 1292.0, 1.179)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.168, 1292.0, 1.179, 2154.0, 1.169)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-4.915612E-9, 5.3382705E-6, 1.180308408, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.168, 1292.0, 1.179, 2154.0, 1.169)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.179, 2154.0, 1.169, 3594.0, 1.136)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.179, 2154.0, 1.169, 3594.0, 1.136)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.169, 3594.0, 1.136)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.663741752, -0.0644607785, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.169, 3594.0, 1.136)
	}

	return jupiterRadii
}

func gasRadius300Myr260K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.727, 28.0, 0.921)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-4.491466E-4, 0.0378479624, 0.2133880181, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.727, 28.0, 0.921)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.727, 28.0, 0.921, 46.0, 1.004)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-5.526297E-5, 0.0087005706, 0.7207101895, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.727, 28.0, 0.921, 46.0, 1.004)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.921, 46.0, 1.004, 77.0, 1.063)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.065054E-5, 0.0032132425, 0.878727391, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.921, 46.0, 1.004, 77.0, 1.063)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.004, 77.0, 1.063, 129.0, 1.116)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-5.279225E-6, 0.0021067512, 0.9320806856, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.004, 77.0, 1.063, 129.0, 1.116)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.063, 129.0, 1.116, 215.0, 1.141)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.281238E-6, 7.3144355E-4, 1.042964864, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.063, 129.0, 1.116, 215.0, 1.141)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.116, 215.0, 1.141, 318.0, 1.146)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.991084E-8, 6.4486167E-5, 1.128518103, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.116, 215.0, 1.141, 318.0, 1.146)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.141, 318.0, 1.146, 464.0, 1.152)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.767774E-8, 7.8379887E-5, 1.12589656, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.141, 318.0, 1.146, 464.0, 1.152)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.146, 464.0, 1.152, 774.0, 1.158)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(1.1597452E-8, 4.9971931E-6, 1.147184417, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.146, 464.0, 1.152, 774.0, 1.158)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.152, 774.0, 1.158, 1292.0, 1.173)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.770889E-8, 8.6204101E-5, 1.107877759, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.152, 774.0, 1.158, 1292.0, 1.173)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.158, 1292.0, 1.173, 2154.0, 1.165)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.320172E-9, 9.0525693E-6, 1.170184852, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.158, 1292.0, 1.173, 2154.0, 1.165)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.173, 2154.0, 1.165, 3594.0, 1.134)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.173, 2154.0, 1.165, 3594.0, 1.134)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.165, 3594.0, 1.134)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.629757403, -0.0605540647, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.165, 3594.0, 1.134)
	}

	return jupiterRadii
}

func gasRadius300Myr260K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.433, 0.754};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.433, 46.0, 0.754)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-2.526882E-4, 0.0365322581, -0.3917956989, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.433, 46.0, 0.754)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.433, 46.0, 0.754, 77.0, 0.923)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-4.15857E-5, 0.0105666537, 0.3559292654, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.433, 46.0, 0.754, 77.0, 0.923)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.754, 77.0, 0.923, 129.0, 1.027)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.605662E-6, 0.0039787664, 0.6735869565, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.754, 77.0, 0.923, 129.0, 1.027)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.923, 129.0, 1.027, 215.0, 1.085)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.284128E-6, 0.0014601586, 0.8766497149, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.923, 129.0, 1.027, 215.0, 1.085)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.027, 215.0, 1.085, 318.0, 1.11)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.07149E-7, 5.1302884E-4, 0.9981417598, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.027, 215.0, 1.085, 318.0, 1.11)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.085, 318.0, 1.11, 464.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.421611E-7, 2.2760832E-4, 1.051996451, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.085, 318.0, 1.11, 464.0, 1.127)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.11, 464.0, 1.127, 774.0, 1.143)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.337242E-8, 6.8167958E-5, 1.098249096, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.11, 464.0, 1.127, 774.0, 1.143)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.127, 774.0, 1.143, 1292.0, 1.164)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.358044E-8, 1.0991773E-4, 1.078040915, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.127, 774.0, 1.143, 1292.0, 1.164)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.143, 1292.0, 1.164, 2154.0, 1.159)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-6.228682E-9, 1.5663572E-9, 1.154159978, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.143, 1292.0, 1.164, 2154.0, 1.159)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.164, 2154.0, 1.159, 3594.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.164, 2154.0, 1.159, 3594.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.159, 3594.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.593773054, -0.0566473508, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.159, 3594.0, 1.13)
	}

	return jupiterRadii
}

func gasRadius300Myr260K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.693, 0.888};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.693, 129.0, 0.888)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.782103E-5, 0.0074211325, 0.2272336957, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.693, 129.0, 0.888)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.693, 129.0, 0.888, 215.0, 0.999)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.157901E-6, 0.0027210156, 0.6061806133, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.693, 129.0, 0.888, 215.0, 0.999)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.888, 215.0, 0.999, 318.0, 1.051)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.037265E-6, 0.0010577167, 0.8195384912, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.888, 215.0, 0.999, 318.0, 1.051)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.999, 318.0, 1.051, 464.0, 1.087)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-3.072889E-7, 4.8687525E-4, 0.9272479522, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.999, 318.0, 1.051, 464.0, 1.087)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.051, 464.0, 1.087, 774.0, 1.12)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-6.095055E-8, 1.819084E-4, 1.015716914, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.051, 464.0, 1.087, 774.0, 1.12)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.087, 774.0, 1.12, 1292.0, 1.149)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-4.056852E-8, 1.3979912E-4, 1.0360991, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.087, 774.0, 1.12, 1292.0, 1.149)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.12, 1292.0, 1.149, 2154.0, 1.149)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-7.541751E-9, 2.5988874E-5, 1.128011548, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.12, 1292.0, 1.149, 2154.0, 1.149)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.149, 2154.0, 1.149, 3594.0, 1.124)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.149, 2154.0, 1.149, 3594.0, 1.124)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.149, 3594.0, 1.124)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.523804357, -0.0488339231, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.149, 3594.0, 1.124)
	}

	return jupiterRadii
}

func gasRadius300My260K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.613, 0.839};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.613, 215.0, 0.839)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.664633E-6, 0.0056085409, 0.0336863924, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.613, 215.0, 0.839)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.613, 215.0, 0.839, 318.0, 0.941)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.106578E-6, 0.0021130972, 0.482060664, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.613, 215.0, 0.839, 318.0, 0.941)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.839, 318.0, 0.941, 464.0, 1.009)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-5.544961E-7, 8.9936934E-4, 0.7110734098, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.839, 318.0, 0.941, 464.0, 1.009)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.941, 464.0, 1.009, 774.0, 1.075)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.545424E-7, 4.0422677E-4, 0.8547111493, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.941, 464.0, 1.009, 774.0, 1.075)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.009, 774.0, 1.075, 1292.0, 1.119)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-5.146447E-8, 1.9126768E-4, 0.9577899426, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.009, 774.0, 1.075, 1292.0, 1.119)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.075, 1292.0, 1.119, 2154.0, 1.131)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.147746E-8, 5.3472442E-5, 1.069072516, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.075, 1292.0, 1.119, 2154.0, 1.131)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.119, 2154.0, 1.131, 3594.0, 1.113)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.119, 2154.0, 1.131, 3594.0, 1.113)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.131, 3594.0, 1.113)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3584.0)
	} else {
		//jupiterRadii = ln_trend(1.400859137, -0.0351604246, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.131, 3594.0, 1.113)
	}

	return jupiterRadii
}

func gasRadius300Myr260K(planet *types.Planet) float64 {
	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	coreMassRadii0 := gasRadius300Myr260K0coreMass(planet)
	coreMassRadii10 := gasRadius300Myr260K10coreMass(planet)
	coreMassRadii25 := gasRadius300Myr260K25coreMass(planet)
	coreMassRadii50 := gasRadius300Myr260K50coreMass(planet)
	coreMassRadii100 := gasRadius300My260K100coreMass(planet)

	var jupiterRadii float64

	if coreEarthMasses <= 10.0 {
		/*jupiterRadii1 := gas_radius_300Myr_260K_0core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_260K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 0.0, 10.0);*/
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		/*jupiterRadii1 := gas_radius_300Myr_260K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_260K_25core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		/*jupiterRadii1 := gas_radius_300Myr_260K_25core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_260K_50core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		/*jupiterRadii1 := gas_radius_300Myr_260K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_260K_100core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}

	return jupiterRadii
}

func gasRadius300Myr78K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64
	if totalEarthMasses < 17.0 {
		//jupiterRadii = quad_trend(-0.0018802521, 0.0866113445, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 0.929, 28.0, 0.951)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-7.662835E-6, 0.0023448276, 0.8913524904, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 17.0, 0.929, 28.0, 0.951)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.929, 28.0, 0.951, 46.0, 0.983)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-1.192305E-5, 0.0026600834, 0.8858653354, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.929, 28.0, 0.951, 46.0, 0.983)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.951, 46.0, 0.983, 77.0, 1.02)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-2.7953E-6, 0.0015373703, 0.9181958205, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.951, 46.0, 0.983, 77.0, 1.02)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.983, 77.0, 1.02, 129.0, 1.07)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-3.934303E-6, 0.0017720049, 0.906882107, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.983, 77.0, 1.02, 129.0, 1.07)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.02, 129.0, 1.07, 215.0, 1.106)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.136091E-6, 8.0941988E-4, 0.9844905224, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.02, 129.0, 1.07, 215.0, 1.106)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.07, 215.0, 1.106, 318.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.961707E-7, 3.6174249E-4, 1.041915857, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.07, 215.0, 1.106, 318.0, 1.127)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.106, 318.0, 1.127, 464.0, 1.146)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.368313E-7, 2.3713903E-4, 1.065426713, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.106, 318.0, 1.127, 464.0, 1.146)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.127, 464.0, 1.146, 774.0, 1.167)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-7.715088E-8, 1.6325473E-4, 1.086860082, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.127, 464.0, 1.146, 774.0, 1.167)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.146, 774.0, 1.167, 1292.0, 1.169)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.372624E-8, 3.2219414E-5, 1.150285234, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.146, 774.0, 1.167, 1292.0, 1.169)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.167, 1292.0, 1.169, 2154.0, 1.156)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.292072E-9, -1.062873E-5, 1.184889125, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.167, 1292.0, 1.169, 2154.0, 1.156)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.169, 2154.0, 1.156, 3594.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.169, 2154.0, 1.156, 3594.0, 1.13)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.156, 3594.0, 1.13)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.545796532, -0.05078728, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.156, 3594.0, 1.13)
	}

	return jupiterRadii
}

func gasRadius300Myr78K10coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii10 := RadiusImproved(10.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 10.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii10
	} else if totalEarthMasses < 17.0 {
		/*double x[] = {10, 17, 28};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.565, 0.733};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.565, 28.0, 0.733)
	} else if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-3.08255E-4, 0.0291442006, 0.1586342738, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 10.0, massRadii10, 17.0, 0.565, 28.0, 0.733)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.565, 28.0, 0.733, 46.0, 0.847)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 17.0, 28.0)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-6.868554E-5, 0.0114160632, 0.4671996928, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 17.0, 0.565, 28.0, 0.733, 46.0, 0.847)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.733, 46.0, 0.847, 77.0, 0.939)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-1.791533E-5, 0.005171328, 0.6470277587, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.733, 46.0, 0.847, 77.0, 0.939)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.847, 77.0, 0.939, 129.0, 1.016)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-6.011641E-6, 0.0027191673, 0.765271405, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.847, 77.0, 0.939, 129.0, 1.016)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.939, 129.0, 1.016, 215.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-1.801498E-6, 0.0012708782, 0.8820354446, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.939, 129.0, 1.016, 215.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.016, 215.0, 1.072, 318.0, 1.104)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-5.050125E-7, 5.7985126E-4, 0.970676181, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.016, 215.0, 1.072, 318.0, 1.104)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.072, 318.0, 1.104, 464.0, 1.131)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-2.21624E-7, 3.5824147E-4, 1.012490718, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.072, 318.0, 1.104, 464.0, 1.131)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.104, 464.0, 1.131, 774.0, 1.157)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-8.730429E-8, 1.9195368E-4, 1.060729756, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.104, 464.0, 1.131, 774.0, 1.157)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.131, 774.0, 1.157, 1292.0, 1.163)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.76406E-8, 4.8028498E-5, 1.130394005, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.131, 774.0, 1.157, 1292.0, 1.163)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.157, 1292.0, 1.163, 2154.0, 1.152)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.998302E-9, -5.874874E-6, 1.17392603, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.157, 1292.0, 1.163, 2154.0, 1.152)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.163, 2154.0, 1.152, 3594.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.163, 2154.0, 1.152, 3594.0, 1.127)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.152, 3594.0, 1.127)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.526804357, -0.0488339231, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.152, 3594.0, 1.127)
	}
	return jupiterRadii
}

func gasRadius300Myr78K25coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii25 := RadiusImproved(25.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 25.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii25
	} else if totalEarthMasses < 28.0 {
		/*double x[] = {25, 28, 46};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.394, 0.664};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.394, 46.0, 0.664)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(-1.994733E-4, 0.029761027, -0.282921659, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 25.0, massRadii25, 28.0, 0.394, 46.0, 0.664)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.394, 46.0, 0.664, 77.0, 0.826)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(-3.608479E-5, 0.0096642351, 0.2958005919, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 0.394, 46.0, 0.664, 77.0, 0.826)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.664, 77.0, 0.826, 129.0, 0.942)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-9.255658E-6, 0.0041374349, 0.5622943144, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 0.664, 77.0, 0.826, 129.0, 0.942)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.826, 129.0, 0.942, 215.0, 1.024)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-2.527832E-6, 0.0018230625, 0.7488905841, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.826, 129.0, 0.942, 215.0, 1.024)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.942, 215.0, 1.024, 318.0, 1.073)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-9.7477288E-8, 4.2377276E-4, 0.9283829688, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.942, 215.0, 1.024, 318.0, 1.073)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.024, 318.0, 1.073, 464.0, 1.146)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-1.124788E-6, 0.001379584, 0.7480353141, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.024, 318.0, 1.073, 464.0, 1.146)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.073, 464.0, 1.146, 774.0, 1.142)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(4.1230371E-8, -6.394642E-5, 1.166794407, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.073, 464.0, 1.146, 774.0, 1.142)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.146, 774.0, 1.142, 1292.0, 1.153)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.127259E-8, 6.5184687E-5, 1.104290949, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.146, 774.0, 1.142, 1292.0, 1.153)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.142, 1292.0, 1.153, 2154.0, 1.146)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-3.109091E-9, 2.5932789E-6, 1.154839378, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.142, 1292.0, 1.153, 2154.0, 1.146)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.153, 2154.0, 1.146, 3594.0, 1.124)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.153, 2154.0, 1.146, 3594.0, 1.124)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.146, 3594.0, 1.124)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.475827834, -0.0429738523, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.146, 3594.0, 1.124)
	}
	return jupiterRadii
}

func gasRadius300Myr78K50coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii50 := RadiusImproved(50.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64
	if totalEarthMasses < 50.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii50
	} else if totalEarthMasses < 77.0 {
		/*double x[] = {50, 77, 129};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.635, 0.823};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.635, 129.0, 0.823)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(-1.541313E-5, 0.0067904903, 0.2035167224, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 50.0, massRadii50, 77.0, 0.635, 129.0, 0.823)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.635, 129.0, 0.823, 215.0, 0.951)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-4.330525E-6, 0.0029780726, 0.5108928957, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 0.635, 129.0, 0.823, 215.0, 0.951)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.823, 215.0, 0.951, 318.0, 1.02)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-1.259994E-6, 0.0013414798, 0.7208250777, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.823, 215.0, 0.951, 318.0, 1.02)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.951, 318.0, 1.02, 464.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-4.85778E-7, 7.0695223E-4, 0.8405511733, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.951, 318.0, 1.02, 464.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.02, 464.0, 1.072, 774.0, 1.119)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-1.411399E-7, 3.2634415E-4, 0.9509631802, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.02, 464.0, 1.072, 774.0, 1.119)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.072, 774.0, 1.119, 1292.0, 1.137)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-2.518046E-8, 8.6771865E-5, 1.066923586, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.072, 774.0, 1.119, 1292.0, 1.137)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.119, 1292.0, 1.137, 2154.0, 1.137)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-5.731731E-9, 1.97151545E-5, 1.121048776, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.119, 1292.0, 1.137, 2154.0, 1.137)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292, 1.137, 2154.0, 1.137, 3594.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292, 1.137, 2154.0, 1.137, 3594.0, 1.118)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.137, 3594.0, 1.118)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.421851312, -0.0371137816, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.137, 3594.0, 1.118)
	}
	return jupiterRadii
}

func gasRadius300Myr78K100coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	massRadii100 := RadiusImproved(100.0/constants.SunMassInEarthMasses, planet)

	var jupiterRadii float64

	if totalEarthMasses < 100.0 {
		//jupiterRadii = radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone(), the_planet) / KM_JUPITER_RADIUS;
		jupiterRadii = massRadii100
	} else if totalEarthMasses < 129.0 {
		/*double x[] = {100, 129, 215};
		  double y[] = {radius_improved(the_planet->getMass(), the_planet->getImf(), the_planet->getRmf(), the_planet->getCmf(), the_planet->getGasGiant(), the_planet->getOrbitZone()) / KM_JUPITER_RADIUS, 0.587, 0.810};
		  double coeff[3];
		  polynomialfit(3, 3, x, y, coeff);
		  jupiterRadii = quad_trend(coeff[2], coeff[1], coeff[0], totalEarthMasses);*/
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.587, 215.0, 0.81)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(-8.069112E-6, 0.0053687976, 0.02870319, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 100.0, massRadii100, 129.0, 0.587, 215.0, 0.81)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.587, 215.0, 0.81, 318.0, 0.92)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(-2.115925E-6, 0.0021957491, 0.4357225711, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 0.587, 215.0, 0.81, 318.0, 0.92)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.81, 318.0, 0.92, 464.0, 0.999)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(-6.702018E-7, 0.0010651937, 0.6490418913, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 0.81, 318.0, 0.92, 464.0, 0.999)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.92, 464.0, 0.999, 774.0, 1.072)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(-2.027975E-7, 4.8654714E-4, 0.8169036124, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 0.92, 464.0, 0.999, 774.0, 1.072)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.999, 774.0, 1.072, 1292.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-3.887424E-8, 1.4788175E-4, 0.9808281498, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 0.999, 774.0, 1.072, 1292.0, 1.107)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.072, 1292.0, 1.107, 2154.0, 1.119)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-9.66744E-9, 4.7235112E-5, 1.062109745, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.072, 1292.0, 1.107, 2154.0, 1.119)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.107, 2154.0, 1.119, 3594.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292, 1.107, 2154.0, 1.119, 3594.0, 1.107)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.119, 3594.0, 1.107)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.298906092, -0.0234402831, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.119, 3594.0, 1.107)
	}
	return jupiterRadii
}

func gasRadius300Myr78K(planet *types.Planet) float64 {

	coreMassRadii0 := gasRadius300Myr78K0coreMass(planet)
	coreMassRadii10 := gasRadius300Myr78K10coreMass(planet)
	coreMassRadii25 := gasRadius300Myr78K25coreMass(planet)
	coreMassRadii50 := gasRadius300Myr78K50coreMass(planet)
	coreMassRadii100 := gasRadius300Myr78K100coreMass(planet)

	coreEarthMasses := planet.DustMass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if coreEarthMasses <= 10.0 {
		/*jupiterRadii1 := gas_radius_300Myr_78K_0core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_78K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 0.0, 10.0);*/
		jupiterRadii = PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
	} else if coreEarthMasses <= 25.0 {
		/*jupiterRadii1 := gas_radius_300Myr_78K_10core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_78K_25core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 0.0, coreMassRadii0, 10.0, coreMassRadii10, 25.0, coreMassRadii25)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 10.0, 25.0)
	} else if coreEarthMasses <= 50.0 {
		/*jupiterRadii1 := gas_radius_300Myr_78K_25core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_78K_50core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 10.0, coreMassRadii10, 25.0, coreMassRadii25, 50.0, coreMassRadii50)
		jupiterRadii2 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 25.0, 50.0)
	} else if coreEarthMasses <= 100.0 {
		/*jupiterRadii1 := gas_radius_300Myr_78K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_78K_100core_mass(totalEarthMasses, the_planet);*/
		jupiterRadii1 := PlanetRadiusHelper(coreEarthMasses, 25.0, coreMassRadii25, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii2 := PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
		jupiterRadii = RangeAdjust(coreEarthMasses, jupiterRadii1, jupiterRadii2, 50.0, 100.0)
	} else {
		/*upper_fraction = coreEarthMasses / 100.0;
		  jupiterRadii = gas_radius_300Myr_78K_100core_mass(totalEarthMasses, the_planet) * pow1_4(upper_fraction);*/
		/*jupiterRadii1 := gas_radius_300Myr_78K_50core_mass(totalEarthMasses, the_planet);
		  jupiterRadii2 := gas_radius_300Myr_78K_100core_mass(totalEarthMasses, the_planet);
		  jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, jupiterRadii1, 100.0, jupiterRadii2);*/
		jupiterRadii = PlanetRadiusHelper2(coreEarthMasses, 50.0, coreMassRadii50, 100.0, coreMassRadii100)
	}
	return jupiterRadii
}

func gasRadius1Gyr1960K0coreMass(planet *types.Planet) float64 {
	totalEarthMasses := planet.Mass * constants.SunMassInEarthMasses

	var jupiterRadii float64

	if totalEarthMasses < 28.0 {
		//jupiterRadii = quad_trend(-0.0016532091, 0.1095041408, 0, totalEarthMasses);
		jupiterRadii = PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 28.0, 1.77, 46.0, 1.539)
	} else if totalEarthMasses < 46.0 {
		//jupiterRadii = quad_trend(1.6183893E-4, -0.024809414, 2.337781874, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 0.0, 0.0, 28.0, 1.77, 46.0, 1.539)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.77, 46.0, 1.539, 77.0, 1.387)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 28.0, 46.0)
	} else if totalEarthMasses < 77.0 {
		//jupiterRadii = quad_trend(4.1002721E-5, -0.0099465604, 1.909780023, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 28.0, 1.77, 46.0, 1.539, 77.0, 1.387)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.539, 77.0, 1.387, 129.0, 1.309)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 46.0, 77.0)
	} else if totalEarthMasses < 129.0 {
		//jupiterRadii = quad_trend(8.5102797E-6, -0.0032531176, 1.587032609, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 46.0, 1.539, 77.0, 1.387, 129.0, 1.309)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.387, 129.0, 1.309, 215.0, 1.281)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 77.0, 129.0)
	} else if totalEarthMasses < 215.0 {
		//jupiterRadii = quad_trend(5.4116627E-7, -5.117426E-4, 1.366009246, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 77.0, 1.387, 129.0, 1.309, 215.0, 1.281)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.309, 215.0, 1.281, 318.0, 1.258)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 129.0, 215.0)
	} else if totalEarthMasses < 318.0 {
		//jupiterRadii = quad_trend(6.2171815E-7, -5.546767E-4, 1.371516579, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 129.0, 1.309, 215.0, 1.281, 318.0, 1.258)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.281, 318.0, 1.258, 464.0, 1.248)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 215.0, 318.0)
	} else if totalEarthMasses < 464.0 {
		//jupiterRadii = quad_trend(5.8240497E-8, -1.140372E-4, 1.2888374324, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 215.0, 1.281, 318.0, 1.258, 464.0, 1.248)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.258, 464.0, 1.248, 774.0, 1.235)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 318.0, 464.0)
	} else if totalEarthMasses < 774.0 {
		//jupiterRadii = quad_trend(7.1630436E-8, -1.30614E-4, 1.293183133, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 318.0, 1.258, 464.0, 1.248, 774.0, 1.235)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.248, 774.0, 1.235, 1292.0, 1.244)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 464.0, 774.0)
	} else if totalEarthMasses < 1292.0 {
		//jupiterRadii = quad_trend(-1.595282E-8, 5.0333039E-5, 1.20559917, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 464.0, 1.248, 774.0, 1.235, 1292.0, 1.244)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.235, 1292.0, 1.244, 2154.0, 1.24)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 774.0, 1292.0)
	} else if totalEarthMasses < 2154.0 {
		//jupiterRadii = quad_trend(-1.035267E-8, 3.1034937E-5, 1.221184205, totalEarthMasses);
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 774.0, 1.235, 1292.0, 1.244, 2154.0, 1.24)
		jupiterRadii2 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.244, 2154.0, 1.24, 3594.0, 1.199)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 1292.0, 2154.0)
	} else if totalEarthMasses < 3594.0 {
		jupiterRadii1 := PlanetRadiusHelper(totalEarthMasses, 1292.0, 1.244, 2154.0, 1.24, 3594.0, 1.199)
		jupiterRadii2 := PlanetRadiusHelper2(totalEarthMasses, 2154.0, 1.24, 3594.0, 1.199)
		jupiterRadii = RangeAdjust(totalEarthMasses, jupiterRadii1, jupiterRadii2, 2154.0, 3594.0)
	} else {
		//jupiterRadii = ln_trend(1.854679146, -0.0800876339, totalEarthMasses);
		jupiterRadii = RangeAdjust(totalEarthMasses, 0.0, 0.0, 2154.0, 3594.0)
	}

	return jupiterRadii

}
