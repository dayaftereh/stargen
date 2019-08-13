package types

type Planet struct {
	Type                            PlanetType     `json:"type"`                            /* Type code */
	SemiMajorAxis                   float64        `json:"semiMajorAxis"`                   /* A - (Distance from primary star) semi-major axis of solar orbit (in AU)*/
	Eccentricity                    float64        `json:"eccentricity"`                    /* E - eccentricity of solar orbit */
	AxialTilt                       float64        `json:"axialTilt"`                       /* units of degrees */
	Mass                            float64        `json:"mass"`                            /* mass (in solar masses) */
	GasGiant                        bool           `json:"gasGiant"`                        /* TRUE if the planet is a gas giant */
	DustMass                        float64        `json:"dustMass"`                        /* mass, ignoring gas */
	GasMass                         float64        `json:"gasMass"`                         /* mass, ignoring dust */
	CoreRadius                      float64        `json:"coreRadius"`                      /* radius of the rocky core (in km) */
	Radius                          float64        `json:"radius"`                          /* equatorial radius (in km)	*/
	OrbitZone                       OrbitZone      `json:"orbitZone"`                       /* the 'zone' of the planet */
	Density                         float64        `json:"density"`                         /* density (in g/cc) */
	OrbitPeriod                     float64        `json:"orbitPeriod"`                     /* length of the local year (days) */
	Day                             float64        `json:"day"`                             /* length of the local day (hours) */
	ResonantPeriod                  bool           `json:"resonantPeriod"`                  /* TRUE if in resonant rotation (Planet's rotation is in a resonant spin lock with the star)		 */
	EscapeVelocity                  float64        `json:"escapeVelocity"`                  /* Escape Velocity (cm/sec) */
	SurfaceAcceleration             float64        `json:"surfaceAcceleration"`             /* Surface acceleration (cm/sec2	) */
	SurfaceGravity                  float64        `json:"surfaceGravity"`                  /* Surface gravity (Earth gravities) */
	RootMeanSquareVelocity          float64        `json:"rootMeanSquareVelocity"`          /* Root Mean Square Velocity units of cm/sec */
	MolecularWeight                 float64        `json:"molecularWeight"`                 /* Molecular weight smallest molecular weight retained */
	VolatileGasInventory            float64        `json:"volatileGasInventory"`            /**/
	SurfacePressure                 float64        `json:"surfacePressure"`                 /* Surface pressure (millibars [mb]) */
	GreenhouseEffect                bool           `json:"greenhouseEffect"`                /* runaway greenhouse effect? */
	BoilPoint                       float64        `json:"boilPoint"`                       /* the boiling point of water (Kelvin) */
	Albedo                          float64        `json:"albedo"`                          /* albedo of the planet */
	ExosphericTemperature           float64        `json:"exosphericTemperature"`           /* units of degrees Kelvin */
	EstimatedTemperature            float64        `json:"estimatedTemperature"`            /* quick non-iterative estimate (K) */
	EstimatedTerrestrialTemperature float64        `json:"estimatedTerrestrialTemperature"` /* for terrestrial moons and the like */
	SurfaceTemperature              float64        `json:"surfaceTemperature"`              /* surface temperature in Kelvin */
	GreenhouseRise                  float64        `json:"greenhouseRise"`                  /* Temperature rise due to greenhouse */
	HighTemperature                 float64        `json:"highTemperature"`                 /* Day-time temperature */
	LowTemperature                  float64        `json:"lowTemperature"`                  /* Night-time temperature */
	MaxTemperature                  float64        `json:"maxTemperature"`                  /* Summer/Day */
	MinTemperature                  float64        `json:"minTemperature"`                  /* Winter/Night */
	Hydrosphere                     float64        `json:"hydrosphere"`                     /* fraction of surface covered water (%) */
	CloudCover                      float64        `json:"cloudCover"`                      /* fraction of surface covered (%) */
	IceCover                        float64        `json:"iceCover"`                        /* fraction of surface covered (%) */
	IceMassFraction                 float64        `json:"iceMassFraction"`                 /* ice mass fraction (%) imf*/
	RockMassFraction                float64        `json:"rockMassFraction"`                /* rock mass fraction (%) rmf*/
	CarbonMassFraction              float64        `json:"carbonMassFraction"`              /* fraction of rock that's carbon instead of silicate (%) cmf*/
	Atmosphere                      []*Gas         `json:"atmosphere"`                      /* The gases of the atmosphere */
	AtmosphereType                  AtmosphereType `json:"atmosphereType"`                  /* the type of the atmosphere */
	Moons                           []*Planet      `json:"moons"`                           /* list of moons for this planet */
}
