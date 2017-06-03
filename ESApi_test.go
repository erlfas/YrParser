package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestFindAllByCity(t *testing.T) {
	url, query := findAllByCity(CityQuery{"Bergen"})
	if url != "http://localhost:9200/yr/weatherdata/_search?pretty" {
		t.Error("url is not as expected")
	}
	expectedQuery := `
	{
		"query": {
			"match": {
				"name": "Bergen"
			}
		}
	}`
	if query != expectedQuery {
		t.Error("Query is not as expected")
	}
}

func TestDoFindAllByCity2(t *testing.T) {
	result := doFindAllByCity2(CityQuery{"Bergen"})
	//fmt.Println(result)

	var queryResult QueryResult
	if err := json.Unmarshal([]byte(result), &queryResult); err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	fmt.Println(queryResult)
}
