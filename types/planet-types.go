package types

type PlanetType string

const (
	PlanetUnknown        PlanetType = "unknown"           // unknown planet type
	PlanetRock           PlanetType = "rock"              // callisto
	PlanetVenusian       PlanetType = "vebusian"          // venuslike
	PlanetTerrestrial    PlanetType = "terrestrial"       // Earthlike
	PlanetGasGiant       PlanetType = "gas-gaint"         // jupiterlike Jovian
	PlanetMartian        PlanetType = "martian"           // planet like mars
	PlanetWater          PlanetType = "water"             // planet with >95% water on the surface
	PlanetIce            PlanetType = "ice"               // pluto
	PlanetSubGasGiant    PlanetType = "sub-gas-gaint"     // gasgiant Sub-Jovian
	PlanetSubSubGasGiant PlanetType = "sub-sub-gas-gaint" // GasDwarf
	PlanetAsteroids      PlanetType = "asteroids"
	Planet1Face          PlanetType = "face"
	PlanetBrownDwarf     PlanetType = "brown-dwarf"
	PlanetIron           PlanetType = "iron"
	PlanetCarbon         PlanetType = "carbon"
	PlanetOil            PlanetType = "oil"
)
