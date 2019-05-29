package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/version", nil)
	response := httptest.NewRecorder()
	handler(response, request)
}
