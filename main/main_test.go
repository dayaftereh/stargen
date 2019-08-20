package main_test

import (
	"testing"

	"github.com/dayaftereh/stargen/main"
)

// Blue 1565959868174024600
// Big Blue 1565961089167918100

const (
	templates = "../templates"
	output    = "../dist"
)

func Generate(t *testing.T, seed int64, name string) {
	err := main.Generate(templates, output, seed, name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultiplyStargens(t *testing.T) {
	// O / Big Blue / 1565961089167918100
	Generate(t, 1565961089167918100, "o-blue")
	// B / blue white / 1566214142
	Generate(t, 1566214142, "b-blue-white")
	// A / white / 1566214203
	Generate(t, 1566214203, "a-white")
	// F / White Yellow / 1566214005
	Generate(t, 1566214005, "f-yellow-white")
	// G / Yellow / 1566214299
	Generate(t, 1566214299, "g-yellow")
	// K / light-orange / 1566214086
	Generate(t, 1566214086, "k-light-orange")
	// M / orange red / 1566213952
	Generate(t, 1566213952, "m-orange-red")
}
