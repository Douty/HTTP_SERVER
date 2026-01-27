package router

import (
	"httpserver/request"
	"testing"
)

func TestRouter(t *testing.T) {

	err := GenerateContentMapFromPath("../pages")
	if err != nil {
		t.Fatalf("Failed to load pages: %v", err)
	}

	t.Logf("Pages map length, %d Available routes:", len(pages))
	for route := range pages {
		t.Logf(" - %s", route)
	}

	if len(pages) == 0 {
		t.Fatal("No pages were loaded")
	}

	for route := range pages {
		req := request.Request{Route: route}

		page, err := Router(req)
		if err != nil {
			t.Fatalf("Expected no error, error message: %v", err)
		}

		if page == nil {
			t.Fatal("Expected page content, got nil")
		}
		content := string(page)
		if len(content) == 0 {
			t.Errorf("Page: %s content is 0", route)
		}
	}
}
