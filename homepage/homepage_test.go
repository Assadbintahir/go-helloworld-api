package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/home", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}

	expected := `Hello from Rahim yar khan!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v expected %v",
			rr.Body.String(), expected)
	}
}
