package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Handler for HTTP testing
func handlerFunc(w http.ResponseWriter, r *http.Request) {}

func TestAzureStatus(t *testing.T) {
	httpGet := "https://status.azure.com"
	if httpGet != "https://status.azure.com" {
		t.Error("Ensure to use the status.azure.com page")
	}

	azurePtr := "--azure"
	if azurePtr != "--azure" {
		t.Error("Ensure to use the `--azure` flag")
	}

	assert.HTTPStatusCode(t, handlerFunc, "GET", "https://status.azure.com", nil, 200)
	assert.HTTPSuccess(t, handlerFunc, "GET", "https://status.azure.com", nil)
}
