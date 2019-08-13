package accretion

import (
	"container/list"
	"math"

	"github.com/dayaftereh/stargen/mathf/random"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/stargen/eccentricity"
	"github.com/dayaftereh/stargen/types"
)

type AccretionProcessor struct {
	CloudEccentricity float64
	Planets           *list.List
	DustBands         *list.List
	random            random.Random
}

func NewAccretionProcessor(random random.Random) *AccretionProcessor {
	return &AccretionProcessor{
		random:            random,
		CloudEccentricity: 0.2,
		DustBands:         list.New(),
		Planets:           list.New(),
	}
}

func (accretionProcessor *AccretionProcessor) StellarDustLimit(stellarMassRatio float64) float64 {
	return (200.0 * math.Pow(stellarMassRatio, (1.0/3.0)))
}

func (accretionProcessor *AccretionProcessor) initialDustBand(innerDust float64, outerDust float64) {
	first := &DustBand{
		InnerEdge:   innerDust,
		OuterEdge:   outerDust,
		DustPresent: true,
		GasPresent:  true,
	}
	// add the first dust band
	accretionProcessor.DustBands.PushBack(first)
}

func (accretionProcessor *AccretionProcessor) nearestPlanet(stellarMassRatio float64) float64 {
	return (0.3 * math.Pow(stellarMassRatio, (1.0/3.0)))
}

func (accretionProcessor *AccretionProcessor) farthestPlanet(stellarMassRatio float64) float64 {
	return (50.0 * math.Pow(stellarMassRatio, (1.0/3.0)))
}

func (accretionProcessor *AccretionProcessor) innerEffectLimit(a, e, mass float64) float64 {
	return (a * (1.0 - e) * (1.0 - mass) / (1.0 + accretionProcessor.CloudEccentricity))

}

func (accretionProcessor *AccretionProcessor) outerEffectLimit(a, e, mass float64) float64 {
	return (a * (1.0 + e) * (1.0 + mass) / (1.0 - accretionProcessor.CloudEccentricity))

}

func (accretionProcessor *AccretionProcessor) isDustAvailable(insideRange, outsideRange float64) bool {
	// find the first dustband intersection
	dustBandElement := Find(accretionProcessor.DustBands, func(v interface{}, index int64) bool {
		return !(v.(*DustBand).OuterEdge < insideRange)
	})

	// no dust band element found
	if dustBandElement == nil {
		return false
	}

	dustHere := dustBandElement.Value.(*DustBand).DustPresent

	// find the next dust bands for the given range
	dustBandElement = FindAfter(dustBandElement, func(v interface{}) bool {
		dustBand := v.(*DustBand)
		// check if dust is till here
		dustHere = dustHere || dustBand.DustPresent
		return !(dustBand.InnerEdge < outsideRange)
	})

	return dustHere
}

func (accretionProcessor *AccretionProcessor) criticalLimit(orbRadius, eccentricity, stellarLuminosityRatio float64) float64 {
	perihelionDist := (orbRadius - orbRadius*eccentricity)
	temp := perihelionDist * math.Sqrt(stellarLuminosityRatio)
	return (constants.B * math.Pow(temp, -0.75))
}

func (accretionProcessor *AccretionProcessor) collectDust(lastMass, a, e, critMass, dustDensity float64, dustBandElement *list.Element) (float64, float64, float64, float64, float64, float64) {
	temp := lastMass / (1.0 + lastMass)
	reducedMass := math.Pow(temp, (1.0 / 4.0))

	rInner := accretionProcessor.innerEffectLimit(a, e, reducedMass)
	rOuter := accretionProcessor.outerEffectLimit(a, e, reducedMass)

	if rInner < 0.0 {
		rInner = 0.0
	}

	// check if a dust band left
	if dustBandElement == nil {
		return 0.0, 0.0, 0.0, rInner, rOuter, reducedMass
	}

	// get the dust band
	dustBand := dustBandElement.Value.(*DustBand)

	tempDensity := dustDensity
	if !dustBand.DustPresent {
		tempDensity = 0.0
	}

	gasDensity := 0.0
	massDensity := tempDensity
	if !(lastMass < critMass || !dustBand.GasPresent) {
		massDensity = constants.K * tempDensity / (1.0 + math.Sqrt(critMass/lastMass)*(constants.K-1.0))
		gasDensity = massDensity - tempDensity
	}

	if dustBand.OuterEdge <= rInner || dustBand.InnerEdge >= rOuter {
		return accretionProcessor.collectDust(lastMass, a, e, critMass, dustDensity, dustBandElement.Next())
	}

	bandwidth := (rOuter - rInner)
	temp1 := rOuter - dustBand.OuterEdge
	if temp1 < 0.0 {
		temp1 = 0.0
	}

	width := bandwidth - temp1

	temp2 := dustBand.InnerEdge - rInner
	if temp2 < 0.0 {
		temp2 = 0.0
	}

	width = width - temp2

	temp = 4.0 * math.Pi * math.Pow(a, 2.0) * reducedMass * (1.0 - e*(temp1-temp2)/bandwidth)
	volume := temp * width

	newMass := (volume * massDensity)

	newGas := (volume * gasDensity)
	newDust := (newMass - newGas)

	nextMass, nextDust, nextGas, rInner, rOuter, reducedMass := accretionProcessor.collectDust(lastMass, a, e, critMass, dustDensity, dustBandElement.Next())

	newMass += nextMass
	newGas += nextGas
	newDust += nextDust

	return newMass, newDust, newGas, rInner, rOuter, reducedMass
}

func (accretionProcessor *AccretionProcessor) updateDustLanes(min, max, mass, critMass, bodyInnerBound, bodyOuterBound float64) bool {
	// check if gas giant
	gas := !(mass > critMass)

	// get each dust band for intersection of min && max
	ForEachElement(accretionProcessor.DustBands, func(element *list.Element, index int64) {
		// get the current dust band
		dustBand := element.Value.(*DustBand)

		// A: boundary is between the found dust band
		// | ------------------------------ | => before
		// |          -          +            => intersection
		// | -------- | -------- | -------- | => after
		if dustBand.InnerEdge < min && dustBand.OuterEdge > max {
			// create a new split dustband
			splitDustBand := &DustBand{
				InnerEdge:   min,
				OuterEdge:   max,
				DustPresent: false,
				GasPresent:  false,
			}

			// insert the split dust band
			middleElement := accretionProcessor.DustBands.InsertAfter(splitDustBand, element)

			// check if the dust band has gas
			if dustBand.GasPresent {
				splitDustBand.GasPresent = gas
			}

			// create the end dust band
			endDustBund := &DustBand{
				InnerEdge:   max,
				OuterEdge:   dustBand.OuterEdge,
				GasPresent:  dustBand.GasPresent,
				DustPresent: dustBand.DustPresent,
			}
			// insert the end dust band after middle
			accretionProcessor.DustBands.InsertAfter(endDustBund, middleElement)

			// set the first dustband back to start of the split dust band
			dustBand.OuterEdge = min

			// goto next dust band
			return
		}

		// B: max contains dust band
		//    | ------------------ | => before
		// -          +              => intersection
		// |----------| ---------- | => after
		if dustBand.InnerEdge < max && dustBand.OuterEdge > max {
			// create the next dust band
			next := &DustBand{
				InnerEdge:   max,
				OuterEdge:   dustBand.OuterEdge,
				DustPresent: dustBand.DustPresent,
				GasPresent:  dustBand.GasPresent,
			}

			// insert after current dustband
			accretionProcessor.DustBands.InsertAfter(next, element)

			// update current dustband
			dustBand.OuterEdge = max
			dustBand.DustPresent = false
			if dustBand.GasPresent {
				dustBand.GasPresent = gas
			} else {
				dustBand.GasPresent = false
			}

			// goto next dust band
			return
		}

		// C: min contains dust band
		// | ------------------ |    => before
		//            -            + => intersection
		// |----------| ---------- | => after
		if dustBand.InnerEdge < min && dustBand.OuterEdge > min {
			// create the next dust band
			next := &DustBand{
				InnerEdge:   min,
				OuterEdge:   dustBand.OuterEdge,
				DustPresent: false,
				GasPresent:  false,
			}

			// check if next dust band has gas
			if dustBand.GasPresent {
				next.GasPresent = gas
			}

			// insert after current dustband
			accretionProcessor.DustBands.InsertAfter(next, element)

			// move current dust band to min
			dustBand.OuterEdge = min

			// goto next dust band
			return
		}

		// D: dust band is between boundary
		//   | ------------------ |    => before
		// -                         + => intersection
		//   | ------------------ |    => after
		if dustBand.InnerEdge >= min && dustBand.OuterEdge <= max {
			// check if dust band has sill gas
			if dustBand.GasPresent {
				dustBand.GasPresent = gas
			}
			// remove the dust
			dustBand.DustPresent = false

			// goto next dust band
			return
		}
	})

	dustLeft := false
	ForEachElement(accretionProcessor.DustBands, func(element *list.Element, index int64) {
		// get the current dust band
		dustBand := element.Value.(*DustBand)

		// | ----------------------------------- | => found dust band
		//           -              +              => bodyInnerBound && bodyOuterBound
		if dustBand.DustPresent && dustBand.OuterEdge >= bodyInnerBound && dustBand.InnerEdge <= bodyOuterBound {
			// the dust band has dust
			dustLeft = true
		}

		// get the nest element
		nextElement := element.Next()
		// check if next element available
		if nextElement != nil {
			nextDustBand := nextElement.Value.(*DustBand)
			// check if dust bands are equal
			if dustBand.DustPresent == nextDustBand.DustPresent && dustBand.GasPresent == nextDustBand.GasPresent {
				// join both dust bands
				dustBand.OuterEdge = nextDustBand.OuterEdge
				// remove the next element
				accretionProcessor.DustBands.Remove(nextElement)
			}
		}
	})

	return dustLeft
}

func (accretionProcessor *AccretionProcessor) accreteDust(seedMass, a, e, critMass, bodyInnerBound, bodyOuterBound, dustDensity float64) (float64, float64, float64, float64, bool) {
	tempMass := 0.0
	newMass := seedMass
	newDust, newGas, rInner, rOuter, reducedMass := 0.0, 0.0, 0.0, 0.0, 0.0

	dustBundElement := accretionProcessor.DustBands.Front()

	for {
		tempMass = newMass

		newMass, newDust, newGas, rInner, rOuter, reducedMass = accretionProcessor.collectDust(newMass, a, e, critMass, dustDensity, dustBundElement)
		if (newMass - tempMass) < (0.0001 * tempMass) {
			break
		}
	}

	seedMass += newMass

	dustLeft := accretionProcessor.updateDustLanes(rInner, rOuter, seedMass, critMass, bodyInnerBound, bodyOuterBound)

	return seedMass, newDust, newGas, reducedMass, dustLeft

}

func (accretionProcessor *AccretionProcessor) coalescePlanetesimals(a, e, mass, critMass, dustMass, gasMass, stellarLuminosityRatio, bodyInnerBound, bodyOuterBound, reducedMass, dustDensity float64, doMoons, dustLeft bool) bool {
	// First we try to find an existing planet with an over-lapping orbit.
	found := FindByElement(accretionProcessor.Planets, func(element *list.Element, index int64) bool {
		// get the current planet
		planet := element.Value.(*types.Planet)

		diff := planet.SemiMajorAxis - a

		var dist1, dist2 float64
		if diff > 0.0 {
			dist1 = (a * (1.0 + e) * (1.0 + reducedMass)) - a
			/* x aphelion	 */
			reducedMass = math.Pow((planet.Mass / (1.0 + planet.Mass)), (1.0 / 4.0))
			dist2 = planet.SemiMajorAxis - (planet.SemiMajorAxis * (1.0 - planet.Eccentricity) * (1.0 - reducedMass))
		} else {
			dist1 = a - (a * (1.0 - e) * (1.0 - reducedMass))
			/* x perihelion */
			reducedMass = math.Pow((planet.Mass / (1.0 + planet.Mass)), (1.0 / 4.0))
			dist2 = (planet.SemiMajorAxis * (1.0 + planet.Eccentricity) * (1.0 + reducedMass)) - planet.SemiMajorAxis
		}

		// check if planets over-lap
		if (math.Abs(diff) <= math.Abs(dist1)) || (math.Abs(diff) <= math.Abs(dist2)) {
			newA := (planet.Mass + mass) / ((planet.Mass / planet.SemiMajorAxis) + (mass / a))

			temp := planet.Mass * math.Sqrt(planet.SemiMajorAxis) * math.Sqrt(1.0-math.Pow(planet.Eccentricity, 2.0))
			temp = temp + (mass * math.Sqrt(a) * math.Sqrt(math.Sqrt(1.0-math.Pow(e, 2.0))))
			temp = temp / ((planet.Mass + mass) * math.Sqrt(newA))
			temp = 1.0 - math.Pow(temp, 2.0)

			if temp < 0.0 || temp >= 1.0 {
				temp = 0.0
			}

			e = math.Sqrt(temp)

			// check if overlap can be a moon
			if doMoons {
				existingMass := 0.0

				// calculate the hole mass
				for _, moon := range planet.Moons {
					existingMass += moon.Mass
				}

				// check if the mas is crit
				if mass < critMass {
					// check if the mass fits into the values
					if (mass*constants.SunMassInEarthMasses) < 2.5 && (mass*constants.SunMassInEarthMasses) > 0.000001 && existingMass < (planet.Mass*0.05) {
						moon := &types.Planet{
							Type:     types.PlanetUnknown,
							Mass:     mass,
							DustMass: dustMass,
							GasMass:  gasMass,
							GasGiant: false,
							Moons:    make([]*types.Planet, 0),
						}

						// check if the moon more mass then the planet
						if (moon.DustMass + moon.GasMass) > (planet.DustMass + planet.GasMass) {
							// swap moon and planet
							tmpDustMass := planet.DustMass
							tmpGasMass := planet.GasMass
							tmpMass := planet.Mass

							planet.DustMass = moon.DustMass
							planet.GasMass = moon.GasMass
							planet.Mass = moon.Mass

							moon.DustMass = tmpDustMass
							moon.GasMass = tmpGasMass
							moon.Mass = tmpMass
						}
						// add the moon
						planet.Moons = append(planet.Moons, moon)

						// stop
						return true
					}
				}
			}

			// Collision between two planetesimals!
			// join both planet to a new planet
			temp = planet.Mass + mass

			seedMass, newDust, newGas := 0.0, 0.0, 0.0
			seedMass, newDust, newGas, _, dustLeft = accretionProcessor.accreteDust(temp, newA, e, stellarLuminosityRatio, bodyInnerBound, bodyOuterBound, dustDensity)

			planet.SemiMajorAxis = newA
			planet.Eccentricity = e
			planet.Mass = seedMass
			planet.DustMass += dustMass + newDust
			planet.GasMass += gasMass + newGas

			// check if planet is a gas giant
			if seedMass >= critMass {
				planet.GasGiant = true
			}

			// find a planet further than new a from the sun to insert before
			nextElement := FindAfter(element, func(value interface{}) bool {
				nextPlanet := value.(*types.Planet)
				return nextPlanet.SemiMajorAxis > newA
			})

			// found a planet
			if nextElement != nil {
				// move planet after the found further planet
				accretionProcessor.Planets.MoveBefore(element, nextElement)
			}

			// found a planet place
			return true
		}

		// goto next planet
		return false
	})

	// Planetesimals didn't collide. Make it a planet.
	if found == nil {
		// create a new planet
		planet := &types.Planet{
			SemiMajorAxis: a,
			Eccentricity:  e,
			Mass:          mass,
			DustMass:      dustMass,
			GasMass:       gasMass,
			GasGiant:      mass >= critMass,
			Moons:         make([]*types.Planet, 0),
			Type:          types.PlanetUnknown,
		}

		// find a planet further than new planet from the sun to insert before
		nextElement := Find(accretionProcessor.Planets, func(value interface{}, index int64) bool {
			nextPlanet := value.(*types.Planet)
			return nextPlanet.SemiMajorAxis > a
		})

		// check if a planet was found
		if nextElement != nil {
			// insert before the found planet
			accretionProcessor.Planets.InsertBefore(planet, nextElement)
		} else {
			// add to the end, because new planet is the furthest
			accretionProcessor.Planets.PushBack(planet)
		}
	}

	return dustLeft

}

func (accretionProcessor *AccretionProcessor) DistPlanetaryMasses(stellarMassRatio, stellarLuminosityRatio, innerDust, outerDust, outerPlanetLimit, dustDensityCoefficient float64, doMoons bool) {
	// create the initial dust band
	accretionProcessor.initialDustBand(innerDust, outerDust)

	planetInnerBound := accretionProcessor.nearestPlanet(stellarMassRatio)
	planetOuterBound := accretionProcessor.farthestPlanet(stellarMassRatio)

	dustLeft := true
	for dustLeft {
		a := accretionProcessor.random.RandFloat64(planetInnerBound, planetOuterBound)
		e := eccentricity.Rand(accretionProcessor.random)

		mass := constants.ProtoPlanetMass

		innerEffectLimit := accretionProcessor.innerEffectLimit(a, e, mass)
		outerEffectLimit := accretionProcessor.outerEffectLimit(a, e, mass)

		if accretionProcessor.isDustAvailable(innerEffectLimit, outerEffectLimit) {
			dustDensity := dustDensityCoefficient * math.Sqrt(stellarMassRatio) * math.Exp(-constants.Alpha*math.Pow(a, (1.0/constants.N)))
			critMass := accretionProcessor.criticalLimit(a, e, stellarLuminosityRatio)

			dustMass, gasMass, reducedMass := 0.0, 0.0, 0.0
			mass, dustMass, gasMass, reducedMass, dustLeft = accretionProcessor.accreteDust(mass, a, e, critMass, planetInnerBound, planetOuterBound, dustDensity)

			dustMass += constants.ProtoPlanetMass

			if mass > constants.ProtoPlanetMass {
				dustLeft = accretionProcessor.coalescePlanetesimals(a, e, mass, critMass, dustMass, gasMass, stellarLuminosityRatio, planetInnerBound, planetOuterBound, reducedMass, dustDensity, doMoons, dustLeft)
			}
		}

	}
}
