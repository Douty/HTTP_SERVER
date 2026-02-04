package router

import (
	"encoding/json"
	"httpserver/request"
	"httpserver/status"
	"slices"
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

func TestAPIGetAllUsers(t *testing.T) {
	t.Run("Returns success on GET request", func(t *testing.T) {
		ctx := Context{
			Method: request.GET,
		}

		data, httpErr := APIGetAllUsers(ctx)
		if httpErr != nil {
			t.Fatalf("Expected no error, got %v", httpErr.Message)
		}

		var got []string
		err := json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		expected := []string{"Alice", "Bob", "Charlie"}

		if !slices.Equal(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("Returns 405 on POST request", func(t *testing.T) {

		ctx := Context{
			Method: request.POST,
		}

		_, httpErr := APIGetAllUsers(ctx)

		if httpErr == nil {
			t.Fatal("Expected an error for POST method, but got nil")
		}

		if httpErr.StatusCode != status.NOT_ALLOWED {
			t.Errorf("Expected status 405, got %d", httpErr.StatusCode)
		}
	})
}
