package utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// Clear clears the screen
func Clear() {
	fmt.Print("\033[2J")
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft() {
	fmt.Print("\033[H")

}

//ScreenClear clears the screen and reprints ascii title
func ScreenClear(message string) {
	MoveTopLeft()
	Clear()
	myFigure := figure.NewFigure(message, "", true)
	myFigure.Print()
}
