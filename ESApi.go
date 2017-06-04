package main

import (
	"bytes"
	"html/template"
	"log"
)

type CityQuery struct {
	City string
}

func findAllByCity2(city CityQuery) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://localhost:9200/yr/weatherdata/_search?q=")
	buffer.WriteString(city.City)
	url := buffer.String()
	return url
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
		log.Panic(err.Error())
		panic(err)
	}
	var buf bytes.Buffer
	err = queryTemplate.Execute(&buf, city)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}
	query := buf.String()

	return url, query
}

func doFindAllByCity2(city CityQuery) string {
	url := findAllByCity2(city)
	byteResults := doPOSTWithBytes(url, []byte(""))
	return string(byteResults)
}