package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//load properties file
	propertiesInput, err := os.Open("./properties.txt")
	if err != nil {
		fmt.Println(err)
	}
	// setup reader
	reader := csv.NewReader(propertiesInput)
	reader.Comma = '\t' // our list is delimited by tabs, not commas. alternatively surround addressess in the file with ""

	// read file into nested slices
	propertiesList, err := reader.ReadAll()
	if err !=nil {
		fmt.Println(err)
	}

	fmt.Print(propertiesList)
}
