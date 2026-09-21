package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProdHandlerAllowsClientVersionHeader(t *testing.T) {
	const gameURL = "https://game.r117.fun"

	request := httptest.NewRequest(http.MethodOptions, "/account/register", nil)
	response := httptest.NewRecorder()
	prodHandler(http.NewServeMux(), gameURL).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected preflight status %d, got %d", http.StatusOK, response.Code)
	}
	if response.Header().Get("Access-Control-Allow-Origin") != gameURL {
		t.Fatalf("unexpected allowed origin: %q", response.Header().Get("Access-Control-Allow-Origin"))
	}
	if !strings.Contains(response.Header().Get("Access-Control-Allow-Headers"), "PKR-Client-Version") {
		t.Fatalf("PKR-Client-Version missing from allowed headers: %q", response.Header().Get("Access-Control-Allow-Headers"))
	}
}
