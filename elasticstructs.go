package main

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

type Hit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source struct {
		Location struct {
			Name     string `json:"name"`
			Type     string `json:"type"`
			Country  string `json:"country"`
			Timezone struct {
				Id               string `json:"id"`
				UtcoffsetMinutes string `json:"utcoffsetMinutes"`
			} `xml:"timezone"`
			Location struct {
				Altitude  int64   `json:"altitude"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Geobase   string  `json:"geobase"`
				Geobaseid int64   `json:"geobaseid"`
			} `json:"location"`
		} `json:"location"`
		Meta struct {
			Lastupdate string `json:"lastupdate"`
			Nextupdate string `json:"nextupdate"`
		} `xml:"meta"`
		Forecast struct {
			Tabular struct {
				Time []struct {
					From   string `json:"from"`
					To     string `json:"to"`
					Symbol struct {
						Number   int64  `json:"number"`
						NumberEx int64  `json:"numberEx"`
						Name     string `json:"name"`
						Var      string `json:"var"`
					} `json:"symbol"`
					Precipitation struct {
						Value    float64 `json:"value"`
						Minvalue float64 `json:"minvalue"`
						Maxvalue float64 `json:"maxvalue"`
					} `json:"precipitation"`
					WindDirection struct {
						Deg  float64 `json:"deg"`
						Code string  `json:"code"`
						Name string  `json:"name"`
					} `json:"windDirection"`
					WindSpeed struct {
						Mps  float64 `json:"mps"`
						Name string  `json:"name"`
					} `json:"windSpeed"`
					Temperature struct {
						Unit  string `json:"unit"`
						Value int64  `json:"value"`
					} `json:"temperature"`
					Pressure struct {
						Unit  string  `jsonxml:"unit"`
						Value float64 `json:"value"`
					} `json:"pressure"`
				} `json:"time"`
			} `json:"tabular"`
		} `json:"forecast"`
	} `json:"_source"`
}

type WeatherdataJSON struct {
	Hit
}

type Hits struct {
	Hits []Hit `json:"hits"`
}

type QueryResult struct {
	Hits Hits `json:"hits"`
}

type SingularQueryResult struct {
	Hit
}

type CRUDResult struct {
	Found   bool   `json:"found"`
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int64  `json:"_version"`
	Created bool   `json:"created"`
}
