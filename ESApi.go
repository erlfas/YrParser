package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
)

type AvgTemperature struct {
	Aggregations struct {
		AvgTemperature struct {
			Value float64 `json:"value"`
		} `json:"avg_temperature"`
	} `json:"aggregations"`
}

func GetAvgTempByCityAsNumber(city CityQuery) float64 {
	url, query := getAvgTempByCityUrl(city)
	byteResults := doPOST(url, query)

	var avgTemp AvgTemperature
	if err := json.Unmarshal(byteResults, &avgTemp); err != nil {
		log.Panic(err.Error())
	}

	return avgTemp.Aggregations.AvgTemperature.Value
}

func GetAvgTempByCity(city CityQuery) string {
	url, query := getAvgTempByCityUrl(city)
	byteResults := doPOST(url, query)
	return string(byteResults)
}

func getAvgTempByCityUrl(city CityQuery) (string, string) {
	url := `http://localhost:9200/yr/weatherdata/_search?pretty`
	queryTemplateContent := `{ "aggs": { "avg_temperature": { "avg": { "field": "Forecast.Tabular.Time.Temperature.Value" } } }, "query": { "bool": { "must": [ { "match": { "Location.Name": "{{.City}}" } } ] } } }`
	queryTemplate, err := template.New("query").Parse(queryTemplateContent)
	if err != nil {
		log.Println(err.Error())
		return "", ""
	}
	var buf bytes.Buffer
	err = queryTemplate.Execute(&buf, city)
	if err != nil {
		log.Println(err.Error())
		return "", ""
	}
	query := buf.String()

	return url, query
}

func getWeatherdataByIDAsWeatherdata(id string) *WeatherdataJSON {
	var buffer bytes.Buffer
	buffer.WriteString("http://localhost:9200/yr/weatherdata/")
	buffer.WriteString(id)
	url := buffer.String()

	byteResults := doGET(url)

	var result WeatherdataJSON
	if err := json.Unmarshal(byteResults, &result); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &result
}

func getWeatherdataByID(id string) *SingularQueryResult {
	var buffer bytes.Buffer
	buffer.WriteString("http://localhost:9200/yr/weatherdata/")
	buffer.WriteString(id)
	url := buffer.String()

	byteResults := doGET(url)

	var queryResult SingularQueryResult
	if err := json.Unmarshal(byteResults, &queryResult); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &queryResult
}

func updateWeatherdataByID(id string, weatherdata *WeatherdataSource) *CRUDResult {
	var buffer bytes.Buffer
	buffer.WriteString("http://localhost:9200/yr/weatherdata/")
	buffer.WriteString(id)
	url := buffer.String()

	byteResults := doUPDATE(url, weatherdata)

	var updateResult CRUDResult
	if err := json.Unmarshal(byteResults, &updateResult); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &updateResult
}

func deleteWeatherdataByID(id string) *CRUDResult {
	var buffer bytes.Buffer
	buffer.WriteString("http://localhost:9200/yr/weatherdata/")
	buffer.WriteString(id)
	url := buffer.String()

	byteResults := doDELETE(url)

	var deleteResult CRUDResult
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
		log.Println(err.Error())
		return "", ""
	}
	var buf bytes.Buffer
	err = queryTemplate.Execute(&buf, city)
	if err != nil {
		log.Println(err.Error())
		return "", ""
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
