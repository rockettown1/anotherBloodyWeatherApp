package darksky

import (
	"encoding/json"
	"fmt"
	"weatherApp/src/api/clients/restclient"
)

//Weather contains the decoded json data from darksky
type Weather struct {
	Lat     float64   `json:"latitutde"`
	Lng     float64   `json:"longitude"`
	Current currently `json:"currently"`
}

type currently struct {
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
}

//GetWeather calls darksky api
func GetWeather(lat float64, lng float64) <-chan Weather {

	c := make(chan Weather)

	go func() {
		defer close(c)
		apiKey := "fb807f90bf55dc3b7a7905299f3e12d6"
		url := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v", apiKey, lat, lng)
		data := restclient.Get(url)
		var weatherResponse Weather

		err := json.Unmarshal(data, &weatherResponse)
		if err != nil {
			fmt.Println("oops", err)
		}

		c <- weatherResponse
	}()

	return c
}
