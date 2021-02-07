package test

import (
	"fmt"
	"github.com/abiran/golang-testing/src/api/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start application")
	go app.StartApp()
	fmt.Println("application started, about to start test cases")
	os.Exit(m.Run())
}
