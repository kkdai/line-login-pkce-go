package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogoutWithoutToken(t *testing.T) {
	// Reset the current access token
	currentAccessToken = ""

	// Create a request to the logout endpoint
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(logout)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check that we get a redirect (303 See Other)
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check that we redirect to the home page
	location := rr.Header().Get("Location")
	if location != "/" {
		t.Errorf("handler returned wrong redirect location: got %v want /", location)
	}
}

func TestLogoutClearsToken(t *testing.T) {
	// Set a fake access token
	currentAccessToken = "fake_token_for_testing"

	// Create a request to the logout endpoint
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(logout)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check that the token was cleared
	if currentAccessToken != "" {
		t.Errorf("logout did not clear the access token: got %v want empty string", currentAccessToken)
	}

	// Check that we get a redirect
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}

func TestBrowseHandler(t *testing.T) {
	// Create a request to the browse endpoint
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(browse)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check that we get a 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
