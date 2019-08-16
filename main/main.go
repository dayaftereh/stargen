package main

import (
	"log"

	"github.com/dayaftereh/stargen/mathf/random"

	"github.com/dayaftereh/stargen/stargen"

	"github.com/dayaftereh/stargen/display"
)

func main() {
	output := "./dist"
	templates := "./templates"

	displayTemplates, err := display.NewDisplayTemplate(templates, output)
	if err != nil {
		log.Panicln(err)
	}

	r := random.NewRandom()
	log.Printf("Seed: %d", r.Seed())
	sun, planets := stargen.Generate(r, true, true, true)

	displayContext := display.NewContext(sun, planets)

	err = displayTemplates.Generate(displayContext, "foo")
	if err != nil {
		log.Panicln(err)
	}
}
