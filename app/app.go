package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"weatherApp/src/api/darksky"
	"weatherApp/src/api/mapbox"
	"weatherApp/src/utils"
)

//StartApp fires the app
func StartApp() {
	check := "yes"

	for check == "yes" {
		//execute weather fetch
		check = weatherLogic()

		if check != "yes" {
			utils.ScreenClear("Bye for now")
			time.Sleep(time.Second)
			utils.Clear()
			utils.MoveTopLeft()
		}

	}

}

func weatherLogic() string {
	utils.ScreenClear("Weather App")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Where would you like the weather for?")
	scanner.Scan()
	location := scanner.Text()
	//wait for the rl channel to recieve the data from the getLocation api call (this is blocking until it's recieved it)
	rl := <-mapbox.GetLocation(location)

	if len(rl.Features) == 0 {
		fmt.Println("Your search query returned no results! Please try again")
		time.Sleep(time.Second * 2)
		return "yes"
	}

	lat, lng := rl.Features[0].Center[1], rl.Features[0].Center[0]

	//wait for the rw channel to recieve the result from the getWeather api call
	rw := <-darksky.GetWeather(lat, lng)

	celcius := utils.FtoC(rw.Current.Temperature)
	utils.ScreenClear("Weather App")
	fmt.Println()
	fmt.Printf("The current temperature in %v is %v°C and %v\n\n", rl.Features[0].Place, celcius, rw.Current.Summary)

	//check for continue
	fmt.Println("Would you like to search again? yes/no")
	scanner.Scan()
	finished := strings.ToLower(scanner.Text())
	return finished

}
