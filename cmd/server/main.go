package main

import {
	"fmt"
	"github.com/lsvishaal/unit-converter/internal/converter"
}

func main(){
	result := converter.ConvertLength(10, "kilometer", "mile")
	fmt.Println(result)
}