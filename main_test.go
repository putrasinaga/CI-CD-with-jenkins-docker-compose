package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHelloServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloServer)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, World putra!!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestEndointsaya(t *testing.T) {
	req, err := http.NewRequest("GET", "/saya", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Endointsaya)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "ini adalah endpoint saya!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestMain(m *testing.M) {
	// start the HTTP server
	go main()

	// wait for the server to start
	resp, err := http.Get("http://localhost:3000/")
	if err != nil || resp.StatusCode != http.StatusOK {
		panic("Could not start server")
	}

	// run the tests
	exitCode := m.Run()

	// stop the server
	resp.Body.Close()

	// exit with the same code as the tests
	os.Exit(exitCode)
}
