package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"gopkg.in/yaml.v3"
)

var (
	templateDataFilename = flag.String("data", "", "data filename")
	outputFilename       = flag.String("output", "", "output filename")
)

type Attribute struct {
	Name           string `yaml:"name"`
	GoName         string `yaml:"goName"`
	ExportedGoName string `yaml:"exportedGoName"`
	Type           string `yaml:"type"`
	Default        string `yaml:"default"`
}

type Element struct {
	Name               string      `yaml:"name"`
	GoName             string      `yaml:"goName"`
	GoType             string      `yaml:"goType"`
	Article            string      `yaml:"article"`
	Container          bool        `yaml:"container"`
	AttributeGroups    []string    `yaml:"attributeGroups"`
	Attributes         []Attribute `yaml:"attributes"`
	GeometryProperties []Attribute `yaml:"geometryProperties"`
}

type TemplateData struct {
	AttributeGroups map[string][]Attribute `yaml:"attributeGroups"`
	Elements        []Element              `yaml:"elements"`
}

func (p *Attribute) generateGoName() {
	if p.GoName != "" {
		return
	}
	components := strings.Split(p.Name, "-")
	goComponents := make([]string, 0, len(components))
	for i, component := range components {
		if i != 0 {
			component = titleize(component)
		}
		goComponents = append(goComponents, component)
	}
	p.GoName = strings.Join(goComponents, "")
}

func (p *Attribute) generateExportedGoName() {
	if p.ExportedGoName != "" {
		return
	}
	p.ExportedGoName = titleize(p.GoName)
	/*
		components := strings.Split(p.Name, "-")
		goComponents := make([]string, 0, len(components))
		for _, component := range components {
			goComponent := titleize(component)
			goComponents = append(goComponents, goComponent)
		}
		p.ExportedGoName = strings.Join(goComponents, "")
	*/
}

func (p *Attribute) generateGoType(defaultType string) {
	if p.Type != "" {
		return
	}
	p.Type = defaultType
}

func titleize(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToTitle(runes[0])
	}
	return string(runes)
}

func untitleize(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}
	return string(runes)
}

func run() error {
	flag.Parse()

	var templateData TemplateData
	if *templateDataFilename != "" {
		dataBytes, err := os.ReadFile(*templateDataFilename)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(dataBytes, &templateData); err != nil {
			return err
		}
	}

	for _, attributes := range templateData.AttributeGroups {
		for i := range attributes {
			attribute := &attributes[i]
			attribute.generateGoName()
			attribute.generateExportedGoName()
			attribute.generateGoType("String")
		}
	}

	for i := range templateData.Elements {
		element := &templateData.Elements[i]
		if element.GoName == "" {
			element.GoName = titleize(element.Name)
		}
		if element.GoType == "" {
			element.GoType = element.GoName + "Element"
		}
		for j := range templateData.Elements[i].Attributes {
			attribute := &element.Attributes[j]
			attribute.generateGoName()
			attribute.generateExportedGoName()
			attribute.generateGoType("String")
		}
		for j := range templateData.Elements[i].GeometryProperties {
			geometryProperty := &element.GeometryProperties[j]
			geometryProperty.generateGoName()
			geometryProperty.generateExportedGoName()
			geometryProperty.generateGoType("Length")
		}
	}

	if flag.NArg() == 0 {
		return errors.New("no arguments")
	}

	templateName := path.Base(flag.Arg(0))
	funcMap := template.FuncMap{
		"allAttributes": func(e Element) []Attribute {
			var allAttributes []Attribute
			for _, attributeGroupName := range e.AttributeGroups {
				allAttributes = append(allAttributes, templateData.AttributeGroups[attributeGroupName]...)
			}
			allAttributes = append(allAttributes, e.Attributes...)
			allAttributes = append(allAttributes, e.GeometryProperties...)
			return allAttributes
		},
		"default": func(defaultValue, value string) string {
			if value != "" {
				return value
			}
			return defaultValue
		},
		"quote":      strconv.Quote,
		"titleize":   titleize,
		"untitleize": untitleize,
	}
	tmpl, err := template.New(templateName).Funcs(funcMap).ParseFiles(flag.Args()...)
	if err != nil {
		return err
	}

	buffer := &bytes.Buffer{}
	if err := tmpl.Execute(buffer, templateData); err != nil {
		return err
	}

	output, err := format.Source(buffer.Bytes())
	if err != nil {
		output = buffer.Bytes()
	}

	if *outputFilename == "" {
		if _, err := os.Stdout.Write(output); err != nil {
			return err
		}
	} else if data, err := os.ReadFile(*outputFilename); err != nil || !bytes.Equal(data, output) {
		if err := os.WriteFile(*outputFilename, output, 0o666); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
