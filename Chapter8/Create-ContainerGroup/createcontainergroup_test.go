// +build unit

package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/stretchr/testify/assert"
)

func TestAzureAuth(t *testing.T) {
	auth, err := auth.NewAuthorizerFromCLI()

	if err != nil {
		t.Error("There was an error authenticating with the current Azure CLI profile")
	} else {
		t.Log(auth)
	}
}

func TestPort(t *testing.T) {
	port := "8080"
	assert.Contains(t, port, "8080")
}

// Table Driven Test scenario
type StringResult struct {
	given    string
	expected string
}

func TestContainerImageString(t *testing.T) {
	image := "adminturneddevops/golangwebapi:latest"

	testString := reflect.TypeOf(image).Kind()

	if testString == reflect.String {
		log.Println("")
	} else {
		t.Error("Image is not in string form. Ensure the Docker image is written as a string and try again")
	}

	var aString = []StringResult{
		{
			"adminturneddevops/golangwebapi:latest", "adminturneddevops/golangwebapi:latest",
		},
	}

	for _, test := range aString {
		result, _ := test.given, test.expected
		if result != test.expected {
			t.Fatal("No string found")
		}
	}
}

func TestContainerPortInt(t *testing.T) {
	port := 8080

	testInt32 := reflect.TypeOf(port).Kind()

	if testInt32 == reflect.Int {
		log.Println("")
	} else {
		t.Error("Port is not of type Int. Please ensure the port is of type Int")
	}
}

func TestLocationString(t *testing.T) {
	location := "eastus"

	testStringLocation := reflect.TypeOf(location).Kind()

	if testStringLocation == reflect.String {
		log.Println("")
	} else {
		t.Error("Location is not of type string. Please ensure the location is of type string")
	}
}

func TestHttpAzureSDK(t *testing.T) {
	handle := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "status successful")
	}

	req, err := http.NewRequest("get", "https://github.com/Azure/azure-sdk-for-go", nil)
	if err != nil {
		t.Fatal("HTTP GET request to the SDK is nil")
	} else {
		w := httptest.NewRecorder()
		handle(w, req)
	}
}
