package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//load properties file
	properties_input, err := os.Open("./properties.txt")
	if err != nil {
		fmt.Println(err)
	}

	// setup reader
	reader := csv.NewReader(properties_input)
	reader.Comma = '\t' // our list is delimited by tabs, not commas. alternatively surround addressess in the file with ""

	// read file into nested slices
	properties_list, err := reader.ReadAll()
	if err !=nil {
		fmt.Println(err)
	}

	// iterate through properties list
	for i := range properties_list {
		// remove blank entries (we'll keep first row headers as a sort of index)
		if properties_list[i][1] == "" {
			continue
		}

		fmt.Println(properties_list[i])
	}
}
