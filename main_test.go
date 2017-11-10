package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	record := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	handler.ServeHTTP(record, request)

	assert.Equal(t, record.Code, http.StatusOK)
	assert.Equal(t, record.Body.String(), "Hello World.")
}
