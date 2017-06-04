# YrParser

Parses xml files from yr.no and load them into an Elasticsearch instance.

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

```
curl -XGET 'localhost:9200/_search?pretty' -H 'Content-Type: application/json' -d'
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
}'
```