package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListHandler(t *testing.T) {
	url := "http://localhost:8080/list"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error("Error creating request")
	}

	rec := httptest.NewRecorder()

	ListHandler(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Fatalf("Expected status %d, got %d", 201, rec.Result().StatusCode)
	}
}
