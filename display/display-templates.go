package display

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"
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

	jsGlob := path.Join(templateDirectory, "*.js")
	tmpl, err = tmpl.ParseGlob(jsGlob)
	return tmpl, nil
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
	}

	return funcMap
}

func (displayTemplate *DisplayTemplate) Generate(context *Context, name string) error {
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

	// execute and output the template
	err = displayTemplate.template.Execute(f, context)

	return err
}
