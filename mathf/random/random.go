package random

import (
	"math/rand"
	"time"
)

type Random interface {
	Seed() int64
	About(value float64, variation float64) float64
	RandFloat64(min float64, max float64) float64
	RandIntn(n int) int
}

type internalRandom struct {
	seed   int64
	random *rand.Rand
}

func NewRandom() Random {
	now := time.Now()
	seed := now.UnixNano()
	return NewRandomWith(seed)
}
func NewRandomWith(seed int64) Random {
	source := rand.NewSource(seed)
	random := rand.New(source)
	return &internalRandom{
		seed:   seed,
		random: random,
	}
}

func (random *internalRandom) Seed() int64 {
	return random.seed
}

func (random *internalRandom) About(value float64, variation float64) float64 {
	return (value + (value * random.RandFloat64(-variation, variation)))
}

func (random *internalRandom) RandFloat64(min float64, max float64) float64 {
	return min + random.random.Float64()*(max-min)
}

func (random *internalRandom) RandIntn(n int) int {
	return random.random.Intn(n)
}
