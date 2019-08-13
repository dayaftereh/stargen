package types

type Sun struct {
	Class                string  `json:"class"`
	Color                string  `json:"color"`
	EffectiveTemperature float64 `json:"effectiveTemperature"` // star temperature in Kelvins
	Mass                 float64 `json:"mass"`                 // star mass (units of solar masses)
	Luminosity           float64 `json:"luminosity"`           // the stellar luminosity ratio is with respect to the sun
	EcosphereRadius      float64 `json:"ecosphereRadius"`      // the estimate range from the sun allowing existence of liquid water, Habitable ecosphere radius (AU)
	Life                 float64 `json:"life"`                 // The max life of the sun (years)
	Age                  float64 `json:"age"`                  // The age of the sun (years)
}
