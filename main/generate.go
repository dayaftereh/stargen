package main

import (
	"log"

	"github.com/dayaftereh/stargen/display"
	"github.com/dayaftereh/stargen/mathf/random"
	"github.com/dayaftereh/stargen/stargen"
)

func Generate(templates string, output string, seed int64, name string) error {
	// load the templates and output for html rendering
	displayTemplates, err := display.NewDisplayTemplate(templates, output)
	if err != nil {
		log.Panicln(err)
	}
	// create a new random with given seed
	r := random.NewRandomWith(seed)
	// generate sun and planets
	sun, planets := stargen.Generate(r, true, true, true)

	log.Printf("Sun-Class: %s", sun.Class)

	// create the output context
	displayContext := display.NewContext(sun, planets)
	// generate and render the html
	err = displayTemplates.Generate(displayContext, name)
	return err
}
