package star

import (
	"github.com/dayaftereh/stargen/mathf/random"
)

type StellarClass struct {
	Class       string
	Color       string
	Mass        *Range
	Temperature *Range
}

// Stellar classification https://en.wikipedia.org/wiki/Stellar_classification

var StellarClassification []*StellarClass = []*StellarClass{
	&StellarClass{
		Class:       "O",
		Color:       "#0000ff",             // blue
		Mass:        NewRange(16.0, 120.0), // >= 16 M
		Temperature: NewRange(30, 50),      // >= 30 *1000 K
	},
	&StellarClass{
		Class:       "B",
		Color:       "#82aeff",           // blue white
		Mass:        NewRange(2.1, 16.0), // 2.1–16 M
		Temperature: NewRange(10, 30),    // 10 - 30 *1000 K
	},
	&StellarClass{
		Class:       "A",
		Color:       "#ffffff",          // white
		Mass:        NewRange(1.4, 2.1), // 1.4 - 2.1 M
		Temperature: NewRange(7.5, 10),  // 7.5 - 10 *1000 K
	},
	&StellarClass{
		Class:       "F",
		Color:       "#fffd9e",           // yellow white
		Mass:        NewRange(1.04, 1.4), // 1.04 - 1.4 M
		Temperature: NewRange(6, 7.5),    // 6 - 7.5 *1000 K
	},
	&StellarClass{
		Class:       "G",
		Color:       "#ffff00",           // yellow
		Mass:        NewRange(0.8, 1.04), // 0.8 - 1.04 M
		Temperature: NewRange(5.2, 6.0),  // 5.2 - 6 *1000 K
	},
	&StellarClass{
		Class:       "K",
		Color:       "#ffd659",           // light orange
		Mass:        NewRange(0.45, 0.8), // 0.45 - 0.8 M
		Temperature: NewRange(3.7, 5.2),  // 3.7 - 5.2 *1000 K
	},
	&StellarClass{
		Class:       "M",
		Color:       "#ff903b",            // orange red
		Mass:        NewRange(0.08, 0.45), // 0.08 - 0.45 M
		Temperature: NewRange(2.4, 3.7),   // 2.4 - 3.7 *1000 K
	},
}

func RandomStellarClass(random random.Random) *StellarClass {
	length := len(StellarClassification)
	index := random.RandIntn(length)
	class := StellarClassification[index]
	return class
}
