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