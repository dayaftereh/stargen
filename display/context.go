package display

import "github.com/dayaftereh/stargen/types"

type Context struct {
	Sun     *types.Sun
	Planets []*types.Planet
}

func NewContext(sun *types.Sun, planets []*types.Planet) *Context {
	return &Context{
		Sun:     sun,
		Planets: planets,
	}
}
