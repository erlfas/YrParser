# YrParser

Parses xml files from yr.no and loads them into an Elasticsearch instance.

## Communicating with Elasticsearch

### Getting all weatherdata

```
curl -XGET 'http://localhost:9200/yr/weatherdata/_search'
```
### Counting weatherdata

```
curl -XGET 'http://localhost:9200/yr/weatherdata/_count?pretty'
```

### Search for weatherdata by location name

With curl:

```
curl -XGET 'localhost:9200/_search?pretty' -H 'Content-Type: application/json' -d'{ "query": { "bool": { "must": [ { "match": { "Location.Name": "Bergen" } } ] } } }'
```
Pretty:

```
GET localhost:9200/_search?pretty 
{
	"query": {
		"bool": {
			"must": [
				{
					"match": {
						"Location.Name": "Bergen"
					}
				}
			]
		}
	}
}
```

### Find average temperature by city

With curl:

```
curl -XGET 'localhost:9200/yr/weatherdata/_search?pretty' -H 'Content-Type: application/json' -d'{ "aggs": { "avg_temperature": { "avg": { "field": "Forecast.Tabular.Time.Temperature.Value" } } }, "query": { "bool": { "must": [ { "match": { "Location.Name": "Bergen" } } ] } } }' | less
```

Pretty:

 ```
 GET localhost:9200/yr/weatherdata/_search?pretty
 {
	"aggs": {
		"avg_temperature": {
			"avg": {
				"field": "Forecast.Tabular.Time.Temperature.Value"
			}
		}
	},
	"query": {
		"bool": {
			"must": [ 
				{
					"match": {
						"Location.Name": "Bergen"
					}
				}
			]
		}
	}
}
 ```