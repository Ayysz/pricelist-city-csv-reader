package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CityPriceRange struct {
	ProvinceID   int
	CodeCityName string
	Price        int
	CityName     string
}

type ShoppingRecord struct {
	Vegetable string
	Fruit     string
}

func createPriceListCity(data [][]string) (CityPriceRangeList []CityPriceRange) {
	for i, line := range data {
		if i > 0 { // omit header line
			var rec CityPriceRange
			for j, field := range line {
				field = strings.TrimSpace(field)
				if j == 2 {
					rec.CodeCityName = field
				} else if j == 3 {
					rec.ProvinceID, _ = strconv.Atoi(field)
				} else if j == 5 {
					rec.CityName = field
				} else if j == 10 {
					rec.Price, _ = strconv.Atoi(strings.Replace(field, ",", "", -1))
				}
			}
			CityPriceRangeList = append(CityPriceRangeList, rec)
		}
	}
	return CityPriceRangeList
}
func main() {
	// Check if a filename is provided as an argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a CSV filename as an argument")
	}

	filename := os.Args[1]

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", filename)
	}

	// Open file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Remember to close the file at the end of the program
	defer f.Close()

	// Read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert records to array of structs
	priceListCity := createPriceListCity(data)

	// Print the array
	fmt.Printf("%+v\n", priceListCity)
}
