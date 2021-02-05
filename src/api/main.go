package main

import (
	"fmt"
	"github.com/abiran/golang-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("US")
	fmt.Println(err)
	fmt.Println(country)
}
