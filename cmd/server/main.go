package main

import (
	"fmt"
	"net/http"

	httphandler "github.com/lsvishaal/unit-converter/internal/http"
)

func main(){
	// 1. on User visit to /convert URL, redirected to  ConvertHandler 
	http.HandleFunc("/convert", httphandler.ConvertHandler)

	//2. Start the infininte loop on port 8080
	
	var err error
	fmt.Println("Server is starting on http://localhost:8080...")

	//ListenAndServe locks program in infinite loop
	//returns error only if server crashes fatally
	err = http.ListenAndServe(":8080", nil)

	// 3. If the loop breaks, catch the error
	if err != nil {
		fmt.Println("The server crashed:", err)
	}
}