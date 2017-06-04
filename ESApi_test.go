package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type AvgTemperature struct {
	Aggregations struct {
		AvgTemperature struct {
			Value float64 `json:"value"`
		} `json:"avg_temperature"`
	} `json:"aggregations"`
}

func TestGetAvgTempByCity(t *testing.T) {
	jsonResult := GetAvgTempByCity(CityQuery{"Bergen"})
	//fmt.Println(jsonResult)

	var avgTemp AvgTemperature
	if err := json.Unmarshal([]byte(jsonResult), &avgTemp); err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(avgTemp)

	if avgTemp.Aggregations.AvgTemperature.Value == 0 {
		t.Error("Go 0 as avg temp")
	}
}

func TestGetWeatherdataByIDAsWeatherdata(t *testing.T) {
	id := "AVxthxdbW33Zbbv0nvtm"
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

	result.Source.Location.Type = "asdf"

	prevVersion := result.Hit.Version

	updateResult := updateWeatherdataByID(id, &result.Hit.Source)

	newVersion := updateResult.Version

	if prevVersion == newVersion {
		t.Error("Expected incremented version number")
	}

	//jsonUpdate, _ := json.Marshal(updateResult)
	//fmt.Println(string(jsonUpdate))

	if updateResult == nil {
		t.Error("No update result")
	}

	if updateResult.Created != false {
		t.Error("Expected created to be false because it should already exist")
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
}

func TestGetWeatherdataByID(t *testing.T) {
	id := "AVxthxdbW33Zbbv0nvtm"
	queryResult := getWeatherdataByID(id)

	//fmt.Println(queryResult)

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
	var id string = "nonExistentID"
	var result *CRUDResult = deleteWeatherdataByID(id)

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

	//fmt.Println(queryResult)
}

func TestDoFindAllByCity(t *testing.T) {
	result := doFindAllByCity(CityQuery{"Bergen"})
	//fmt.Println(result)

	var queryResult QueryResult
	if err := json.Unmarshal([]byte(result), &queryResult); err != nil {
		log.Panic(err.Error())
	}

	//fmt.Println(queryResult)
}
