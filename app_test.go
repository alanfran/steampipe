package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestQueryRoute tests a successful query.
func TestQueryRoute(t *testing.T) {
	app := newApp(time.Second * 10)

	req, _ := http.NewRequest("GET", "/query/zs.nekonet.xyz:27015", nil)
	resp := httptest.NewRecorder()

	app.engine.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != 200 {
		t.Error("Query failed.")
	}
}

// TestQueryNoPort tests that a query without a port will fail.
func TestQueryNoPort(t *testing.T) {
	app := newApp(time.Second * 10)

	req, _ := http.NewRequest("GET", "/query/zs.nekonet.xyz", nil)
	resp := httptest.NewRecorder()

	app.engine.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != 200 {
		t.Error("Query failed to add port to request.")
	}
}

// TestQueryWrongPort tests that a query with the wrong port will fail.
func TestQueryWrongPort(t *testing.T) {
	app := newApp(time.Second * 10)

	req, _ := http.NewRequest("GET", "/query/zs.nekonet.xyz:27016", nil)
	resp := httptest.NewRecorder()

	app.engine.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != 500 {
		t.Error("Query did not fail when it should have.")
	}
}

func TestIndex(t *testing.T) {
	app := newApp(time.Second * 10)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	app.engine.ServeHTTP(resp, req)

	result := resp.Result()
	if result.StatusCode != 200 {
		t.Error("Index not found.")
	}
}
