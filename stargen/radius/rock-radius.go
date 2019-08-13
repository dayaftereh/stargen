package radius

func rockRadius(mass, cmf float64) float64 {
	adjustForCarbon := true
	var radius float64

	if mass <= 0.00623 {
		adjustForCarbon = false
		radius1 := FractionRadius(mass, 0, 1, cmf)
		radius2 := PlanetRadiusHelper(mass, 0.00623, 0.2029, 0.008748, 0.227, 0.01227, 0.254)
		radius = RangeAdjust(mass, radius1, radius2, 0.0, 0.00623)
	} else if mass <= 0.008748 {
		//radius = quad_trend(-315.395589, 14.2950833, 0.1260830485, mass);
		radius = PlanetRadiusHelper(mass, 0.00623, 0.2029, 0.008748, 0.227, 0.01227, 0.254)
	} else if mass <= 0.01227 {
		//radius = quad_trend(-188.1341907, 11.62030323, 0.1397430274, mass);
		radius1 := PlanetRadiusHelper(mass, 0.00623, 0.2029, 0.008748, 0.227, 0.01227, 0.254)
		radius2 := PlanetRadiusHelper(mass, 0.008748, 0.227, 0.01227, 0.254, 0.01717, 0.2838)
		radius = RangeAdjust(mass, radius1, radius2, 0.008748, 0.01227)
	} else if mass <= 0.01717 {
		//radius = quad_trend(-103.5492715, 9.130123207, 0.1575630314, mass);
		radius1 := PlanetRadiusHelper(mass, 0.008748, 0.227, 0.01227, 0.254, 0.01717, 0.2838)
		radius2 := PlanetRadiusHelper(mass, 0.01227, 0.254, 0.01717, 0.2838, 0.02399, 0.317)
		radius = RangeAdjust(mass, radius1, radius2, 0.01227, 0.01717)
	} else if mass <= 0.02399 {
		//radius = quad_trend(-60.94197703, 7.376406965, 0.1751133296, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01227, 0.254, 0.01717, 0.2838, 0.02399, 0.317)
		radius2 := PlanetRadiusHelper(mass, 0.01717, 0.2838, 0.02399, 0.317, 0.03343, 0.3536)
		radius = RangeAdjust(mass, radius1, radius2, 0.01717, 0.02399)
	} else if mass <= 0.03343 {
		//radius = quad_trend(-34.72166711, 5.870836769, 0.1961416432, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01717, 0.2838, 0.02399, 0.317, 0.03343, 0.3536)
		radius2 := PlanetRadiusHelper(mass, 0.02399, 0.317, 0.03343, 0.3536, 0.04644, 0.3939)
		radius = RangeAdjust(mass, radius1, radius2, 0.02399, 0.03343)
	} else if mass <= 0.04644 {
		//radius = quad_trend(-20.17535427, 4.709022763, 0.2187246368, mass);
		radius1 := PlanetRadiusHelper(mass, 0.02399, 0.317, 0.03343, 0.3536, 0.04644, 0.3939)
		radius2 := PlanetRadiusHelper(mass, 0.03343, 0.3536, 0.04644, 0.3939, 0.0643, 0.4381)
		radius = RangeAdjust(mass, radius1, radius2, 0.03343, 0.04644)
	} else if mass <= 0.0643 {
		//radius = quad_trend(-11.55708954, 3.754636127, 0.2444595682, mass);
		radius1 := PlanetRadiusHelper(mass, 0.03343, 0.3536, 0.04644, 0.3939, 0.0643, 0.4381)
		radius2 := PlanetRadiusHelper(mass, 0.04644, 0.3939, 0.0643, 0.4381, 0.08866, 0.4865)
		radius = RangeAdjust(mass, radius1, radius2, 0.04644, 0.0643)
	} else if mass <= 0.08866 {
		//radius = quad_trend(-6.878989996, 3.039074021, 0.2711286558, mass);
		radius1 := PlanetRadiusHelper(mass, 0.04644, 0.3939, 0.0643, 0.4381, 0.08866, 0.4865)
		radius2 := PlanetRadiusHelper(mass, 0.0643, 0.4381, 0.08866, 0.4865, 0.1217, 0.5391)
		radius = RangeAdjust(mass, radius1, radius2, 0.0643, 0.08866)
	} else if mass <= 0.1217 {
		//radius = quad_trend(-4.009273679, 2.435400496, 0.302092671, mass);
		radius1 := PlanetRadiusHelper(mass, 0.0643, 0.4381, 0.08866, 0.4865, 0.1217, 0.5391)
		radius2 := PlanetRadiusHelper(mass, 0.08866, 0.4865, 0.1217, 0.5391, 0.1661, 0.596)
		radius = RangeAdjust(mass, radius1, radius2, 0.08866, 0.1217)
	} else if mass <= 0.1661 {
		//radius = quad_trend(-2.404094408, 1.973429902, 0.3345403587, mass);
		radius1 := PlanetRadiusHelper(mass, 0.08866, 0.4865, 0.1217, 0.5391, 0.1661, 0.596)
		radius2 := PlanetRadiusHelper(mass, 0.1217, 0.5391, 0.1661, 0.596, 0.2255, 0.6573)
		radius = RangeAdjust(mass, radius1, radius2, 0.1217, 0.1661)
	} else if mass <= 0.2255 {
		//radius = quad_trend(-1.446140999, 1.598295347, 0.3704210305, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1217, 0.5391, 0.1661, 0.596, 0.2255, 0.6573)
		radius2 := PlanetRadiusHelper(mass, 0.1661, 0.596, 0.2255, 0.6573, 0.3042, 0.7228)
		radius = RangeAdjust(mass, radius1, radius2, 0.1661, 0.2255)
	} else if mass <= 0.3042 {
		//radius = quad_trend(-0.8656081038, 1.290787073, 0.4102439036, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1661, 0.596, 0.2255, 0.6573, 0.3042, 0.7228)
		radius2 := PlanetRadiusHelper(mass, 0.2255, 0.6573, 0.3042, 0.7228, 0.4075, 0.7925)
		radius = RangeAdjust(mass, radius1, radius2, 0.2255, 0.3042)
	} else if mass <= 0.4075 {
		//radius = quad_trend(-0.5362568931, 1.056387816, 0.4510707737, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2255, 0.6573, 0.3042, 0.7228, 0.4075, 0.7925)
		radius2 := PlanetRadiusHelper(mass, 0.3042, 0.7228, 0.4075, 0.7925, 0.542, 0.8661)
		radius = RangeAdjust(mass, radius1, radius2, 0.3042, 0.4075)
	} else if mass <= 0.542 {
		//radius = quad_trend(-0.330761771, 0.8612701975, 0.4964574539, mass);
		radius1 := PlanetRadiusHelper(mass, 0.3042, 0.7228, 0.4075, 0.7925, 0.542, 0.8661)
		radius2 := PlanetRadiusHelper(mass, 0.4075, 0.7925, 0.542, 0.8661, 0.7143, 0.9429)
		radius = RangeAdjust(mass, radius1, radius2, 0.4075, 0.542)
	} else if mass <= 0.7143 {
		//radius = quad_trend(-0.2162384806, 0.7173945877, 0.5407952145, mass);
		radius1 := PlanetRadiusHelper(mass, 0.4075, 0.7925, 0.542, 0.8661, 0.7143, 0.9429)
		radius2 := PlanetRadiusHelper(mass, 0.542, 0.8661, 0.7143, 0.9429, 0.927, 1.02)
		radius = RangeAdjust(mass, radius1, radius2, 0.542, 0.7143)
	} else if mass <= 0.927 {
		//radius = quad_trend(-0.1354314862, 0.5847660678, 0.5943020587, mass);
		radius1 := PlanetRadiusHelper(mass, 0.542, 0.8661, 0.7143, 0.9429, 0.927, 1.02)
		radius2 := PlanetRadiusHelper(mass, 0.7143, 0.9429, 0.927, 1.02, 1.2, 1.101)
		radius = RangeAdjust(mass, radius1, radius2, 0.7143, 0.927)
	} else if mass <= 1.2 {
		//radius = quad_trend(-0.0814344419, 0.4699143547, 0.6543683708, mass);
		radius1 := PlanetRadiusHelper(mass, 0.7143, 0.9429, 0.927, 1.02, 1.2, 1.101)
		radius2 := PlanetRadiusHelper(mass, 0.927, 1.02, 1.2, 1.101, 1.545, 1.186)
		radius = RangeAdjust(mass, radius1, radius2, 0.927, 1.2)
	} else if mass <= 1.545 {
		//radius = quad_trend(-0.0570319452, 0.4029295012, 0.6996105997, mass);
		radius1 := PlanetRadiusHelper(mass, 0.927, 1.02, 1.2, 1.101, 1.545, 1.186)
		radius2 := PlanetRadiusHelper(mass, 1.2, 1.101, 1.545, 1.186, 1.981, 1.274)
		radius = RangeAdjust(mass, radius1, radius2, 1.2, 1.545)
	} else if mass <= 1.981 {
		//radius = quad_trend(-0.0352606639, 0.3261639633, 0.766244763, mass);
		radius1 := PlanetRadiusHelper(mass, 1.2, 1.101, 1.545, 1.186, 1.981, 1.274)
		radius2 := PlanetRadiusHelper(mass, 1.545, 1.186, 1.981, 1.274, 2.525, 1.365)
		radius = RangeAdjust(mass, radius1, radius2, 1.545, 1.981)
	} else if mass <= 2.525 {
		//radius = quad_trend(-0.024640974, 0.2783647176, 0.8193647176, mass);
		radius1 := PlanetRadiusHelper(mass, 1.545, 1.186, 1.981, 1.274, 2.525, 1.365)
		radius2 := PlanetRadiusHelper(mass, 1.981, 1.274, 2.525, 1.365, 3.203, 1.458)
		radius = RangeAdjust(mass, radius1, radius2, 1.981, 2.525)
	} else if mass <= 3.203 {
		//radius = quad_trend(-0.0150740628, 0.2235123732, 0.8967378292, mass);
		radius1 := PlanetRadiusHelper(mass, 1.981, 1.274, 2.525, 1.365, 3.203, 1.458)
		radius2 := PlanetRadiusHelper(mass, 2.525, 1.365, 3.203, 1.458, 4.043, 1.554)
		radius = RangeAdjust(mass, radius1, radius2, 2.525, 3.203)
	} else if mass <= 4.043 {
		//radius = quad_trend(-0.0098938279, 0.1859763911, 0.9638204674, mass);
		radius1 := PlanetRadiusHelper(mass, 2.525, 1.365, 3.203, 1.458, 4.043, 1.554)
		radius2 := PlanetRadiusHelper(mass, 3.203, 1.458, 4.043, 1.554, 5.077, 1.653)
		radius = RangeAdjust(mass, radius1, radius2, 3.203, 4.043)
	} else if mass <= 5.077 {
		//radius = quad_trend(-0.0075670613, 0.1647562803, 1.011580401, mass);
		radius1 := PlanetRadiusHelper(mass, 3.203, 1.458, 4.043, 1.554, 5.077, 1.653)
		radius2 := PlanetRadiusHelper(mass, 4.043, 1.554, 5.077, 1.653, 6.297, 1.749)
		radius = RangeAdjust(mass, radius1, radius2, 4.043, 5.077)
	} else if mass <= 6.297 {
		//radius = quad_trend(-0.0049514253, 0.1350060359, 1.095201943, mass);
		radius1 := PlanetRadiusHelper(mass, 4.043, 1.554, 5.077, 1.653, 6.297, 1.749)
		radius2 := PlanetRadiusHelper(mass, 5.077, 1.653, 6.297, 1.749, 7.714, 1.842)
		radius = RangeAdjust(mass, radius1, radius2, 5.077, 6.297)
	} else if mass <= 7.714 {
		//radius = quad_trend(-0.003587277, 0.1158929542, 1.161465525, mass);
		radius1 := PlanetRadiusHelper(mass, 5.077, 1.653, 6.297, 1.749, 7.714, 1.842)
		radius2 := PlanetRadiusHelper(mass, 6.297, 1.749, 7.714, 1.842, 9.423, 1.935)
		radius = RangeAdjust(mass, radius1, radius2, 6.297, 7.714)
	} else if mass <= 9.423 {
		//radius = quad_trend(-0.002262228, 0.0931855895, 1.257782041, mass);
		radius1 := PlanetRadiusHelper(mass, 6.297, 1.749, 7.714, 1.842, 9.423, 1.935)
		radius2 := PlanetRadiusHelper(mass, 7.714, 1.842, 9.423, 1.935, 11.47, 2.029)
		radius = RangeAdjust(mass, radius1, radius2, 7.714, 9.423)
	} else if mass <= 11.47 {
		//radius = quad_trend(-0.0017062124, 0.0815687551, 1.317877216, mass);
		radius1 := PlanetRadiusHelper(mass, 7.714, 1.842, 9.423, 1.935, 11.47, 2.029)
		radius2 := PlanetRadiusHelper(mass, 9.423, 1.935, 11.47, 2.029, 13.87, 2.121)
		radius = RangeAdjust(mass, radius1, radius2, 9.423, 11.47)
	} else if mass <= 13.87 {
		//radius = quad_trend(-0.0011721485, 0.0680355766, 1.402840849, mass);
		radius1 := PlanetRadiusHelper(mass, 9.423, 1.935, 11.47, 2.029, 13.87, 2.121)
		radius2 := PlanetRadiusHelper(mass, 11.47, 2.029, 13.87, 2.121, 16.73, 2.213)
		radius = RangeAdjust(mass, radius1, radius2, 11.47, 13.87)
	} else if mass <= 16.73 {
		//radius = quad_trend(-8.290313E-4, 0.057536189, 1.482459524, mass);
		radius1 := PlanetRadiusHelper(mass, 11.47, 2.029, 13.87, 2.121, 16.73, 2.213)
		radius2 := PlanetRadiusHelper(mass, 13.87, 2.121, 16.73, 2.213, 20.1, 2.304)
		radius = RangeAdjust(mass, radius1, radius2, 13.87, 16.73)
	} else if mass <= 20.1 {
		//radius = quad_trend(-5.903191E-4, 0.0487444197, 1.562731982, mass);
		radius1 := PlanetRadiusHelper(mass, 13.87, 2.121, 16.73, 2.213, 20.1, 2.304)
		radius2 := PlanetRadiusHelper(mass, 16.73, 2.213, 20.1, 2.304, 24.07, 2.394)
		radius = RangeAdjust(mass, radius1, radius2, 16.73, 20.1)
	} else if mass <= 24.07 {
		//radius = quad_trend(-3.920942E-4, 0.0399888266, 1.658634568, mass);
		radius1 := PlanetRadiusHelper(mass, 16.73, 2.213, 20.1, 2.304, 24.07, 2.394)
		radius2 := PlanetRadiusHelper(mass, 20.1, 2.304, 24.07, 2.394, 28.68, 2.483)
		radius = RangeAdjust(mass, radius1, radius2, 20.1, 24.07)
	} else if mass <= 28.68 {
		//radius = quad_trend(-3.135081E-4, 0.035843407, 1.712884759, mass);
		radius1 := PlanetRadiusHelper(mass, 20.1, 2.304, 24.07, 2.394, 28.68, 2.483)
		radius2 := PlanetRadiusHelper(mass, 24.07, 2.394, 28.68, 2.483, 33.99, 2.569)
		radius = RangeAdjust(mass, radius1, radius2, 24.07, 28.68)
	} else if mass <= 33.99 {
		//radius = quad_trend(-2.488584E-4, 0.0317918141, 1.775907375, mass);
		radius1 := PlanetRadiusHelper(mass, 24.07, 2.394, 28.68, 2.483, 33.99, 2.569)
		radius2 := PlanetRadiusHelper(mass, 28.68, 2.483, 33.99, 2.569, 40.05, 2.65)
		radius = RangeAdjust(mass, radius1, radius2, 28.68, 33.99)
	} else if mass <= 40.05 {
		//radius = quad_trend(-1.5098E-4, 0.0245448927, 1.909149277, mass);
		radius1 := PlanetRadiusHelper(mass, 28.68, 2.483, 33.99, 2.569, 40.05, 2.65)
		radius2 := PlanetRadiusHelper(mass, 33.99, 2.569, 40.05, 2.65, 46.88, 2.728)
		radius = RangeAdjust(mass, radius1, radius2, 33.99, 40.05)
	} else if mass <= 46.88 {
		//radius = quad_trend(-1.447629E-4, 0.024004441, 1.920822137, mass);
		radius1 := PlanetRadiusHelper(mass, 33.99, 2.569, 40.05, 2.65, 46.88, 2.728)
		radius2 := PlanetRadiusHelper(mass, 40.05, 2.65, 46.88, 2.728, 54.49, 2.799)
		radius = RangeAdjust(mass, radius1, radius2, 40.05, 46.88)
	} else if mass <= 54.49 {
		//radius = quad_trend(-9.444827E-5, 0.0189040505, 2.04935033, mass);
		radius1 := PlanetRadiusHelper(mass, 40.05, 2.65, 46.88, 2.728, 54.49, 2.799)
		radius2 := PlanetRadiusHelper(mass, 46.88, 2.728, 54.49, 2.799, 63.08, 2.866)
		radius = RangeAdjust(mass, radius1, radius2, 46.88, 54.49)
	} else if mass <= 63.08 {
		//radius = quad_trend(-7.602327E-5, 0.0167378233, 2.112681274, mass);
		radius1 := PlanetRadiusHelper(mass, 46.88, 2.728, 54.49, 2.799, 63.08, 2.866)
		radius2 := PlanetRadiusHelper(mass, 54.49, 2.799, 63.08, 2.866, 72.75, 2.928)
		radius = RangeAdjust(mass, radius1, radius2, 54.49, 63.08)
	} else if mass <= 72.75 {
		//radius = quad_trend(-5.173226E-5, 0.0134383755, 2.224154419, mass);
		radius1 := PlanetRadiusHelper(mass, 54.49, 2.799, 63.08, 2.866, 72.75, 2.928)
		radius2 := PlanetRadiusHelper(mass, 63.08, 2.866, 72.75, 2.928, 83.59, 2.986)
		radius = RangeAdjust(mass, radius1, radius2, 63.08, 72.75)
	} else if mass <= 83.59 {
		//radius = quad_trend(-4.216159E-5, 0.119420964, 2.282355333, mass);
		radius1 := PlanetRadiusHelper(mass, 63.08, 2.866, 72.75, 2.928, 83.59, 2.986)
		radius2 := PlanetRadiusHelper(mass, 72.75, 2.928, 83.59, 2.986, 95.68, 3.039)
		radius = RangeAdjust(mass, radius1, radius2, 72.75, 83.59)
	} else if mass <= 95.68 {
		//radius = quad_trend(-2.871507E-5, 0.0095315392, 2.389899117, mass);
		radius1 := PlanetRadiusHelper(mass, 72.75, 2.928, 83.59, 2.986, 95.68, 3.039)
		radius2 := PlanetRadiusHelper(mass, 83.59, 2.986, 95.68, 3.039, 109.1, 3.088)
		radius = RangeAdjust(mass, radius1, radius2, 83.59, 95.68)
	} else if mass <= 109.1 {
		//radius = quad_trend(-2.46556E-5, 0.0087002408, 2.432774665, mass);
		radius1 := PlanetRadiusHelper(mass, 83.59, 2.986, 95.68, 3.039, 109.1, 3.088)
		radius2 := PlanetRadiusHelper(mass, 95.68, 3.039, 109.1, 3.088, 124.0, 3.132)
		radius = RangeAdjust(mass, radius1, radius2, 95.68, 109.1)
	} else if mass <= 124.0 {
		//radius = quad_trend(-1.99273E-5, 0.007598074, 2.496241005, mass);
		radius1 := PlanetRadiusHelper(mass, 95.68, 3.039, 109.1, 3.088, 124.0, 3.132)
		radius2 := PlanetRadiusHelper(mass, 109.1, 3.088, 124.0, 3.132, 140.3, 3.17)
		radius = RangeAdjust(mass, radius1, radius2, 109.1, 124.0)
	} else if mass <= 140.3 {
		//radius = quad_trend(-1.262711E-5, 0.0056686323, 2.623243968, mass);
		radius1 := PlanetRadiusHelper(mass, 109.1, 3.088, 124.0, 3.132, 140.3, 3.17)
		radius2 := PlanetRadiusHelper(mass, 124.0, 3.132, 140.3, 3.17, 158.2, 3.204)
		radius = RangeAdjust(mass, radius1, radius2, 124.0, 140.3)
	} else if mass <= 158.2 {
		//radius = quad_trend(-1.255653E-5, 0.0056475658, 2.624810399, mass);
		radius1 := PlanetRadiusHelper(mass, 124.0, 3.132, 140.3, 3.17, 158.2, 3.204)
		radius2 := PlanetRadiusHelper(mass, 140.3, 3.17, 158.2, 3.204, 177.8, 3.232)
		radius = RangeAdjust(mass, radius1, radius2, 140.3, 158.2)
	} else if mass <= 177.8 {
		//radius = quad_trend(-8.629392E-6, 0.0043280472, 2.735272806, mass);
		radius1 := PlanetRadiusHelper(mass, 140.3, 3.17, 158.2, 3.204, 177.8, 3.232)
		radius2 := PlanetRadiusHelper(mass, 158.2, 3.204, 177.8, 3.232, 199.2, 3.255)
		radius = RangeAdjust(mass, radius1, radius2, 158.2, 177.8)
	} else if mass <= 199.2 {
		//radius = quad_trend(-5.801246E-6, 0.0032618362, 2.835439392, mass);
		radius1 := PlanetRadiusHelper(mass, 158.2, 3.204, 177.8, 3.232, 199.2, 3.255)
		radius2 := PlanetRadiusHelper(mass, 177.8, 3.232, 199.2, 3.255, 222.5, 3.274)
		radius = RangeAdjust(mass, radius1, radius2, 177.8, 199.2)
	} else if mass <= 222.5 {
		//radius = quad_trend(-5.492344E-6, 0.0031315723, 2.849130541, mass);
		radius1 := PlanetRadiusHelper(mass, 177.8, 3.232, 199.2, 3.255, 222.5, 3.274)
		radius2 := PlanetRadiusHelper(mass, 199.2, 3.255, 222.5, 3.274, 248.1, 3.288)
		radius = RangeAdjust(mass, radius1, radius2, 199.2, 222.5)
	} else if mass <= 248.1 {
		radius1 := PlanetRadiusHelper(mass, 199.2, 3.255, 222.5, 3.274, 248.1, 3.288)
		radius2 := PlanetRadiusHelper2(mass, 222.5, 3.274, 248.1, 3.288)
		radius = RangeAdjust(mass, radius1, radius2, 222.5, 248.1)
	} else {
		//radius = ln_trend(2.57918226, 0.128552657, mass);
		radius = PlanetRadiusHelper2(mass, 222.5, 3.274, 248.1, 3.288)
	}

	if adjustForCarbon {
		rmf := 1.0
		carbonFraction := rmf * cmf
		growFactor := (0.05 * carbonFraction) + 1.0
		radius *= growFactor
	}

	return radius
}

func halfRockHalfWaterRadius(mass, cmf float64) float64 {
	adjustForCarbon := true

	var radius float64
	if mass <= 0.008278 {
		adjustForCarbon = false
		radius1 := FractionRadius(mass, 0.5, 0.5, cmf)
		radius2 := PlanetRadiusHelper(mass, 0.008278, 0.2963, 0.01156, 0.3286, 0.01615, 0.3647)
		radius = RangeAdjust(mass, radius1, radius2, 0.0, 0.008278)
	} else if mass <= 0.01156 {
		//radius = quad_trend(-251.097088, 14.82282406, 0.1908031617, mass);
		radius = PlanetRadiusHelper(mass, 0.008278, 0.2963, 0.01156, 0.3286, 0.01615, 0.3647)
	} else if mass <= 0.01615 {
		//radius = quad_trend(-144.1013419, 11.85797193, 0.2107786256, mass);
		radius1 := PlanetRadiusHelper(mass, 0.008278, 0.2963, 0.01156, 0.3286, 0.01615, 0.3647)
		radius2 := PlanetRadiusHelper(mass, 0.01156, 0.3286, 0.01615, 0.3647, 0.02255, 0.4049)
		radius = RangeAdjust(mass, radius1, radius2, 0.01156, 0.01615)
	} else if mass <= 0.02255 {
		//radius = quad_trend(-86.45685997, 9.627130481, 0.2317717371, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01156, 0.3286, 0.01615, 0.3647, 0.02255, 0.4049)
		radius2 := PlanetRadiusHelper(mass, 0.01615, 0.3647, 0.02255, 0.4049, 0.0313, 0.4484)
		radius = RangeAdjust(mass, radius1, radius2, 0.01615, 0.02255)
	} else if mass <= 0.0313 {
		//radius = quad_trend(-52.75830191, 7.812463129, 0.2555566849, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01615, 0.3647, 0.02255, 0.4049, 0.0313, 0.4484)
		radius2 := PlanetRadiusHelper(mass, 0.02255, 0.4049, 0.0313, 0.4484, 0.04314, 0.4944)
		radius = RangeAdjust(mass, radius1, radius2, 0.02255, 0.0313)
	} else if mass <= 0.04314 {
		//radius = quad_trend(-27.9087717, 5.9626641, 0.2891105582, mass);
		radius1 := PlanetRadiusHelper(mass, 0.02255, 0.4049, 0.0313, 0.4484, 0.04314, 0.4944)
		radius2 := PlanetRadiusHelper(mass, 0.02255, 0.4484, 0.04314, 0.4944, 0.05943, 0.5449)
		radius = RangeAdjust(mass, radius1, radius2, 0.0313, 0.04314)
	} else if mass <= 0.05943 {
		//radius = quad_trend(-15.88199788, 4.72907791, 0.3199449236, mass);
		radius1 := PlanetRadiusHelper(mass, 0.02255, 0.4484, 0.04314, 0.4944, 0.05943, 0.5449)
		radius2 := PlanetRadiusHelper(mass, 0.04314, 0.4944, 0.05943, 0.5449, 0.0817, 0.6003)
		radius = RangeAdjust(mass, radius1, radius2, 0.04314, 0.05943)
	} else if mass <= 0.0817 {
		//radius = quad_trend(-9.401792072, 3.814526464, 0.3514091157, mass);
		radius1 := PlanetRadiusHelper(mass, 0.04314, 0.4944, 0.05943, 0.5449, 0.0817, 0.6003)
		radius2 := PlanetRadiusHelper(mass, 0.05943, 0.5449, 0.0817, 0.6003, 0.112, 0.6607)
		radius = RangeAdjust(mass, radius1, radius2, 0.05943, 0.0817)
	} else if mass <= 0.112 {
		//radius = quad_trend(-5.4572037, 3.050459697, 0.3875036772, mass);
		radius1 := PlanetRadiusHelper(mass, 0.05943, 0.5449, 0.0817, 0.6003, 0.112, 0.6607)
		radius2 := PlanetRadiusHelper(mass, 0.0817, 0.6003, 0.112, 0.6607, 0.1528, 0.7262)
		radius = RangeAdjust(mass, radius1, radius2, 0.0817, 0.112)
	} else if mass <= 0.1528 {
		//radius = quad_trend(-3.312524816, 2.482548728, 0.4242068537, mass);
		radius1 := PlanetRadiusHelper(mass, 0.0817, 0.6003, 0.112, 0.6607, 0.1528, 0.7262)
		radius2 := PlanetRadiusHelper(mass, 0.112, 0.6607, 0.1528, 0.7262, 0.2074, 0.7966)
		radius = RangeAdjust(mass, radius1, radius2, 0.112, 0.1528)
	} else if mass <= 0.2074 {
		//radius = quad_trend(-1.983760111, 2.003927681, 0.466316364, mass);
		radius1 := PlanetRadiusHelper(mass, 0.112, 0.6607, 0.1528, 0.7262, 0.2074, 0.7966)
		radius2 := PlanetRadiusHelper(mass, 0.1528, 0.7262, 0.2074, 0.7966, 0.2799, 0.8718)
		radius = RangeAdjust(mass, radius1, radius2, 0.1528, 0.2074)
	} else if mass <= 0.2799 {
		//radius = quad_trend(-1.206466699, 1.625152602, 0.5114392259, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1528, 0.7262, 0.2074, 0.7966, 0.2799, 0.8718)
		radius2 := PlanetRadiusHelper(mass, 0.2074, 0.7966, 0.2799, 0.8718, 0.3754, 0.9515)
		radius = RangeAdjust(mass, radius1, radius2, 0.2074, 0.2799)
	} else if mass <= 0.3754 {
		//radius = quad_trend(-0.7427608407, 1.321286153, 0.5601628686, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2074, 0.7966, 0.2799, 0.8718, 0.3754, 0.9515)
		radius2 := PlanetRadiusHelper(mass, 0.2799, 0.8718, 0.3754, 0.9515, 0.4998, 1.035)
		radius = RangeAdjust(mass, radius1, radius2, 0.2799, 0.3754)
	} else if mass <= 0.4998 {
		//radius = quad_trend(-0.4356758148, 1.052525338, 0.617779672, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2799, 0.8718, 0.3754, 0.9515, 0.4998, 1.035)
		radius2 := PlanetRadiusHelper(mass, 0.3754, 0.9515, 0.4998, 1.035, 0.6607, 1.123)
		radius = RangeAdjust(mass, radius1, radius2, 0.3754, 0.4998)
	} else if mass <= 0.6607 {
		//radius = quad_trend(-0.279489115, 0.8712706729, 0.6693553098, mass);
		radius1 := PlanetRadiusHelper(mass, 0.3754, 0.9515, 0.4998, 1.035, 0.6607, 1.123)
		radius2 := PlanetRadiusHelper(mass, 0.4998, 1.035, 0.6607, 1.123, 0.8653, 1.214)
		radius = RangeAdjust(mass, radius1, radius2, 0.4998, 0.6607)
	} else if mass <= 0.8653 {
		//radius = quad_trend(-0.1823992238, 0.7231114989, 0.7248619608, mass);
		radius1 := PlanetRadiusHelper(mass, 0.4998, 1.035, 0.6607, 1.123, 0.8653, 1.214)
		radius2 := PlanetRadiusHelper(mass, 0.6607, 1.123, 0.8653, 1.214, 1.117, 1.305)
		radius = RangeAdjust(mass, radius1, radius2, 0.6607, 0.8653)
	} else if mass <= 1.117 {
		//radius = quad_trend(-0.1131126774, 0.5857647781, 0.7918301862, mass);
		radius1 := PlanetRadiusHelper(mass, 0.6607, 1.123, 0.8653, 1.214, 1.117, 1.305)
		radius2 := PlanetRadiusHelper(mass, 0.8653, 1.214, 1.117, 1.305, 1.437, 1.4)
		radius = RangeAdjust(mass, radius1, radius2, 0.8653, 1.117)
	} else if mass <= 1.437 {
		//radius = quad_trend(-0.0723180077, 0.4815751916, 0.8573108937, mass);
		radius1 := PlanetRadiusHelper(mass, 0.8653, 1.214, 1.117, 1.305, 1.437, 1.4)
		radius2 := PlanetRadiusHelper(mass, 1.117, 1.305, 1.437, 1.4, 1.842, 1.499)
		radius = RangeAdjust(mass, radius1, radius2, 1.117, 1.437)
	} else if mass <= 1.842 {
		//radius = quad_trend(-0.0460469153, 0.3954322797, 0.9268492667, mass);
		radius1 := PlanetRadiusHelper(mass, 1.117, 1.305, 1.437, 1.4, 1.842, 1.499)
		radius2 := PlanetRadiusHelper(mass, 1.437, 1.4, 1.842, 1.499, 2.351, 1.602)
		radius = RangeAdjust(mass, radius1, radius2, 1.437, 1.842)
	} else if mass <= 2.351 {
		//radius = quad_trend(-0.0313721992, 0.3339011952, 0.990398741, mass);
		radius1 := PlanetRadiusHelper(mass, 1.437, 1.4, 1.842, 1.499, 2.351, 1.602)
		radius2 := PlanetRadiusHelper(mass, 1.842, 1.499, 2.351, 1.602, 2.988, 1.708)
		radius = RangeAdjust(mass, radius1, radius2, 1.842, 2.351)
	} else if mass <= 2.988 {
		//radius = quad_trend(-0.0192555176, 0.269210232, 1.075515861, mass);
		radius1 := PlanetRadiusHelper(mass, 1.842, 1.499, 2.351, 1.602, 2.988, 1.708)
		radius2 := PlanetRadiusHelper(mass, 2.351, 1.602, 2.988, 1.708, 3.78, 1.818)
		radius = RangeAdjust(mass, radius1, radius2, 2.351, 2.988)
	} else if mass <= 3.78 {
		//radius = quad_trend(-0.0123380727, 0.2223929652, 1.15364591, mass);
		radius1 := PlanetRadiusHelper(mass, 2.351, 1.602, 2.988, 1.708, 3.78, 1.818)
		radius2 := PlanetRadiusHelper(mass, 2.988, 1.708, 3.78, 1.818, 4.763, 1.933)
		radius = RangeAdjust(mass, radius1, radius2, 2.988, 3.78)
	} else if mass <= 4.763 {
		//radius = quad_trend(-0.0092219964, 0.1957723254, 1.209748184, mass);
		radius1 := PlanetRadiusHelper(mass, 2.988, 1.708, 3.78, 1.818, 4.763, 1.933)
		radius2 := PlanetRadiusHelper(mass, 3.78, 1.818, 4.763, 1.933, 5.972, 2.05)
		radius = RangeAdjust(mass, radius1, radius2, 3.78, 4.763)
	} else if mass <= 5.972 {
		//radius = quad_trend(-0.006005431, 0.1612424952, 1.301242218, mass);
		radius1 := PlanetRadiusHelper(mass, 3.78, 1.818, 4.763, 1.933, 5.972, 2.05)
		radius2 := PlanetRadiusHelper(mass, 4.763, 1.933, 5.972, 2.05, 7.392, 2.165)
		radius = RangeAdjust(mass, radius1, radius2, 4.763, 5.972)
	} else if mass <= 7.392 {
		//radius = quad_trend(-0.0046758803, 0.1434743798, 1.359935265, mass);
		radius1 := PlanetRadiusHelper(mass, 4.763, 1.933, 5.972, 2.05, 7.392, 2.165)
		radius2 := PlanetRadiusHelper(mass, 5.972, 2.05, 7.392, 2.165, 9.043, 2.275)
		radius = RangeAdjust(mass, radius1, radius2, 5.972, 7.392)
	} else if mass <= 9.043 {
		//radius = quad_trend(-0.0029585423, 0.1152499298, 1.474732193, mass);
		radius1 := PlanetRadiusHelper(mass, 5.972, 2.05, 7.392, 2.165, 9.043, 2.275)
		radius2 := PlanetRadiusHelper(mass, 7.392, 2.165, 9.043, 2.275, 11.03, 2.386)
		radius = RangeAdjust(mass, radius1, radius2, 7.392, 9.043)
	} else if mass <= 11.03 {
		//radius = quad_trend(-0.0019751495, 0.0955102856, 1.572820013, mass);
		radius1 := PlanetRadiusHelper(mass, 7.392, 2.165, 9.043, 2.275, 11.03, 2.386)
		radius2 := PlanetRadiusHelper(mass, 9.043, 2.275, 11.03, 2.386, 13.4, 2.498)
		radius = RangeAdjust(mass, radius1, radius2, 9.043, 11.03)
	} else if mass <= 13.4 {
		//radius = quad_trend(-0.0014930172, 0.0837317945, 1.644080125, mass);
		radius1 := PlanetRadiusHelper(mass, 9.043, 2.275, 11.03, 2.386, 13.4, 2.498)
		radius2 := PlanetRadiusHelper(mass, 11.03, 2.386, 13.4, 2.498, 16.18, 2.608)
		radius = RangeAdjust(mass, radius1, radius2, 11.03, 13.4)
	} else if mass <= 16.18 {
		//radius = quad_trend(-0.0010753359, 0.0713767814, 1.734638445, mass);
		radius1 := PlanetRadiusHelper(mass, 11.03, 2.386, 13.4, 2.498, 16.18, 2.608)
		radius2 := PlanetRadiusHelper(mass, 13.4, 2.498, 16.18, 2.608, 19.48, 2.717)
		radius = RangeAdjust(mass, radius1, radius2, 13.4, 16.18)
	} else if mass <= 19.48 {
		//radius = quad_trend(-7.235726E-4, 0.058832903, 1.845509445, mass);
		radius1 := PlanetRadiusHelper(mass, 13.4, 2.498, 16.18, 2.608, 19.48, 2.717)
		radius2 := PlanetRadiusHelper(mass, 16.18, 2.608, 19.48, 2.717, 23.36, 2.825)
		radius = RangeAdjust(mass, radius1, radius2, 16.18, 19.48)
	} else if mass <= 23.36 {
		//radius = quad_trend(-5.028683E-4, 0.049377928, 1.945941584, mass);
		radius1 := PlanetRadiusHelper(mass, 16.18, 2.608, 19.48, 2.717, 23.36, 2.825)
		radius2 := PlanetRadiusHelper(mass, 19.48, 2.717, 23.36, 2.825, 27.94, 2.933)
		radius = RangeAdjust(mass, radius1, radius2, 19.48, 23.36)
	} else if mass <= 27.94 {
		//radius = quad_trend(-4.006219E-4, 0.0441326902, 2.012675568, mass);
		radius1 := PlanetRadiusHelper(mass, 19.48, 2.717, 23.36, 2.825, 27.94, 2.933)
		radius2 := PlanetRadiusHelper(mass, 23.36, 2.825, 27.94, 2.933, 33.24, 3.037)
		radius = RangeAdjust(mass, radius1, radius2, 23.36, 27.94)
	} else if mass <= 33.24 {
		//radius = quad_trend(-2.667319E-4, 0.0359413008, 2.137022627, mass);
		radius1 := PlanetRadiusHelper(mass, 23.36, 2.825, 27.94, 2.933, 33.24, 3.037)
		radius2 := PlanetRadiusHelper(mass, 27.94, 2.933, 33.24, 3.037, 39.33, 3.138)
		radius = RangeAdjust(mass, radius1, radius2, 27.94, 33.24)
	} else if mass <= 39.33 {
		//radius = quad_trend(-2.098119E-4, 0.0318106148, 2.211435834, mass);
		radius1 := PlanetRadiusHelper(mass, 27.94, 2.933, 33.24, 3.037, 39.33, 3.138)
		radius2 := PlanetRadiusHelper(mass, 33.24, 3.037, 39.33, 3.138, 46.26, 3.234)
		radius = RangeAdjust(mass, radius1, radius2, 33.24, 39.33)
	} else if mass <= 46.26 {
		//radius = quad_trend(-1.493274E-4, 0.0266337427, 2.321481765, mass);
		radius1 := PlanetRadiusHelper(mass, 33.24, 3.037, 39.33, 3.138, 46.26, 3.234)
		radius2 := PlanetRadiusHelper(mass, 39.33, 3.138, 46.26, 3.234, 54.07, 3.325)
		radius = RangeAdjust(mass, radius1, radius2, 39.33, 49.26)
	} else if mass <= 54.07 {
		//radius = quad_trend(-1.209301E-4, 0.0237846466, 2.392511186, mass);
		radius1 := PlanetRadiusHelper(mass, 39.33, 3.138, 46.26, 3.234, 54.07, 3.325)
		radius2 := PlanetRadiusHelper(mass, 46.26, 3.234, 54.07, 3.325, 62.77, 3.409)
		radius = RangeAdjust(mass, radius1, radius2, 49.26, 54.07)
	} else if mass <= 62.77 {
		//radius = quad_trend(-8.655674E-5, 0.019768462, 2.509173507, mass);
		radius1 := PlanetRadiusHelper(mass, 46.26, 3.234, 54.07, 3.325, 62.77, 3.409)
		radius2 := PlanetRadiusHelper(mass, 54.07, 3.325, 62.77, 3.409, 72.58, 3.488)
		radius = RangeAdjust(mass, radius1, radius2, 54.07, 62.77)
	} else if mass <= 72.58 {
		//radius = quad_trend(-6.040906E-5, 0.0162293732, 2.628298337, mass);
		radius1 := PlanetRadiusHelper(mass, 54.07, 3.325, 62.77, 3.409, 72.58, 3.488)
		radius2 := PlanetRadiusHelper(mass, 62.77, 3.409, 72.58, 3.488, 83.62, 3.563)
		radius = RangeAdjust(mass, radius1, radius2, 62.77, 72.58)
	} else if mass <= 83.62 {
		//radius = quad_trend(-5.504085E-5, 0.0153908584, 2.66087877, mass);
		radius1 := PlanetRadiusHelper(mass, 62.77, 3.409, 72.58, 3.488, 83.62, 3.563)
		radius2 := PlanetRadiusHelper(mass, 72.58, 3.488, 83.62, 3.563, 95.97, 3.631)
		radius = RangeAdjust(mass, radius1, radius2, 72.58, 83.62)
	} else if mass <= 95.97 {
		//radius = quad_trend(-3.355431E-5, 0.0115320909, 2.833308489, mass);
		radius1 := PlanetRadiusHelper(mass, 72.58, 3.488, 83.62, 3.563, 95.97, 3.631)
		radius2 := PlanetRadiusHelper(mass, 83.62, 3.563, 95.97, 3.631, 109.8, 3.695)
		radius = RangeAdjust(mass, radius1, radius2, 83.62, 95.97)
	} else if mass <= 109.8 {
		//radius = quad_trend(-2.64817E-5, 0.0100767608, 2.907836114, mass);
		radius1 := PlanetRadiusHelper(mass, 83.62, 3.563, 95.97, 3.631, 109.8, 3.695)
		radius2 := PlanetRadiusHelper(mass, 95.97, 3.631, 109.8, 3.695, 125.1, 3.754)
		radius = RangeAdjust(mass, radius1, radius2, 95.97, 109.8)
	} else if mass <= 125.1 {
		//radius = quad_trend(-2.23638E-5, 0.0091094666, 2.964399476, mass);
		radius1 := PlanetRadiusHelper(mass, 95.97, 3.631, 109.8, 3.695, 125.1, 3.754)
		radius2 := PlanetRadiusHelper(mass, 109.8, 3.695, 125.1, 3.754, 142.0, 3.807)
		radius = RangeAdjust(mass, radius1, radius2, 109.8, 125.1)
	} else if mass <= 142.0 {
		//radius = quad_trend(-1.716093E-5, 0.0077197786, 3.056824394, mass);
		radius1 := PlanetRadiusHelper(mass, 109.8, 3.695, 125.1, 3.754, 142.0, 3.807)
		radius2 := PlanetRadiusHelper(mass, 125.1, 3.754, 142.0, 3.807, 160.6, 3.854)
		radius = RangeAdjust(mass, radius1, radius2, 125.1, 142.0)
	} else if mass <= 160.6 {
		//radius = quad_trend(-1.325841E-5, 0.0065388751, 3.145822216, mass);
		radius1 := PlanetRadiusHelper(mass, 125.1, 3.754, 142.0, 3.807, 160.6, 3.854)
		radius2 := PlanetRadiusHelper(mass, 142.0, 3.807, 160.6, 3.854, 181.0, 3.895)
		radius = RangeAdjust(mass, radius1, radius2, 142.0, 160.6)
	} else if mass <= 181.0 {
		//radius = quad_trend(-1.031141E-5, 0.0055321815, 3.231487246, mass);
		radius1 := PlanetRadiusHelper(mass, 142.0, 3.807, 160.6, 3.854, 181.0, 3.895)
		radius2 := PlanetRadiusHelper(mass, 160.6, 3.854, 181.0, 3.895, 203.3, 3.93)
		radius = RangeAdjust(mass, radius1, radius2, 160.6, 181.0)
	} else if mass <= 203.3 {
		//radius = quad_trend(-8.158076E-6, 0.0047046552, 3.310724123, mass);
		radius1 := PlanetRadiusHelper(mass, 160.6, 3.854, 181.0, 3.895, 203.3, 3.93)
		radius2 := PlanetRadiusHelper(mass, 181.0, 3.895, 203.3, 3.93, 227.7, 3.959)
		radius = RangeAdjust(mass, radius1, radius2, 181.0, 203.3)
	} else if mass <= 227.7 {
		//radius = quad_trend(-5.557254E-6, 0.0035837009, 3.431119842, mass);
		radius1 := PlanetRadiusHelper(mass, 181.0, 3.895, 203.3, 3.93, 227.7, 3.959)
		radius2 := PlanetRadiusHelper(mass, 203.3, 3.93, 227.7, 3.959, 254.2, 3.983)
		radius = RangeAdjust(mass, radius1, radius2, 203.3, 227.7)
	} else if mass <= 254.2 {
		//radius = quad_trend(-4.545683E-6, 0.0030962202, 3.489671491, mass);
		radius1 := PlanetRadiusHelper(mass, 203.3, 3.93, 227.7, 3.959, 254.2, 3.983)
		radius2 := PlanetRadiusHelper(mass, 227.7, 3.959, 254.2, 3.983, 283.3, 4.002)
		radius = RangeAdjust(mass, radius1, radius2, 227.7, 254.2)
	} else if mass <= 283.3 {
		radius1 := PlanetRadiusHelper(mass, 227.7, 3.959, 254.2, 3.983, 283.3, 4.002)
		radius2 := PlanetRadiusHelper2(mass, 254.2, 3.983, 283.3, 4.002)
		radius = RangeAdjust(mass, radius1, radius2, 254.2, 283.3)
	} else {
		//radius = ln_trend(3.012162161, 0.1753009325, mass);
		radius = PlanetRadiusHelper2(mass, 254.2, 3.983, 283.3, 4.002)
	}

	if adjustForCarbon {
		rmf := 0.5
		carbonFraction := rmf * cmf
		growFactor := (0.05 * carbonFraction) + 1.0 // not very scientific but guested by eye-balling an image on wikipedia.
		radius *= growFactor
	}
	return radius
}
