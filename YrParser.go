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
	resp, err := http.Get("http://www.yr.no/sted/Norge/Hordaland/Bergen/Bergen/varsel_time_for_time.xml")
	if err != nil {
		fmt.Println("Error reading from yr.no")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xmlReader := bytes.NewReader(body)
	weatherData := new(Weatherdata)
	if err := xml.NewDecoder(xmlReader).Decode(weatherData); err != nil {
		log.Panic(err.Error())
	}

	saveWeatherdata(weatherData)
}

func saveWeatherdata(x interface{}) (err error) {
	var xBytes []byte
	xBytes, err = json.MarshalIndent(x, "", " ")
	if err != nil {
		return
	}
	fmt.Println(string(xBytes))
	return
}
