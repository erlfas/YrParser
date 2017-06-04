package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
)

func deleteWeatherdataByID(id string) *DeleteResult {
	var buffer bytes.Buffer
	buffer.WriteString("http://localhost:9200/yr/weatherdata/")
	buffer.WriteString(id)
	url := buffer.String()

	byteResults := doDELETE(url)

	var deleteResult DeleteResult
	if err := json.Unmarshal(byteResults, &deleteResult); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &deleteResult
}

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
	queryTemplateContent := `{ "query": { "bool": { "must": [ { "match": { "Location.Name": "{{.City}}" } } ] } } }`
	queryTemplate, err := template.New("query").Parse(queryTemplateContent)
	if err != nil {
		log.Panic(err.Error())
	}
	var buf bytes.Buffer
	err = queryTemplate.Execute(&buf, city)
	if err != nil {
		log.Panic(err.Error())
	}
	query := buf.String()

	return url, query
}

func doFindAllByCity(city CityQuery) string {
	url, query := findAllByCity(city)
	byteResults := doPOST(url, query)
	return string(byteResults)
}

func doFindAllByCity2(city CityQuery) string {
	url := findAllByCity2(city)
	byteResults := doPOSTWithBytes(url, []byte(""))
	return string(byteResults)
}
