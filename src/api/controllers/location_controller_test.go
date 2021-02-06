package controllers

import (
	"encoding/json"
	"github.com/abiran/golang-testing/src/api/domain/locations"
	"github.com/abiran/golang-testing/src/api/services"
	"github.com/abiran/golang-testing/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationsServiceMock struct {
}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
	// mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "US"},
	}
	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)
	var apiErr errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)

}
