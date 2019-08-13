package stargen

import (
	"github.com/dayaftereh/stargen/mathf/random"
	"github.com/dayaftereh/stargen/stargen/accretion"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/stargen/environment"
	"github.com/dayaftereh/stargen/stargen/star"
	"github.com/dayaftereh/stargen/types"
)

func Generate(random random.Random, randomTilt bool, doMoons bool, doGases bool) (*types.Sun, []*types.Planet) {
	// create a star producer
	starProducer := star.NewStarProducer(random)

	// generate a new sun
	sun := starProducer.Create()

	// create the accretion processor for the planetesimal
	accretionProcessor := accretion.NewAccretionProcessor(random)
	// calculate the limit of stellar dust based on the mass of the sun
	outerDustLimit := accretionProcessor.StellarDustLimit(sun.Mass)

	// create the planetesimal
	accretionProcessor.DistPlanetaryMasses(
		sun.Mass,
		sun.Luminosity,
		0.0,
		outerDustLimit,
		0.0,
		constants.DustDensityCoefficient,
		true,
	)

	// get the planetesimal
	planets := make([]*types.Planet, 0)
	// get each planet
	accretion.ForEach(accretionProcessor.Planets, func(value interface{}, index int64) {
		planet := value.(*types.Planet)
		// create planet environment builder
		planetEnvironment := environment.NewPlanetEnvironment(random)
		// generate the planetesimal to a planet
		planetEnvironment.GeneratePlanet(sun, planet, randomTilt, doMoons, doGases, false)
		// append the found planet
		planets = append(planets, planet)
	})

	return sun, planets
}
