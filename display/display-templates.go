package display

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"

	"github.com/dayaftereh/stargen/mathf"
	"github.com/dayaftereh/stargen/stargen"
	"github.com/dayaftereh/stargen/stargen/constants"
	"github.com/dayaftereh/stargen/types"
)

type DisplayTemplate struct {
	output   string
	template *template.Template
}

func NewDisplayTemplate(templateDirectory string, output string) (*DisplayTemplate, error) {
	tmpl, err := loadTemplates(templateDirectory)
	if err != nil {
		return nil, err
	}
	return &DisplayTemplate{
		output:   output,
		template: tmpl,
	}, nil
}

func loadTemplates(templateDirectory string) (*template.Template, error) {
	htmlGlob := path.Join(templateDirectory, "*.html")
	tmpl, err := template.New("index.html").Funcs(defaultFuncMap()).ParseGlob(htmlGlob)
	if err != nil {
		return nil, err
	}

	jsGlob := path.Join(templateDirectory, "js", "*.js")
	tmpl, err = tmpl.ParseGlob(jsGlob)
	if err != nil {
		return nil, err
	}

	shaderGlob := path.Join(templateDirectory, "shader", "*")
	tmpl, err = tmpl.ParseGlob(shaderGlob)

	return tmpl, err
}

func defaultFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"json": func(v interface{}) string {
			bytes, err := json.Marshal(v)
			if err != nil {
				return err.Error()
			}

			return string(bytes)
		},
		"orbits": func(context *Context) string {
			return Orbits(context.Planets, context.Sun)
		},
	}

	return funcMap
}

func Orbits(planets []*types.Planet, sun *types.Sun) string {
	orbits := stargen.Orbits(planets, sun)

	planetsPositions := make([][]*mathf.Vec3, 0)

	for _, o := range orbits {
		period := o.Period()

		positions := make([]*mathf.Vec3, 0)

		loops := 1000
		for i := 1; i < loops; i++ {
			dt := float64(i) * (period / float64(loops))
			newOrbit := o.Update(dt)
			position := newOrbit.Position()
			positions = append(positions, mathf.NewVec3(
				position.X/constants.KMPerAU, position.Y/constants.KMPerAU, position.Z/constants.KMPerAU,
			))
		}

		planetsPositions = append(planetsPositions, positions)
	}

	bytes, err := json.Marshal(planetsPositions)
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

func (displayTemplate *DisplayTemplate) Generate(context *Context, name string) error {
	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)

	// execute and output the template
	err := displayTemplate.template.Execute(buffer, context)
	if err != nil {
		return err
	}

	// check if the output directory exists
	exists, err := Exists(displayTemplate.output)
	if err != nil {
		return err
	}

	if !exists {
		// create all directories
		err = os.MkdirAll(displayTemplate.output, os.ModePerm)
		if err != nil {
			return err
		}
	}

	outputFilename := fmt.Sprintf("%s.html", name)
	outputFile := path.Join(displayTemplate.output, outputFilename)

	// create the output file
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	// close the file at the end
	defer f.Close()

	io.WriteString(f, buffer.String())

	return err
}
