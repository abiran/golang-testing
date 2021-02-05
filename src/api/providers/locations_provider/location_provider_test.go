package locations_provider

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryRestClientError(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/US",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get country US", err.Message)
}

func TestGetCountryCountryNotFound(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/US",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found","error": "not_found","status": 404,"cause": []}`,
	})
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/US",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found","error": "not_found","status": "404"","cause": []}`,
	})
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country US", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/US",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 123,"name": "United States of America","time_zone": "GMT-04:00"}`,
	})
	country, err := GetCountry("US")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for US", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/US",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"US","name":"United States of America","locale":"en_US","currency_id":"USD","decimal_separator":".","thousands_separator":",","time_zone":"GMT-04:00","geo_information":{"location":{"latitude":38.8867036,"longitude":-77.0081972}},"states":[{"id":"US-AL","name":"Alabama"},{"id":"US-AK","name":"Alaska"},{"id":"US-AZ","name":"Arizona"},{"id":"US-AR","name":"Arkansas"},{"id":"US-CA","name":"California"},{"id":"US-CO","name":"Colorado"},{"id":"US-CT","name":"Connecticut"},{"id":"US-DE","name":"Delaware"},{"id":"US-DC","name":"District of Columbia"},{"id":"US-FL","name":"Florida"},{"id":"US-GA","name":"Georgia"},{"id":"US-HI","name":"Hawaii"},{"id":"US-ID","name":"Idaho"},{"id":"US-IL","name":"Illinois"},{"id":"US-IN","name":"Indiana"},{"id":"US-IA","name":"Iowa"},{"id":"US-KS","name":"Kansas"},{"id":"US-KY","name":"Kentucky"},{"id":"US-LA","name":"Louisiana"},{"id":"US-ME","name":"Maine"},{"id":"US-MD","name":"Maryland"},{"id":"US-MA","name":"Massachusetts"},{"id":"US-MI","name":"Michigan"},{"id":"US-MN","name":"Minnesota"},{"id":"US-MS","name":"Mississippi"},{"id":"US-MO","name":"Missouri"},{"id":"US-MT","name":"Montana"},{"id":"US-NE","name":"Nebraska"},{"id":"US-NV","name":"Nevada"},{"id":"US-NH","name":"New Hampshire"},{"id":"US-NJ","name":"New Jersey"},{"id":"US-NM","name":"New Mexico"},{"id":"US-NY","name":"New York"},{"id":"US-NC","name":"North Carolina"},{"id":"US-ND","name":"North Dakota"},{"id":"US-OH","name":"Ohio"},{"id":"US-OK","name":"Oklahoma"},{"id":"US-OR","name":"Oregon"},{"id":"US-PA","name":"Pennsylvania"},{"id":"US-PR","name":"Puerto Rico"},{"id":"US-RI","name":"Rhode Island"},{"id":"US-SC","name":"South Carolina"},{"id":"US-SD","name":"South Dakota"},{"id":"US-TN","name":"Tennessee"},{"id":"US-TX","name":"Texas"},{"id":"US-UT","name":"Utah"},{"id":"US-VT","name":"Vermont"},{"id":"US-VI","name":"Virgin Islands"},{"id":"US-VA","name":"Virginia"},{"id":"US-WA","name":"Washington"},{"id":"US-WV","name":"West Virginia"},{"id":"US-WI","name":"Wisconsin"},{"id":"US-WY","name":"Wyoming"}]}`,
	})
	country, err := GetCountry("US")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "US", country.Id)
	assert.EqualValues(t, "United States of America", country.Name)
	assert.EqualValues(t, "GMT-04:00", country.TimeZone)
	assert.EqualValues(t, 53, len(country.States))
}
