package http

import (
	"fmt"
	"net/http"
	"strconv" // The Translator

	"github.com/lsvishaal/unit-converter/internal/converter"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	
	// STEP 1: Grab the text from the web address
	// Example: ?value=10&from=kilometer&to=mile
	var valueText string = r.URL.Query().Get("value")
	var fromUnit string = r.URL.Query().Get("from")
	var toUnit string = r.URL.Query().Get("to")

	// STEP 2: Translate the text into a real number
	var realNumber float64
	var err error
	
	// Attempt the translation
	realNumber, err = strconv.ParseFloat(valueText, 64)
	
	// If they typed "?value=apple", stop them
	if err != nil {
		fmt.Fprint(w, "Stop: Please type a real number.")
		return 
	}

	// STEP 3: Do the math
	var finalAnswer float64
	var mathErr error
	
	finalAnswer, mathErr = converter.ConvertLength(realNumber, fromUnit, toUnit)

	// If they typed a bad unit, stop them
	if mathErr != nil {
		fmt.Fprint(w, "Stop: ", mathErr)
		return
	}

	// STEP 4: Put the answer on the silver tray (w) to send to the browser
	fmt.Fprintf(w, "Success! Your answer is: %f", finalAnswer)
}
