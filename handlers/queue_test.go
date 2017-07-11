package handlers

import (
	"net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
)

func TestQueueHandler(t *testing.T) {

	url := "http://localhost:8080/ydl?who=test&v=http://unittests.com"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error("Error creating request")
	}

	rec := httptest.NewRecorder()

	QueueHandler(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Fatalf("Expected status %d, got %d", 201, rec.Result().StatusCode)
	}

}
