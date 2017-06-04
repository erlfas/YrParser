package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestDeleteWeatherdataByID(t *testing.T) {
	var id string = "AVxtldwDW33Zbbv0nvtu"
	var result *DeleteResult = deleteWeatherdataByID(id)

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
