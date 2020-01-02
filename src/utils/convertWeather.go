package utils

import (
	"math"
)

//FtoC takes a farenheit temp and returns it in celcius
func FtoC(temp float64) float64 {
	newTemp := math.Round((temp - 32) * (5.0 / 9))
	return newTemp
}

//Note for future dan
//the division above 5/9 uses untyped constants - if they were both untyped ints the result would be an untyped int with a value of 0. As we want the result of the division to be a floating-point constant, we make one of the operands an untyped float so the result is a float. Just be careful with types basically. See you in the future.
