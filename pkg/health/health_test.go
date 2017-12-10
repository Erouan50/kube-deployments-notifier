package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckReply)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("healthCheckReply handler didn't return an HTTP 200 status code")
	}

	if rr.Body.String() != "ok\n" {
		t.Errorf("healthCheckReply didn't return 'ok\n'")
	}
}
