package locations_provider

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetCountryRestClientError(t *testing.T) {
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get country US", err.Message)
}

func TestGetCountryCountryNotFound(t *testing.T) {
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country US", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for US", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	country, err := GetCountry("US")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "US", country.Id)
	assert.EqualValues(t, "United States of America", country.Name)
	assert.EqualValues(t, "GMT-04:00", country.TimeZone)
	assert.EqualValues(t, 53, len(country.States))
}
