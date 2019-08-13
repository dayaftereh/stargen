package star

import (
	"github.com/dayaftereh/stargen/mathf/random"
)

type Range struct {
	Min float64
	Max float64
}

func NewRange(min, max float64) *Range {
	return &Range{
		Min: min,
		Max: max,
	}
}

func (r *Range) random(random random.Random) float64 {
	return random.RandFloat64(r.Min, r.Max)
}
