package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetWeatherdataByID(t *testing.T) {
	id := "AVxtAEeRW33Zbbv0nvti"
	queryResult := getWeatherdataByID(id)

	fmt.Println(queryResult)

	if queryResult == nil {
		t.Error("Got no weatherdata")
	}

	if queryResult.Id != id {
		t.Error("Unexpectd id")
	}

	if queryResult.Index != "yr" {
		t.Error("Unexpectd index")
	}

	if queryResult.Type != "weatherdata" {
		t.Error("Unexpectd type")
	}

	if queryResult.Source.Location.Name != "Bergen" {
		t.Error("Unexpectd location name")
	}
}

func TestDeleteWeatherdataByID(t *testing.T) {
	var id string = "AVxtldwDW33Zbbv0nvtu"
	var result *CRUDResult = deleteWeatherdataByID(id)

	fmt.Println(*result)

	if result.ID != id {
		t.Error("Returned id is different from input id")
	}

	if result.Index != "yr" {
		t.Error("Unexpected index")
	}

	if result.Type != "weatherdata" {
		t.Error("Unexpected type")
	}

	if result.Found == true {
		t.Error("Did not expect to delete a document")
	}
}
func TestFindAllByCity(t *testing.T) {
	url, query := findAllByCity(CityQuery{"Bergen"})
	if url != "http://localhost:9200/yr/weatherdata/_search?pretty" {
		t.Error("url is not as expected")
	}
	expectedQuery := `{ "query": { "bool": { "must": [ { "match": { "Location.Name": "Bergen" } } ] } } }`
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
	}

	fmt.Println(queryResult)
}

func TestDoFindAllByCity(t *testing.T) {
	result := doFindAllByCity(CityQuery{"Bergen"})
	//fmt.Println(result)

	var queryResult QueryResult
	if err := json.Unmarshal([]byte(result), &queryResult); err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(queryResult)
}
