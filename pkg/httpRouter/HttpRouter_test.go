package httpRouter

import (
	"net/http"
	"testing"
)

func TestRouter_RegisterSingle(t *testing.T) {
	router := New()
	if router == nil {
		t.Fatal("New() returned nil")
	}
	router.Register("GET", "/test", nil)

	if len(router.routes) != 1 {
		t.Fatalf("Expected 1 route, got %d", len(router.routes))
	}

	if router.routes[0].Method != "GET" {
		t.Fatalf("Expected method GET, got %s", router.routes[0].Method)
	}

	if router.routes[0].Path != "/test" {
		t.Fatalf("Expected path /test, got %s", router.routes[0].Path)
	}

	if router.routes[0].Handler != nil {
		t.Fatalf("Expected nil handler, got %v", router.routes[0].Handler)
	}
	req, _ := http.NewRequest("GET", "/nonexistent", nil)

	if router.Match(req) != nil {
		t.Fatalf("Expected nil, got %v", router.Match(req))
	}
}
func TestRouter_RegisterMultiple(t *testing.T) {
	router := New()
	if router == nil {
		t.Fatal("New() returned nil")
	}
	router.Register("GET", "/test", &Handler)
	router.Register("POST", "/test", nil)
	if router.routes == nil {
		t.Fatal("Expected non-nil routes")
	}
	if len(router.routes) != 2 {
		t.Fatal("Expected 2 routes")
	}

	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	if router.Match(req) == nil {
		t.Fatal("Expected route")
	}
}
