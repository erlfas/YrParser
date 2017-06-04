package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
)

func main() {

	doneChannel := make(chan bool, 5)

	urls := make([]string, 5)
	urls[0] = "http://www.yr.no/place/Norway/Hordaland/Bergen/Bergen/forecast.xml"
	urls[1] = "http://www.yr.no/place/Norway/M%C3%B8re_og_Romsdal/Rauma/%C3%85ndalsnes/forecast.xml"
	urls[2] = "http://www.yr.no/place/Norway/M%C3%B8re_og_Romsdal/Molde/Molde/forecast.xml"
	urls[3] = "http://www.yr.no/place/Norway/M%C3%B8re_og_Romsdal/Kristiansund/Kristiansund/forecast.xml"
	urls[4] = "http://www.yr.no/place/Norway/M%C3%B8re_og_Romsdal/%C3%85lesund/%C3%85lesund/forecast.xml"

	for i := 0; i < len(urls); i++ {
		go downloadAndSave(urls[i], doneChannel)
	}

	// wait for go routines to finish
	for i := 0; i < len(urls); i++ {
		<-doneChannel
	}
}

func getWeatherdata(url string) *Weatherdata {
	body := doGET(url)
	xmlReader := bytes.NewReader(body)
	weatherData := new(Weatherdata)
	if err := xml.NewDecoder(xmlReader).Decode(weatherData); err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	fmt.Println(weatherData.Location.Name)

	return weatherData
}

func downloadAndSave(url string, done chan bool) {
	weatherData := getWeatherdata(url)

	saveWeatherdata(weatherData)

	done <- true
}

func saveWeatherdata(x interface{}) (err error) {
	doPOSTX("http://localhost:9200/yr/weatherdata", x)
	return
}
