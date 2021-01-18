package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

	// set up waitgroup for goroutines
	var wg sync.WaitGroup

	// iterate through properties list
	for i, v := range properties_list {
		// always print the header row
		if i == 0{
			fmt.Println(v)
			continue
		}
		// remove blank entries
		if properties_list[i][1] == "" {
			continue
		}
		// remove every tenth entry
		if i%10 == 0{
			continue
		}
		// remove duplicate entries
		if sliceContains(properties_list, i, v) {
			continue
		}

		wg.Add(1)
		go arbitraryFilter(v, &wg)

		wg.Wait()
	}
}


func sliceContains(slice [][]string, current_index int, compare_value []string) bool {
	for i, v := range slice {
		// do not check against the literal same entry!
		if i == current_index {
			continue
		}
		if compare_value[1] == v[1] && compare_value[2] == v[2] && compare_value[3] == v[3] {
			return true
		}
	}
	return false
}

func arbitraryFilter(slice []string, wg *sync.WaitGroup) {
	defer wg.Done()

	printable := true
	price, err := strconv.Atoi(slice[4])
	if err != nil {
		fmt.Println(err)
	}
	if price < 400000 {
		printable = false
	}
	if strings.Contains(slice[1], "AVE") || strings.Contains(slice[1], "CRES") || strings.Contains(slice[1], "PL") {
		printable = false
	}
	if printable == true {
		fmt.Println(slice)
	}
}
