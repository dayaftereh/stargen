package radius

func ironRadius(mass float64) float64 {
	var radius float64
	if mass <= 0.001496 {
		radius1 := FractionRadius(mass, 0, 0, 0)
		radius2 := PlanetRadiusHelper(mass, 0.001496, 0.09947, 0.002096, 0.1112, 0.002931, 0.1243)
		radius = RangeAdjust(mass, radius1, radius2, 0.0, 0.001496)
	} else if mass <= 0.002096 {
		//radius = quad_trend(-2681.77957240149, 29.1753892167919, 0.0618254642138286, mass);
		radius = PlanetRadiusHelper(mass, 0.001496, 0.09947, 0.002096, 0.1112, 0.002931, 0.1243)
	} else if mass <= 0.002931 {
		//radius = quad_trend(-1688.63294694608, 24.2372061422203, 0.0678174361405534, mass);
		radius1 := PlanetRadiusHelper(mass, 0.001496, 0.09947, 0.002096, 0.1112, 0.002931, 0.1243)
		radius2 := PlanetRadiusHelper(mass, 0.002096, 0.1112, 0.002931, 0.1243, 0.00409, 0.1387)
		radius = RangeAdjust(mass, radius1, radius2, 0.002096, 0.002931)
	} else if mass <= 0.00409 {
		//radius = quad_trend(-915.413258520028, 18.8629573194402, 0.0768735211649926, mass);
		radius1 := PlanetRadiusHelper(mass, 0.002096, 0.1112, 0.002931, 0.1243, 0.00409, 0.1387)
		radius2 := PlanetRadiusHelper(mass, 0.002931, 0.1243, 0.00409, 0.1387, 0.005694, 0.1546)
		radius = RangeAdjust(mass, radius1, radius2, 0.002931, 0.00409)
	} else if mass <= 0.005694 {
		//radius = quad_trend(-512.025546524638, 14.9254112879076, 0.086217430425244, mass);
		radius1 := PlanetRadiusHelper(mass, 0.002931, 0.1243, 0.00409, 0.1387, 0.005694, 0.1546)
		radius2 := PlanetRadiusHelper(mass, 0.00409, 0.1387, 0.005694, 0.1546, 0.007904, 0.1722)
		radius = RangeAdjust(mass, radius1, radius2, 0.00409, 0.005694)
	} else if mass <= 0.007904 {
		//radius = quad_trend(-292.3316061, 11.93892609, 0.0960976238, mass);
		radius1 := PlanetRadiusHelper(mass, 0.00409, 0.1387, 0.005694, 0.1546, 0.007904, 0.1722)
		radius2 := PlanetRadiusHelper(mass, 0.005694, 0.1546, 0.007904, 0.1722, 0.01094, 0.1915)
		radius = RangeAdjust(mass, radius1, radius2, 0.005694, 0.007904)
	} else if mass <= 0.01094 {
		//radius = quad_trend(-192.8510361, 10.06829345, 0.104668233, mass);
		radius1 := PlanetRadiusHelper(mass, 0.005694, 0.1546, 0.007904, 0.1722, 0.01094, 0.1915)
		radius2 := PlanetRadiusHelper(mass, 0.007904, 0.1722, 0.01094, 0.1915, 0.01507, 0.2126)
		radius = RangeAdjust(mass, radius1, radius2, 0.007904, 0.01094)
	} else if mass <= 0.01507 {
		//radius = quad_trend(-99.45862138, 7.642892436, 0.1200091513, mass);
		radius1 := PlanetRadiusHelper(mass, 0.007904, 0.1722, 0.01094, 0.1915, 0.01507, 0.2126)
		radius2 := PlanetRadiusHelper(mass, 0.01094, 0.1915, 0.01507, 0.2126, 0.0207, 0.2356)
		radius = RangeAdjust(mass, radius1, radius2, 0.01094, 0.01507)
	} else if mass <= 0.0207 {
		//radius = quad_trend(-59.86761779, 6.226722237, 0.1323595252, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01094, 0.1915, 0.01507, 0.2126, 0.0207, 0.2356)
		radius2 := PlanetRadiusHelper(mass, 0.01507, 0.2126, 0.0207, 0.2356, 0.02829, 0.2606)
		radius = RangeAdjust(mass, radius1, radius2, 0.01507, 0.0207)
	} else if mass <= 0.02829 {
		//radius = quad_trend(-36.47985573, 5.080955774, 0.1460554689, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01507, 0.2126, 0.0207, 0.2356, 0.02829, 0.2606)
		radius2 := PlanetRadiusHelper(mass, 0.0207, 0.2356, 0.02829, 0.2606, 0.0385, 0.2876)
		radius = RangeAdjust(mass, radius1, radius2, 0.0207, 0.02829)
	} else if mass <= 0.0385 {
		//radius = quad_trend(-21.31356831, 4.067999437, 0.1625740583, mass);
		radius1 := PlanetRadiusHelper(mass, 0.0207, 0.2356, 0.02829, 0.2606, 0.0385, 0.2876)
		radius2 := PlanetRadiusHelper(mass, 0.02829, 0.2606, 0.0385, 0.2876, 0.05212, 0.3167)
		radius = RangeAdjust(mass, radius1, radius2, 0.02829, 0.0385)
	} else if mass <= 0.05212 {
		//radius = quad_trend(-12.81381824, 3.297752085, 0.1796298268, mass);
		radius1 := PlanetRadiusHelper(mass, 0.02829, 0.2606, 0.0385, 0.2876, 0.05212, 0.3167)
		radius2 := PlanetRadiusHelper(mass, 0.0385, 0.2876, 0.05212, 0.3167, 0.07021, 0.348)
		radius = RangeAdjust(mass, radius1, radius2, 0.0385, 0.05212)
	} else if mass <= 0.07021 {
		//radius = quad_trend(-7.888269424, 2.695209699, 0.1976541102, mass);
		radius1 := PlanetRadiusHelper(mass, 0.0385, 0.2876, 0.05212, 0.3167, 0.07021, 0.348)
		radius2 := PlanetRadiusHelper(mass, 0.05212, 0.3167, 0.07021, 0.348, 0.09408, 0.3814)
		radius = RangeAdjust(mass, radius1, radius2, 0.05212, 0.07021)
	} else if mass <= 0.09408 {
		//radius = quad_trend(-4.757963763, 2.180931782, 0.218330896, mass);
		radius1 := PlanetRadiusHelper(mass, 0.05212, 0.3167, 0.07021, 0.348, 0.09408, 0.3814)
		radius2 := PlanetRadiusHelper(mass, 0.07021, 0.348, 0.09408, 0.3814, 0.1254, 0.417)
		radius = RangeAdjust(mass, radius1, radius2, 0.07021, 0.09408)
	} else if mass <= 0.1254 {
		//radius = quad_trend(-2.941685354, 1.782294997, 0.2397586803, mass);
		radius1 := PlanetRadiusHelper(mass, 0.07021, 0.348, 0.09408, 0.3814, 0.1254, 0.417)
		radius2 := PlanetRadiusHelper(mass, 0.09408, 0.3814, 0.1254, 0.417, 0.1663, 0.4548)
		radius = RangeAdjust(mass, radius1, radius2, 0.09408, 0.1254)
	} else if mass <= 0.1663 {
		//radius = quad_trend(-1.784894626, 1.444859141, 0.2638824172, mass);
		radius1 := PlanetRadiusHelper(mass, 0.09408, 0.3814, 0.1254, 0.417, 0.1663, 0.4548)
		radius2 := PlanetRadiusHelper(mass, 0.1254, 0.417, 0.1663, 0.4548, 0.2193, 0.4949)
		radius = RangeAdjust(mass, radius1, radius2, 0.1254, 0.1663)
	} else if mass <= 0.2193 {
		//radius = quad_trend(-1.162328645, 1.204797699, 0.2865871433, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1254, 0.417, 0.1663, 0.4548, 0.2193, 0.4949)
		radius2 := PlanetRadiusHelper(mass, 0.1663, 0.4548, 0.2193, 0.4949, 0.2877, 0.537)
		radius = RangeAdjust(mass, radius1, radius2, 0.1663, 0.2193)
	} else if mass <= 0.2877 {
		//radius = quad_trend(-0.699716184, 0.9702531813, 0.3157745709, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1663, 0.4548, 0.2193, 0.4949, 0.2877, 0.537)
		radius2 := PlanetRadiusHelper(mass, 0.2193, 0.4949, 0.2877, 0.537, 0.3754, 0.5814)
		radius = RangeAdjust(mass, radius1, radius2, 0.2193, 0.2877)
	} else if mass <= 0.3754 {
		//radius = quad_trend(-0.4577736374, 0.8098210786, 0.3419049902, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2193, 0.4949, 0.2877, 0.537, 0.3754, 0.5814)
		radius2 := PlanetRadiusHelper(mass, 0.2877, 0.537, 0.3754, 0.5814, 0.4875, 0.6279)
		radius = RangeAdjust(mass, radius1, radius2, 0.2877, 0.3754)
	} else if mass <= 0.4875 {
		//radius = quad_trend(-0.2880355042, 0.6633540435, 0.3729683416, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2877, 0.537, 0.3754, 0.5814, 0.4875, 0.6279)
		radius2 := PlanetRadiusHelper(mass, 0.3754, 0.5814, 0.4875, 0.6279, 0.6298, 0.6765)
		radius = RangeAdjust(mass, radius1, radius2, 0.2877, 0.3754)
	} else if mass <= 0.6298 {
		//radius = quad_trend(-0.1883400931, 0.5519643698, 0.4035775744, mass);
		radius1 := PlanetRadiusHelper(mass, 0.3754, 0.5814, 0.4875, 0.6279, 0.6298, 0.6765)
		radius2 := PlanetRadiusHelper(mass, 0.4875, 0.6279, 0.6298, 0.6765, 0.8096, 0.727)
		radius = RangeAdjust(mass, radius1, radius2, 0.3754, 0.6298)
	} else if mass <= 0.8096 {
		//radius = quad_trend(-0.1194866451, 0.4528567076, 0.4386849891, mass);
		radius1 := PlanetRadiusHelper(mass, 0.4875, 0.6279, 0.6298, 0.6765, 0.8096, 0.727)
		radius2 := PlanetRadiusHelper(mass, 0.6298, 0.6765, 0.8096, 0.727, 1.036, 0.7796)
		radius = RangeAdjust(mass, radius1, radius2, 0.6298, 0.8096)
	} else if mass <= 1.036 {
		//radius = quad_trend(-0.0787318553, 0.3776396675, 0.4728678898, mass);
		radius1 := PlanetRadiusHelper(mass, 0.6298, 0.6765, 0.8096, 0.727, 1.036, 0.7796)
		radius2 := PlanetRadiusHelper(mass, 0.8096, 0.727, 1.036, 0.7796, 1.319, 0.834)
		radius = RangeAdjust(mass, radius1, radius2, 0.8096, 1.036)
	} else if mass <= 1.319 {
		//radius = quad_trend(-0.0512867047, 0.313006338, 0.5103712488, mass);
		radius1 := PlanetRadiusHelper(mass, 0.8096, 0.727, 1.036, 0.7796, 1.319, 0.834)
		radius2 := PlanetRadiusHelper(mass, 1.036, 0.7796, 1.319, 0.834, 1.671, 0.8902)
		radius = RangeAdjust(mass, radius1, radius2, 1.036, 1.319)
	} else if mass <= 1.671 {
		//radius = quad_trend(-0.0344294192, 0.2626030543, 0.5475255322, mass);
		radius1 := PlanetRadiusHelper(mass, 1.036, 0.7796, 1.319, 0.834, 1.671, 0.8902)
		radius2 := PlanetRadiusHelper(mass, 1.319, 0.834, 1.671, 0.8902, 2.108, 0.9481)
		radius = RangeAdjust(mass, radius1, radius2, 1.319, 1.671)
	} else if mass <= 2.108 {
		//radius = quad_trend(-0.0239715508, 0.2230827695, 0.584363039, mass);
		radius1 := PlanetRadiusHelper(mass, 1.319, 0.834, 1.671, 0.8902, 2.108, 0.9481)
		radius2 := PlanetRadiusHelper(mass, 1.671, 0.8902, 2.108, 0.9481, 2.648, 1.007)
		radius = RangeAdjust(mass, radius1, radius2, 1.671, 2.108)
	} else if mass <= 2.648 {
		//radius = quad_trend(-0.0140840757, 0.176057938, 0.6395547667, mass);
		radius1 := PlanetRadiusHelper(mass, 1.671, 0.8902, 2.108, 0.9481, 2.648, 1.007)
		radius2 := PlanetRadiusHelper(mass, 2.108, 0.9481, 2.648, 1.007, 3.31, 1.068)
		radius = RangeAdjust(mass, radius1, radius2, 2.108, 2.648)
	} else if mass <= 3.31 {
		//radius = quad_trend(-0.0105419379, 0.154953881, 0.6706011795, mass);
		radius1 := PlanetRadiusHelper(mass, 2.108, 0.9481, 2.648, 1.007, 3.31, 1.068)
		radius2 := PlanetRadiusHelper(mass, 2.648, 1.007, 3.31, 1.068, 4.119, 1.13)
		radius = RangeAdjust(mass, radius1, radius2, 2.648, 3.31)
	} else if mass <= 4.119 {
		//radius = quad_trend(-0.0070348211, 0.1288995104, 0.718416824, mass);
		radius1 := PlanetRadiusHelper(mass, 2.648, 1.007, 3.31, 1.068, 4.119, 1.13)
		radius2 := PlanetRadiusHelper(mass, 3.31, 1.068, 4.119, 1.13, 5.103, 1.193)
		radius = RangeAdjust(mass, radius1, radius2, 3.31, 4.119)
	} else if mass <= 5.103 {
		//radius = quad_trend(-0.0047115353, 0.1074741683, 0.7672505662, mass);
		radius1 := PlanetRadiusHelper(mass, 3.31, 1.068, 4.119, 1.13, 5.103, 1.193)
		radius2 := PlanetRadiusHelper(mass, 4.119, 1.13, 5.103, 1.193, 6.293, 1.257)
		radius = RangeAdjust(mass, radius1, radius2, 4.119, 5.103)
	} else if mass <= 6.293 {
		//radius = quad_trend(-0.003487465, 0.0935246637, 0.8065593536, mass);
		radius1 := PlanetRadiusHelper(mass, 4.119, 1.13, 5.103, 1.193, 6.293, 1.257)
		radius2 := PlanetRadiusHelper(mass, 5.103, 1.193, 6.293, 1.257, 7.727, 1.321)
		radius = RangeAdjust(mass, radius1, radius2, 5.103, 6.293)
	} else if mass <= 7.727 {
		//radius = quad_trend(-0.0021560003, 0.0748575287, 0.8713031702, mass);
		radius1 := PlanetRadiusHelper(mass, 5.103, 1.193, 6.293, 1.257, 7.727, 1.321)
		radius2 := PlanetRadiusHelper(mass, 6.293, 1.257, 7.727, 1.321, 9.445, 1.386)
		radius = RangeAdjust(mass, radius1, radius2, 6.293, 7.727)
	} else if mass <= 9.445 {
		//radius = quad_trend(-0.00160772, 0.0654424596, 0.9113174962, mass);
		radius1 := PlanetRadiusHelper(mass, 6.293, 1.257, 7.727, 1.321, 9.445, 1.386)
		radius2 := PlanetRadiusHelper(mass, 7.727, 1.321, 9.445, 1.386, 11.49, 1.451)
		radius = RangeAdjust(mass, radius1, radius2, 7.727, 9.445)
	} else if mass <= 11.49 {
		//radius = quad_trend(-0.0012172944, 0.0572688997, 0.9536876732, mass);
		radius1 := PlanetRadiusHelper(mass, 7.727, 1.321, 9.445, 1.386, 11.49, 1.451)
		radius2 := PlanetRadiusHelper(mass, 9.445, 1.386, 11.49, 1.451, 13.92, 1.515)
		radius = RangeAdjust(mass, radius1, radius2, 9.445, 11.49)
	} else if mass <= 13.92 {
		//radius = quad_trend(-7.485494E-4, 0.0453580881, 1.028659131, mass);
		radius1 := PlanetRadiusHelper(mass, 9.445, 1.386, 11.49, 1.451, 13.92, 1.515)
		radius2 := PlanetRadiusHelper(mass, 11.49, 1.451, 13.92, 1.515, 16.78, 1.579)
		radius = RangeAdjust(mass, radius1, radius2, 11.49, 13.92)
	} else if mass <= 16.78 {
		//radius = quad_trend(-5.83219E-4, 0.0402824467, 1.067276595, mass);
		radius1 := PlanetRadiusHelper(mass, 11.49, 1.451, 13.92, 1.515, 16.78, 1.579)
		radius2 := PlanetRadiusHelper(mass, 13.92, 1.515, 16.78, 1.579, 20.14, 1.642)
		radius = RangeAdjust(mass, radius1, radius2, 13.92, 16.78)
	} else if mass <= 20.14 {
		//radius = quad_trend(-3.979673E-4, 0.0334429539, 1.129882258, mass);
		radius1 := PlanetRadiusHelper(mass, 13.92, 1.515, 16.78, 1.579, 20.14, 1.642)
		radius2 := PlanetRadiusHelper(mass, 16.78, 1.579, 20.14, 1.642, 24.05, 1.704)
		radius = RangeAdjust(mass, radius1, radius2, 16.78, 20.14)
	} else if mass <= 24.05 {
		//radius = quad_trend(-2.896199E-4, 0.0286550795, 1.182362194, mass);
		radius1 := PlanetRadiusHelper(mass, 16.78, 1.579, 20.14, 1.642, 24.05, 1.704)
		radius2 := PlanetRadiusHelper(mass, 20.14, 1.642, 24.05, 1.704, 28.6, 1.765)
		radius = RangeAdjust(mass, radius1, radius2, 20.14, 24.05)
	} else if mass <= 28.6 {
		//radius = quad_trend(-2.251678E-4, 0.0252616764, 1.226694282, mass);
		radius1 := PlanetRadiusHelper(mass, 20.14, 1.642, 24.05, 1.704, 28.6, 1.765)
		radius2 := PlanetRadiusHelper(mass, 24.05, 1.704, 28.6, 1.765, 33.87, 1.824)
		radius = RangeAdjust(mass, radius1, radius2, 24.05, 28.6)
	} else if mass <= 33.87 {
		//radius = quad_trend(-1.591712E-4, 0.0211388691, 1.290623996, mass);
		radius1 := PlanetRadiusHelper(mass, 24.05, 1.704, 28.6, 1.765, 33.87, 1.824)
		radius2 := PlanetRadiusHelper(mass, 28.6, 1.765, 33.87, 1.824, 39.94, 1.881)
		radius = RangeAdjust(mass, radius1, radius2, 28.6, 33.87)
	} else if mass <= 39.94 {
		//radius = quad_trend(-1.04791E-4, 0.171250664, 1.364187783, mass);
		radius1 := PlanetRadiusHelper(mass, 28.6, 1.765, 33.87, 1.824, 39.94, 1.881)
		radius2 := PlanetRadiusHelper(mass, 33.87, 1.824, 39.94, 1.881, 46.92, 1.937)
		radius = RangeAdjust(mass, radius1, radius2, 33.87, 39.94)
	} else if mass <= 46.92 {
		//radius = quad_trend(-9.380878E-5, 0.0161711529, 1.38476825, mass);
		radius1 := PlanetRadiusHelper(mass, 33.87, 1.824, 39.94, 1.881, 46.92, 1.937)
		radius2 := PlanetRadiusHelper(mass, 39.94, 1.881, 46.92, 1.937, 54.93, 1.99)
		radius = RangeAdjust(mass, radius1, radius2, 39.94, 46.92)
	} else if mass <= 54.93 {
		//radius = quad_trend(-5.440961E-5, 0.0121583483, 1.486312324, mass);
		radius1 := PlanetRadiusHelper(mass, 39.94, 1.881, 46.92, 1.937, 54.93, 1.99)
		radius2 := PlanetRadiusHelper(mass, 46.92, 1.937, 54.93, 1.99, 64.08, 2.042)
		radius = RangeAdjust(mass, radius1, radius2, 46.92, 54.93)
	} else if mass <= 64.08 {
		//radius = quad_trend(-5.031019E-5, 0.0116704759, 1.500741944, mass);
		radius1 := PlanetRadiusHelper(mass, 46.92, 1.937, 54.93, 1.99, 64.08, 2.042)
		radius2 := PlanetRadiusHelper(mass, 54.93, 1.99, 64.08, 2.042, 74.51, 2.091)
		radius = RangeAdjust(mass, radius1, radius2, 54.93, 64.08)
	} else if mass <= 74.51 {
		//radius = quad_trend(-3.297829E-5, 0.0092684477, 1.583494853, mass);
		radius1 := PlanetRadiusHelper(mass, 54.93, 1.99, 64.08, 2.042, 74.51, 2.091)
		radius2 := PlanetRadiusHelper(mass, 64.08, 2.042, 74.51, 2.091, 86.37, 2.138)
		radius = RangeAdjust(mass, radius1, radius2, 64.08, 74.51)
	} else if mass <= 86.37 {
		//radius = quad_trend(-2.420693E-5, 0.0078573106, 1.639942343, mass);
		radius1 := PlanetRadiusHelper(mass, 64.08, 2.042, 74.51, 2.091, 86.37, 2.138)
		radius2 := PlanetRadiusHelper(mass, 74.51, 2.091, 86.37, 2.138, 99.8, 2.183)
		radius = RangeAdjust(mass, radius1, radius2, 74.51, 86.37)
	} else if mass <= 99.8 {
		//radius = quad_trend(-1.822424E-5, 0.0067435142, 1.691511445, mass);
		radius1 := PlanetRadiusHelper(mass, 74.51, 2.091, 86.37, 2.138, 99.8, 2.183)
		radius2 := PlanetRadiusHelper(mass, 86.37, 2.138, 99.8, 2.183, 115.0, 2.226)
		radius = RangeAdjust(mass, radius1, radius2, 86.37, 99.8)
	} else if mass <= 115.0 {
		//radius = quad_trend(-1.516304E-5, 0.0060859676, 1.726644882, mass);
		radius1 := PlanetRadiusHelper(mass, 86.37, 2.138, 99.8, 2.183, 115.0, 2.226)
		radius2 := PlanetRadiusHelper(mass, 99.8, 2.183, 115.0, 2.226, 132.1, 2.266)
		radius = RangeAdjust(mass, radius1, radius2, 99.8, 115.0)
	} else if mass <= 132.1 {
		radius1 := PlanetRadiusHelper(mass, 99.8, 2.183, 115.0, 2.226, 132.1, 2.266)
		radius2 := PlanetRadiusHelper2(mass, 115.0, 2.226, 132.1, 2.266)
		radius = RangeAdjust(mass, radius1, radius2, 115.0, 132.1)
	} else {
		//radius = ln_trend(0.8568787518, 0.2885439056, mass);
		radius = PlanetRadiusHelper2(mass, 115.0, 2.226, 132.1, 2.266)
	}

	return radius
}

func halfRockHalfIronRadius(mass, cmf float64) float64 {
	adjustForCarbon := true

	var radius float64
	if mass <= 0.00177 {
		adjustForCarbon = false
		radius1 := FractionRadius(mass, 0, 0.5, cmf)
		radius2 := PlanetRadiusHelper(mass, 0.00177, 0.121, 0.002481, 0.1354, 0.003472, 0.1513)
		radius = RangeAdjust(mass, radius1, radius2, 0.0, 0.00177)
	} else if mass <= 0.002481 {
		radius = PlanetRadiusHelper(mass, 0.00177, 0.121, 0.002481, 0.1354, 0.003472, 0.1513)
	} else if mass <= 0.003472 {
		//radius = quad_trend(-1373.39582492175, 24.2607963117367, 0.0836627150671445, mass);
		radius1 := PlanetRadiusHelper(mass, 0.00177, 0.121, 0.002481, 0.1354, 0.003472, 0.1513)
		radius2 := PlanetRadiusHelper(mass, 0.002481, 0.1354, 0.003472, 0.1513, 0.004848, 0.169)
		radius = RangeAdjust(mass, radius1, radius2, 0.002481, 0.003472)
	} else if mass <= 0.004848 {
		//radius = quad_trend(-799.3217849, 19.51372934, 0.0931839832, mass);
		radius1 := PlanetRadiusHelper(mass, 0.002481, 0.1354, 0.003472, 0.1513, 0.004848, 0.169)
		radius2 := PlanetRadiusHelper(mass, 0.003472, 0.1513, 0.004848, 0.169, 0.006752, 0.1885)
		radius = RangeAdjust(mass, radius1, radius2, 0.003472, 0.004848)
	} else if mass <= 0.006752 {
		//radius = quad_trend(-446.2529913, 15.41813134, 0.1047412297, mass);
		radius1 := PlanetRadiusHelper(mass, 0.003472, 0.1513, 0.004848, 0.169, 0.006752, 0.1885)
		radius2 := PlanetRadiusHelper(mass, 0.004848, 0.169, 0.006752, 0.1885, 0.00938, 0.2101)
		radius = RangeAdjust(mass, radius1, radius2, 0.004848, 0.006752)
	} else if mass <= 0.00938 {
		//radius = quad_trend(-260.7214329, 12.42513624, 0.1164916409, mass);
		radius1 := PlanetRadiusHelper(mass, 0.004848, 0.169, 0.006752, 0.1885, 0.00938, 0.2101)
		radius2 := PlanetRadiusHelper(mass, 0.006752, 0.1885, 0.00938, 0.2101, 0.01299, 0.2339)
		radius = RangeAdjust(mass, radius1, radius2, 0.006752, 0.00938)
	} else if mass <= 0.01299 {
		//radius = quad_trend(-152.0702736, 9.994609805, 0.1297303718, mass);
		radius1 := PlanetRadiusHelper(mass, 0.006752, 0.1885, 0.00938, 0.2101, 0.01299, 0.2339)
		radius2 := PlanetRadiusHelper(mass, 0.00938, 0.2101, 0.01299, 0.2339, 0.01792, 0.26)
		radius = RangeAdjust(mass, radius1, radius2, 0.00938, 0.01299)
	} else if mass <= 0.01792 {
		//radius = quad_trend(-89.11289838, 8.048597336, 0.14438564, mass);
		radius1 := PlanetRadiusHelper(mass, 0.00938, 0.2101, 0.01299, 0.2339, 0.01792, 0.26)
		radius2 := PlanetRadiusHelper(mass, 0.01299, 0.2339, 0.01792, 0.26, 0.02464, 0.2886)
		radius = RangeAdjust(mass, radius1, radius2, 0.01299, 0.01792)
	} else if mass <= 0.02464 {
		//radius = quad_trend(-52.58495246, 6.493967957, 0.1605145107, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01299, 0.2339, 0.01792, 0.26, 0.02464, 0.2886)
		radius2 := PlanetRadiusHelper(mass, 0.01792, 0.26, 0.02464, 0.2886, 0.03372, 0.03372)
		radius = RangeAdjust(mass, radius1, radius2, 0.01792, 0.02464)
	} else if mass <= 0.03372 {
		//radius = quad_trend(-31.03774295, 5.236472811, 0.1784172424, mass);
		radius1 := PlanetRadiusHelper(mass, 0.01792, 0.26, 0.02464, 0.2886, 0.03372, 0.03372)
		radius2 := PlanetRadiusHelper(mass, 0.02464, 0.2886, 0.03372, 0.03372, 0.04595, 0.3535)
		radius = RangeAdjust(mass, radius1, radius2, 0.02464, 0.03372)
	} else if mass <= 0.04595 {
		//radius = quad_trend(-18.41664973, 4.230950314, 0.1979727934, mass);
		radius1 := PlanetRadiusHelper(mass, 0.02464, 0.2886, 0.03372, 0.03372, 0.04595, 0.3535)
		radius2 := PlanetRadiusHelper(mass, 0.03372, 0.03372, 0.04595, 0.3535, 0.06231, 0.3901)
		radius = RangeAdjust(mass, radius1, radius2, 0.03372, 0.04595)
	} else if mass <= 0.06231 {
		//radius = quad_trend(-11.08681604, 3.437422519, 0.2189591664, mass);
		radius1 := PlanetRadiusHelper(mass, 0.03372, 0.03372, 0.04595, 0.3535, 0.06231, 0.3901)
		radius2 := PlanetRadiusHelper(mass, 0.04595, 0.3535, 0.06231, 0.3901, 0.08408, 0.4296)
		radius = RangeAdjust(mass, radius1, radius2, 0.04595, 0.06231)
	} else if mass <= 0.08408 {
		//radius = quad_trend(-6.715816383, 2.797551879, 0.241858942, mass);
		radius1 := PlanetRadiusHelper(mass, 0.04595, 0.3535, 0.06231, 0.3901, 0.08408, 0.4296)
		radius2 := PlanetRadiusHelper(mass, 0.06231, 0.3901, 0.08408, 0.4296, 0.1129, 0.4721)
		radius = RangeAdjust(mass, radius1, radius2, 0.06231, 0.08408)
	} else if mass <= 0.1129 {
		//radius = quad_trend(-4.069306668, 2.276242395, 0.2669812848, mass);
		radius1 := PlanetRadiusHelper(mass, 0.06231, 0.3901, 0.08408, 0.4296, 0.1129, 0.4721)
		radius2 := PlanetRadiusHelper(mass, 0.08408, 0.4296, 0.1129, 0.4721, 0.1508, 0.5177)
		radius = RangeAdjust(mass, radius1, radius2, 0.08408, 0.1129)
	} else if mass <= 0.1508 {
		//radius = quad_trend(-2.532586328, 1.871009242, 0.2931444403, mass);
		radius1 := PlanetRadiusHelper(mass, 0.08408, 0.4296, 0.1129, 0.4721, 0.1508, 0.5177)
		radius2 := PlanetRadiusHelper(mass, 0.1129, 0.4721, 0.1508, 0.5177, 0.2003, 0.5663)
		radius = RangeAdjust(mass, radius1, radius2, 0.1129, 0.1508)
	} else if mass <= 0.2003 {
		//radius = quad_trend(-1.57175725, 1.533662152, 0.3221665132, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1129, 0.4721, 0.1508, 0.5177, 0.2003, 0.5663)
		radius2 := PlanetRadiusHelper(mass, 0.1508, 0.5177, 0.2003, 0.5663, 0.2647, 0.618)
		radius = RangeAdjust(mass, radius1, radius2, 0.1508, 0.2003)
	} else if mass <= 0.2647 {
		//radius = quad_trend(-0.9812585362, 1.25908025, 0.3534744066, mass);
		radius1 := PlanetRadiusHelper(mass, 0.1508, 0.5177, 0.2003, 0.5663, 0.2647, 0.618)
		radius2 := PlanetRadiusHelper(mass, 0.2003, 0.5663, 0.2647, 0.618, 0.348, 0.6738)
		radius = RangeAdjust(mass, radius1, radius2, 0.2003, 0.2647)
	} else if mass <= 0.348 {
		//radius = quad_trend(-0.6134611137, 1.03373077, 0.3873542869, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2003, 0.5663, 0.2647, 0.618, 0.348, 0.6738)
		radius2 := PlanetRadiusHelper(mass, 0.2647, 0.618, 0.348, 0.6738, 0.455, 0.7307)
		radius = RangeAdjust(mass, radius1, radius2, 0.2647, 0.348)
	} else if mass <= 0.455 {
		//radius = quad_trend(-0.3947160298, 0.8580784672, 0.4219903835, mass);
		radius1 := PlanetRadiusHelper(mass, 0.2647, 0.618, 0.348, 0.6738, 0.455, 0.7307)
		radius2 := PlanetRadiusHelper(mass, 0.348, 0.6738, 0.455, 0.7307, 0.5919, 0.7916)
		radius = RangeAdjust(mass, radius1, radius2, 0.348, 0.455)
	} else if mass <= 0.5919 {
		//radius = quad_trend(-0.251475037, 0.7081194719, 0.4605672598, mass);
		radius1 := PlanetRadiusHelper(mass, 0.348, 0.6738, 0.455, 0.7307, 0.5919, 0.7916)
		radius2 := PlanetRadiusHelper(mass, 0.455, 0.7307, 0.5919, 0.7916, 0.7659, 0.8554)
		radius = RangeAdjust(mass, radius1, radius2, 0.455, 0.5919)
	} else if mass <= 0.7659 {
		//radius = quad_trend(-0.1614376955, 0.5858667696, 0.501384447, mass);
		radius1 := PlanetRadiusHelper(mass, 0.455, 0.7307, 0.5919, 0.7916, 0.7659, 0.8554)
		radius2 := PlanetRadiusHelper(mass, 0.5919, 0.7916, 0.7659, 0.8554, 0.986, 0.9221)
		radius = RangeAdjust(mass, radius1, radius2, 0.5919, 0.7659)
	} else if mass <= 0.986 {
		//radius = quad_trend(-0.1082398009, 0.4926693781, 0.5415582947, mass);
		radius1 := PlanetRadiusHelper(mass, 0.5919, 0.7916, 0.7659, 0.8554, 0.986, 0.9221)
		radius2 := PlanetRadiusHelper(mass, 0.7659, 0.8554, 0.986, 0.9221, 1.261, 0.9907)
		radius = RangeAdjust(mass, radius1, radius2, 0.7659, 0.986)
	} else if mass <= 1.261 {
		//radius = quad_trend(-0.0690127077, 0.4045260997, 0.5903311441, mass);
		radius1 := PlanetRadiusHelper(mass, 0.7659, 0.8554, 0.986, 0.9221, 1.261, 0.9907)
		radius2 := PlanetRadiusHelper(mass, 0.986, 0.9221, 1.261, 0.9907, 1.606, 1.062)
		radius = RangeAdjust(mass, radius1, radius2, 0.986, 1.261)
	} else if mass <= 1.606 {
		//radius = quad_trend(-0.0446111528, 0.3345668417, 0.6397483435, mass);
		radius1 := PlanetRadiusHelper(mass, 0.986, 0.9221, 1.261, 0.9907, 1.606, 1.062)
		radius2 := PlanetRadiusHelper(mass, 1.261, 0.9907, 1.606, 1.062, 2.036, 1.136)
		radius = RangeAdjust(mass, radius1, radius2, 1.261, 1.606)
	} else if mass <= 2.036 {
		//radius = quad_trend(-0.0323446774, 0.2898923383, 0.679857461, mass);
		radius1 := PlanetRadiusHelper(mass, 1.261, 0.9907, 1.606, 1.062, 2.036, 1.136)
		radius2 := PlanetRadiusHelper(mass, 1.606, 1.062, 2.036, 1.136, 2.568, 1.211)
		radius = RangeAdjust(mass, radius1, radius2, 1.606, 2.036)
	} else if mass <= 2.568 {
		//radius = quad_trend(-0.0188541262, 0.2277818406, 0.7503921064, mass);
		radius1 := PlanetRadiusHelper(mass, 1.606, 1.062, 2.036, 1.136, 2.568, 1.211)
		radius2 := PlanetRadiusHelper(mass, 2.036, 1.136, 2.568, 1.211, 3.226, 1.289)
		radius = RangeAdjust(mass, radius1, radius2, 2.036, 2.568)
	} else if mass <= 3.226 {
		//radius = quad_trend(-0.0131731218, 0.1948661011, 0.7974556375, mass);
		radius1 := PlanetRadiusHelper(mass, 2.036, 1.136, 2.568, 1.211, 3.226, 1.289)
		radius2 := PlanetRadiusHelper(mass, 2.568, 1.211, 3.226, 1.289, 4.032, 1.369)
		radius = RangeAdjust(mass, radius1, radius2, 2.568, 3.226)
	} else if mass <= 4.032 {
		//radius = quad_trend(-0.0089795106, 0.1644288708, 0.8520029117, mass);
		radius1 := PlanetRadiusHelper(mass, 2.568, 1.211, 3.226, 1.289, 4.032, 1.369)
		radius2 := PlanetRadiusHelper(mass, 3.226, 1.289, 4.032, 1.369, 5.018, 1.451)
		radius = RangeAdjust(mass, radius1, radius2, 3.226, 4.032)
	} else if mass <= 5.018 {
		//radius = quad_trend(-0.0063563019, 0.1406888322, 0.9050771807, mass);
		radius1 := PlanetRadiusHelper(mass, 3.226, 1.289, 4.032, 1.369, 5.018, 1.451)
		radius2 := PlanetRadiusHelper(mass, 4.032, 1.369, 5.018, 1.451, 6.216, 1.534)
		radius = RangeAdjust(mass, radius1, radius2, 4.032, 5.018)
	} else if mass <= 6.216 {
		//radius = quad_trend(-0.0042731856, 0.1172871044, 0.970053509, mass);
		radius1 := PlanetRadiusHelper(mass, 4.032, 1.369, 5.018, 1.451, 6.216, 1.534)
		radius2 := PlanetRadiusHelper(mass, 5.018, 1.451, 6.216, 1.534, 7.665, 1.618)
		radius = RangeAdjust(mass, radius1, radius2, 5.018, 6.216)
	} else if mass <= 7.665 {
		//radius = quad_trend(-0.0032875812, 0.1036059286, 1.017013265, mass);
		radius1 := PlanetRadiusHelper(mass, 5.018, 1.451, 6.216, 1.534, 7.665, 1.618)
		radius2 := PlanetRadiusHelper(mass, 6.216, 1.534, 7.665, 1.618, 9.39, 1.7)
		radius = RangeAdjust(mass, radius1, radius2, 6.216, 7.665)
	} else if mass <= 9.39 {
		//radius = quad_trend(-0.0019141267, 0.0801816636, 1.115866754, mass);
		radius1 := PlanetRadiusHelper(mass, 6.216, 1.534, 7.665, 1.618, 9.39, 1.7)
		radius2 := PlanetRadiusHelper(mass, 7.665, 1.618, 9.39, 1.7, 11.45, 1.783)
		radius = RangeAdjust(mass, radius1, radius2, 7.665, 9.39)
	} else if mass <= 11.45 {
		//radius = quad_trend(-0.0015393648, 0.0723716241, 1.156159475, mass);
		radius1 := PlanetRadiusHelper(mass, 7.665, 1.618, 9.39, 1.7, 11.45, 1.783)
		radius2 := PlanetRadiusHelper(mass, 9.39, 1.7, 11.45, 1.783, 13.91, 1.865)
		radius = RangeAdjust(mass, radius1, radius2, 9.39, 11.45)
	} else if mass <= 13.91 {
		//radius = quad_trend(-9.435581E-4, 0.057261966, 1.251053311, mass);
		radius1 := PlanetRadiusHelper(mass, 9.39, 1.7, 11.45, 1.783, 13.91, 1.865)
		radius2 := PlanetRadiusHelper(mass, 11.45, 1.783, 13.91, 1.865, 16.81, 1.947)
		radius = RangeAdjust(mass, radius1, radius2, 11.45, 13.91)
	} else if mass <= 16.81 {
		//radius = quad_trend(-7.534048E-4, 0.0514204578, 1.295516297, mass);
		radius1 := PlanetRadiusHelper(mass, 11.45, 1.783, 13.91, 1.865, 16.81, 1.947)
		radius2 := PlanetRadiusHelper(mass, 13.91, 1.865, 16.81, 1.947, 20.21, 2.027)
		radius = RangeAdjust(mass, radius1, radius2, 13.91, 16.81)
	} else if mass <= 20.21 {
		//radius = quad_trend(-5.047244E-4, 0.0422143076, 1.380000531, mass);
		radius1 := PlanetRadiusHelper(mass, 13.91, 1.865, 16.81, 1.947, 20.21, 2.027)
		radius2 := PlanetRadiusHelper(mass, 16.81, 1.947, 20.21, 2.027, 24.2, 2.106)
		radius = RangeAdjust(mass, radius1, radius2, 16.81, 20.21)
	} else if mass <= 24.2 {
		//radius = quad_trend(-3.750415E-4, 0.0364550938, 1.443426061, mass);
		radius1 := PlanetRadiusHelper(mass, 16.81, 1.947, 20.21, 2.027, 24.2, 2.106)
		radius2 := PlanetRadiusHelper(mass, 20.21, 2.027, 24.2, 2.106, 28.85, 2.183)
		radius = RangeAdjust(mass, radius1, radius2, 20.21, 24.2)
	} else if mass <= 28.85 {
		//radius = quad_trend(-2.610787E-4, 0.0304093647, 1.522991503, mass);
		radius1 := PlanetRadiusHelper(mass, 20.21, 2.027, 24.2, 2.106, 28.85, 2.183)
		radius2 := PlanetRadiusHelper(mass, 24.2, 2.106, 28.85, 2.183, 34.23, 2.258)
		radius = RangeAdjust(mass, radius1, radius2, 24.2, 28.85)
	} else if mass <= 34.23 {
		//radius = quad_trend(-1.90016E-4, 0.0259267321, 1.593168402, mass);
		radius1 := PlanetRadiusHelper(mass, 24.2, 2.106, 28.85, 2.183, 34.23, 2.258)
		radius2 := PlanetRadiusHelper(mass, 28.85, 2.183, 34.23, 2.258, 40.45, 2.331)
		radius = RangeAdjust(mass, radius1, radius2, 28.85, 34.23)
	} else if mass <= 40.45 {
		//radius = quad_trend(-1.446417E-4, 0.022538175, 1.655993898, mass);
		radius1 := PlanetRadiusHelper(mass, 28.85, 2.183, 34.23, 2.258, 40.45, 2.331)
		radius2 := PlanetRadiusHelper(mass, 34.23, 2.258, 40.45, 2.331, 47.59, 2.401)
		radius = RangeAdjust(mass, radius1, radius2, 34.23, 40.45)
	} else if mass <= 47.59 {
		//radius = quad_trend(-1.04715E-4, 0.0190230332, 1.732853309, mass);
		radius1 := PlanetRadiusHelper(mass, 34.23, 2.258, 40.45, 2.331, 47.59, 2.401)
		radius2 := PlanetRadiusHelper(mass, 40.45, 2.331, 47.59, 2.401, 55.76, 2.468)
		radius = RangeAdjust(mass, radius1, radius2, 40.45, 47.59)
	} else if mass <= 55.76 {
		//radius = quad_trend(-8.202615E-5, 0.016678137, 1.793060949, mass);
		radius1 := PlanetRadiusHelper(mass, 40.45, 2.331, 47.59, 2.401, 55.76, 2.468)
		radius2 := PlanetRadiusHelper(mass, 47.59, 2.401, 55.76, 2.468, 65.07, 2.531)
		radius = RangeAdjust(mass, radius1, radius2, 47.59, 55.76)
	} else if mass <= 65.07 {
		//radius = quad_trend(-5.984705E-5, 0.0139982359, 1.873533463, mass);
		radius1 := PlanetRadiusHelper(mass, 47.59, 2.401, 55.76, 2.468, 65.07, 2.531)
		radius2 := PlanetRadiusHelper(mass, 55.76, 2.468, 65.07, 2.531, 75.65, 2.59)
		radius = RangeAdjust(mass, radius1, radius2, 55.76, 65.07)
	} else if mass <= 75.65 {
		//radius = quad_trend(-4.398699E-5, 0.0117664086, 1.951605315, mass);
		radius1 := PlanetRadiusHelper(mass, 55.76, 2.468, 65.07, 2.531, 75.65, 2.59)
		radius2 := PlanetRadiusHelper(mass, 65.07, 2.531, 75.65, 2.59, 87.65, 2.645)
		radius = RangeAdjust(mass, radius1, radius2, 65.07, 75.65)
	} else if mass <= 87.65 {
		//radius = quad_trend(-2.918571E-5, 0.0093493602, 2.049748472, mass);
		radius1 := PlanetRadiusHelper(mass, 65.07, 2.531, 75.65, 2.59, 87.65, 2.645)
		radius2 := PlanetRadiusHelper(mass, 75.65, 2.59, 87.65, 2.645, 101.2, 2.697)
		radius = RangeAdjust(mass, radius1, radius2, 75.65, 87.65)
	} else if mass <= 101.2 {
		//radius = quad_trend(-2.427672E-5, 0.0084222976, 2.093292089, mass);
		radius1 := PlanetRadiusHelper(mass, 75.65, 2.59, 87.65, 2.645, 101.2, 2.697)
		radius2 := PlanetRadiusHelper(mass, 87.65, 2.645, 101.2, 2.697, 116.5, 2.745)
		radius = RangeAdjust(mass, radius1, radius2, 87.65, 101.2)
	} else if mass <= 116.5 {
		//radius = quad_trend(-1.821786E-5, 0.0071032835, 2.164724855, mass);
		radius1 := PlanetRadiusHelper(mass, 87.65, 2.645, 101.2, 2.697, 116.5, 2.745)
		radius2 := PlanetRadiusHelper(mass, 101.2, 2.697, 116.5, 2.745, 133.8, 2.789)
		radius = RangeAdjust(mass, radius1, radius2, 101.2, 116.5)
	} else if mass <= 133.8 {
		//radius = quad_trend(-1.286376E-5, 0.0057631526, 2.248182937, mass);
		radius1 := PlanetRadiusHelper(mass, 101.2, 2.697, 116.5, 2.745, 133.8, 2.789)
		radius2 := PlanetRadiusHelper(mass, 116.5, 2.745, 133.8, 2.789, 153.1, 2.829)
		radius = RangeAdjust(mass, radius1, radius2, 116.5, 133.8)
	} else if mass <= 153.1 {
		radius1 := PlanetRadiusHelper(mass, 116.5, 2.745, 133.8, 2.789, 153.1, 2.829)
		radius2 := PlanetRadiusHelper2(mass, 133.8, 2.789, 153.1, 2.829)
		radius = RangeAdjust(mass, radius1, radius2, 133.8, 153.1)
	} else {
		//radius = ln_trend(1.335486915, 0.2968566848, mass);
		radius = PlanetRadiusHelper2(mass, 133.8, 2.789, 153.1, 2.829)
	}

	if adjustForCarbon {
		rmf := 0.5
		carbonFraction := rmf * cmf
		growFactor := (0.05 * carbonFraction) + 1.0
		radius *= growFactor
	}

	return radius
}
