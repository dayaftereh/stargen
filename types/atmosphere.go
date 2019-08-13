package types

type AtmosphereType string

const (
	None         AtmosphereType = "none"
	Toxic        AtmosphereType = "toxic"
	Breathable   AtmosphereType = "breathable"
	Unbreathable AtmosphereType = "unbreathable"
)
