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
	for i, v := range properties_list {
		// remove blank entries (we'll keep first row headers as a sort of index)
		if properties_list[i][1] == "" {
			continue
		}
		// remove an entry if an entry with matching address and date exists earlier in the slice
		if sliceContainsBefore(properties_list, i, v) {
			continue
		}

		// print the entry if it has passed filters
		fmt.Println(properties_list[i])
	}
}


func sliceContainsBefore(slice [][]string, before_index int, compare_value []string) bool {
	for _, v := range slice[:before_index] {
		if compare_value[1] == v[1] && compare_value[2] == v[2] && compare_value[3] == v[3] {
			return true
		}
	}
	return false
}
