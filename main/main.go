package main

import (
	"log"
	"time"
)

func main() {
	output := "./dist"
	templates := "./templates"

	// create a new seed
	now := time.Now()
	seed := now.Unix()
	log.Printf("Seed: %d", seed)

	// generate and render star system
	err := Generate(templates, output, seed, "main")
	if err != nil {
		log.Panicln(err)
	}
}
