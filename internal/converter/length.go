package converter

import (
	"errors"

)

// unitToMeter stores how many meters are in each supported unit.
// Example:
// 1 kilometer = 1000 meters
// 1 inch = 0.0254 meters

var unitToMeter = map[string]float64{
	"millimeter": 0.001,
	"centimeter": 0.01,
	"meter":      1,
	"kilometer":  1000,
	"inch":       0.0254,
	"foot":       0.3048,
	"yard":       0.9144,
	"mile":       1609.34,
}


// ConvertLength converts a value from one length unit to another.
//
// Parameters:
//   - value: the numeric value to convert
//   - from:  the source unit (e.g. "kilometer")
//   - to:    the target unit (e.g. "mile")
//
// Example:
//   ConvertLength(10, "kilometer", "mile") ≈ 6.2137
//
// Conversion strategy:
// 1. Convert the input value to meters
// 2. Convert meters to the target unit
func ConvertLength(value float64, from string, to string) (float64, error) {
	

	//1. retreive the "from" rate
	var fromRate float64
	var fromExists bool
	fromRate, fromExists = unitToMeter[from]

	//if typo, stop and return error
	if fromExists == false {
		return 0.0, errors.New("Source unit NOT FOUND, please try again.")
	}

	//2. Retrieve the "to" rate
	var toRate float64
	var toExists bool
	toRate, toExists = unitToMeter[to]

	// if unit was typo, stop and return error
	if toExists == false{
		return 0.0, errors.New("target unit was not found, please try again")
	}

	//3. else(no errors), we Execute the Math fn,
	var meters float64
	meters = value * fromRate
	
	var result float64
	result = meters / toRate

	//return final number and "nil" to signal no errors occured.

	return result, nil
}

