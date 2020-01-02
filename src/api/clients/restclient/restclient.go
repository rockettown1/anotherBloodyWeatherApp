package restclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Get calls an api
func Get(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The API call has failed with error %v\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return data

}
