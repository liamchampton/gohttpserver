package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	t.Run("returns message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Home(response, request)

		got := response.Body.String()
		want := "This is a test for liamstoolchain"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
