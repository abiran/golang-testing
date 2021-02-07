package test

import (
	"encoding/json"
	"fmt"
	"github.com/abiran/golang-testing/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCountriesNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/US",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"status": 404, "error": "not_found", "message": "no country with id US"}`,
	})
	response, err := http.Get("http://localhost:8080/locations/countries/US")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	var apiError errors.ApiError
	err = json.Unmarshal(bytes, &apiError)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "not_found", apiError.Error)
	assert.EqualValues(t, "no country with id US", apiError.Message)
}
