package star

import (
	"math"

	"github.com/dayaftereh/stargen/mathf/random"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

type StarProducer struct {
	random random.Random
}

func NewStarProducer(random random.Random) *StarProducer {
	return &StarProducer{
		random: random,
	}
}

func (starProducer *StarProducer) Create() *types.Sun {
	// get a random stellar class
	stellarClass := RandomStellarClass(starProducer.random)
	// generate a random star mass based on the stellar class
	mass := stellarClass.Mass.random(starProducer.random)

	// calculate luminosity based on the mass
	luminosity := massToLuminosity(mass)

	// calculate the life of the sun
	life := 1e10 * (mass / luminosity)

	// check if life lager then max age
	maxAge := constants.MaxSunAge
	if life > maxAge {
		maxAge = life
	}

	// generate a random age of the sun
	age := starProducer.random.RandFloat64(constants.MinSunAge, maxAge)

	// generate a random effective temperature based on the stellar class
	effectiveTemperature := stellarClass.Temperature.random(starProducer.random) * 1000.0

	ecosphereRadius := math.Sqrt(luminosity)

	// create a new sun
	return &types.Sun{
		Age:                  age,
		Life:                 life,
		Mass:                 mass,
		Luminosity:           luminosity,
		EcosphereRadius:      ecosphereRadius,
		EffectiveTemperature: effectiveTemperature,
		Class:                stellarClass.Class,
		Color:                stellarClass.Color,
	}
}
