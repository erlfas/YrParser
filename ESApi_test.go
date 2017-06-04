package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetWeatherdataByIDAsWeatherdata(t *testing.T) {
	id := "AVxtAEeRW33Zbbv0nvti"
	result := getWeatherdataByIDAsWeatherdata(id)

	//jsonString, _ := json.Marshal(result)
	//fmt.Println(string(jsonString))

	if result == nil {
		t.Error("Got no weatherdata")
	}

	if result.Id != id {
		t.Error("Unexpectd id")
	}

	if result.Index != "yr" {
		t.Error("Unexpectd index")
	}

	if result.Type != "weatherdata" {
		t.Error("Unexpectd type")
	}

	if result.Source.Location.Name != "Bergen" {
		t.Error("Unexpectd location name")
	}

	if result.Source.Location.Type != "By" {
		t.Error("Type was not updated from By to Storby")
	}

	result.Source.Location.Type = "Storby"

	//jsonString, _ := json.Marshal(result)
	//fmt.Println(string(jsonString))

	updateResult := updateWeatherdataByID(id, result)

	jsonUpdate, _ := json.Marshal(updateResult)
	fmt.Println(string(jsonUpdate))

	if updateResult == nil {
		t.Error("No update result")
	}

	if updateResult.Created != false {
		t.Error("Expected created to be false because it should already exist")
	}

	if updateResult.Found == false {
		t.Error("Excpected found to be true")
	}

	result2 := getWeatherdataByIDAsWeatherdata(id)

	//jsonString3, _ := json.Marshal(result2)
	//fmt.Println(string(jsonString3))

	if result2 == nil {
		t.Error("Got no weatherdata")
	}

	if result2.Id != id {
		t.Error("Unexpectd id")
	}

	if result2.Index != "yr" {
		t.Error("Unexpectd index")
	}

	if result2.Type != "weatherdata" {
		t.Error("Unexpectd type")
	}

	if result2.Source.Location.Name != "Bergen" {
		t.Error("Unexpectd location name")
	}

	if result2.Source.Location.Type != "Storby" {
		t.Error("Type was not updated from By to Storby")
	}
}

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
