package main

import (
	"fmt"
	"time"
	"weatherApp/app"
	"weatherApp/src/utils"
)

func main() {
	utils.MoveTopLeft()
	utils.Clear()
	fmt.Println("Starting Weather app...")
	time.Sleep(time.Second)
	app.StartApp()

}
