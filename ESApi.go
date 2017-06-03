package main

import (
	"bytes"
	"html/template"
)

type CityQuery struct {
	City string
}

func findAllByCity(city CityQuery) (string, string) {
	url := `http://localhost:9200/yr/weatherdata/_search?pretty`
	queryTemplateContent := `
	{
		"query": {
			"match": {
				"name": "{{.City}}"
			}
		}
	}`
	queryTemplate, err := template.New("query").Parse(queryTemplateContent)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = queryTemplate.Execute(&buf, city)
	if err != nil {
		panic(err)
	}
	query := buf.String()

	return url, query
}
