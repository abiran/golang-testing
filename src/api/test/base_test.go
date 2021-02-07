package test

import (
	"github.com/abiran/golang-testing/src/api/app"
	"github.com/mercadolibre/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
