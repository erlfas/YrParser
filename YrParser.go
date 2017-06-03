package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Weatherdata struct {
	Location struct {
		Name     string `xml:"name"`
		Type     string `xml:"type"`
		Country  string `xml:"country"`
		Timezone struct {
			Id               string `xml:"id,attr"`
			UtcoffsetMinutes string `xml:"utcoffsetMinutes"`
		} `xml:"timezone"`
		Location struct {
			Altitude  int64   `xml:"altitude,attr"`
			Latitude  float64 `xml:"latitude,attr"`
			Longitude float64 `xml:"longitude,attr"`
			Geobase   string  `xml:"geobase,attr"`
			Geobaseid int64   `xml:"geobaseid,attr"`
		} `xml:"location"`
	} `xml:"location"`
	Meta struct {
		Lastupdate string `xml:"lastupdate"`
		Nextupdate string `xml:"nextupdate"`
	} `xml:"meta"`
	Forecast struct {
		Tabular struct {
			Time []struct {
				From   string `xml:"from,attr"`
				To     string `xml:"to,attr"`
				Symbol struct {
					Number   int64  `xml:"number,attr"`
					NumberEx int64  `xml:"numberEx,attr"`
					Name     string `xml:"name,attr"`
					Var      string `xml:"var,attr"`
				} `xml:"symbol"`
				Precipitation struct {
					Value    float64 `xml:"value,attr"`
					Minvalue float64 `xml:"minvalue,attr"`
					Maxvalue float64 `xml:"maxvalue,attr"`
				} `xml:"precipitation"`
				WindDirection struct {
					Deg  float64 `xml:"deg,attr"`
					Code string  `xml:"code,attr"`
					Name string  `xml:"name,attr"`
				} `xml:"windDirection"`
				WindSpeed struct {
					Mps  float64 `xml:"mps,attr"`
					Name string  `xml:"name,attr"`
				} `xml:"windSpeed"`
				Temperature struct {
					Unit  string `xml:"unit,attr"`
					Value int64  `xml:"value,attr"`
				} `xml:"temperature"`
				Pressure struct {
					Unit  string  `xml:"unit,attr"`
					Value float64 `xml:"value,attr"`
				} `xml:"pressure"`
			} `xml:"time"`
		} `xml:"tabular"`
	} `xml:"forecast"`
}

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
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("Error reading from ")
		fmt.Println(url)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xmlReader := bytes.NewReader(body)
	weatherData := new(Weatherdata)
	if err := xml.NewDecoder(xmlReader).Decode(weatherData); err != nil {
		log.Panic(err.Error())
	}

	return weatherData
}

func downloadAndSave(url string, done chan bool) {
	weatherData := getWeatherdata(url)

	saveWeatherdata(weatherData)

	done <- true
}

func saveWeatherdata(x interface{}) (err error) {
	var xBytes []byte
	xBytes, err = json.MarshalIndent(x, "", " ")
	if err != nil {
		return
	}
	http.Post("http://localhost:9200/yr/weatherdata", "application/json", bytes.NewBuffer(xBytes))
	return
}
