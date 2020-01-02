package mapbox

import (
	"encoding/json"
	"fmt"
	"net/url"
	"weatherApp/src/api/clients/restclient"
)

//Location defines the response from the mapbox api
type Location struct {
	Features []Locations
	Message  string
}

//Locations represents a slice of the locations that return from an api call
type Locations struct {
	Place  string `json:"place_name"`
	Center []float64
}

//GetLocation returns the api data from mapbox
func GetLocation(query string) <-chan Location {

	c := make(chan Location)

	go func() {
		defer close(c)
		apiKey := "pk.eyJ1Ijoicm9ja2V0dG93biIsImEiOiJjanQ2ZmEyZnowZjloNDRtd2VtemR3dzZmIn0.JLgxwoeoCASsZ8WDYI3-5A"

		locationString := url.PathEscape(query)

		locationURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%v.json?access_token=%v", locationString, apiKey)

		data := restclient.Get(locationURL)
		var locationResponse Location

		err := json.Unmarshal(data, &locationResponse)
		if err != nil {
			fmt.Println("Oops there has been an error", err)

		}

		c <- locationResponse
	}()

	return c
}
