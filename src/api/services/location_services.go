package services

import (
	"github.com/abiran/golang-testing/src/api/domain/locations"
	"github.com/abiran/golang-testing/src/api/providers/locations_provider"
	"github.com/abiran/golang-testing/src/api/utils/errors"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
