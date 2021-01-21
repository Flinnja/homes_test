package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"strconv"
	"strings"
)

type record struct {
	id string
	address string
	town string
	date string
	price string
}

func main() {
	// load properties file
	properties_input, err := os.Open("./properties.txt")
	if err != nil {
		fmt.Println(err)
	}

	// declare properties slice
	properties := make([]record, 0)
	// declare map to check duplicates
	duplicates := make(map[string]int)

	// setup reader
	reader := csv.NewReader(properties_input)
	reader.Comma = '\t' // our list is delimited by tabs, not commas. alternatively surround addressess in the file with ""

	// read file into a slice of records
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		// skip the header row and empty rows
		if row[0] == "ID" || row[0] == "" {
			continue
		}
		var newProperty = record{row[0], row[1], row[2], row[3], row[4]}

		// count duplicates
		var duplicate_key = row[1]+row[2]+row[3]
		if v, found := duplicates[duplicate_key]; found {
			v++
		} else {
			duplicates[duplicate_key] = 0
		}

		// append to properties slice
		properties = append(properties, newProperty)
	}

	// iterate through properties list
	for i, v := range properties {
		// remove every tenth entry
		if i%10 == 9 {
			continue
		}

		// remove properties that are valuated below 400k
		price, err := strconv.Atoi(v.price)
		if err != nil {
			fmt.Println(err)
		}
		if price < 400000 {
			continue
		}

		// remove properties with AVE, CRES, or PL in street name
		if strings.Contains(v.address, "AVE") || strings.Contains(v.address, "CRES") || strings.Contains(v.address, "PL") {
			continue
		}

		// remove duplicate entries
		if duplicates[v.address + v.town + v.date] > 0 {
			continue
		}

		// print anything that passed the filters
		fmt.Println(v)
	}
}
