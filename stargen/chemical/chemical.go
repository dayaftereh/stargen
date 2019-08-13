package chemical

//  				             Weight
// &Chemical{1, "H", "Hydrogen", 1.0079, 14.06, 20.40, 8.99e-05, 0.0, 0.0, 0.0, 0.00125893, 27925.4, 1, 0.0, 0.0}

type Chemical struct {
	Num          int64   `json:"num"` // Atomic number
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Weight       float64 `json:"num"`          // Standard atomic weight
	MeltingPoint float64 `json:"meltingPoint"` // (kelvin)
	BoilingPoint float64 `json:"boilingPoint"` // (kelvin)
	Density      float64 `json:"density"`      // 8.99e-05 (0.08988 g/L)
	PZero        float64 `json:"num"`          // 0.0
	C            float64 `json:"num"`          // 0.0
	N            float64 `json:"num"`          // 0.0
	Abunde       float64 `json:"num"`          // 0.00125893
	Abunds       float64 `json:"num"`          // 27925.4
	Reactivity   float64 `json:"num"`          // 1
	MaxIpp       float64 `json:"num"`          // 0.0
	MinIpp       float64 `json:"num"`          // 0.0
}
