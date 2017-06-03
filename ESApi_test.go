package main

import "testing"

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
