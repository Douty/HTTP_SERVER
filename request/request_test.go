package request

import (
	"io"
	"strings"
	"testing"
)

func CompareRequests(res Request, expectedRes Request, isVaild bool, t *testing.T) {
	t.Helper()

	if res.Method != expectedRes.Method {

		t.Errorf("Wrong method, was expecting %q. Got %q", expectedRes.Method, res.Method)
	}

	if res.Route != expectedRes.Route {
		t.Errorf("Wrong Route, was expecting %q. Got %q", expectedRes.Route, res.Route)
	}

	if res.Version != expectedRes.Version {

		t.Errorf("Wrong Version, was expecting %s. Got %s", expectedRes.Version, res.Version)
	}

	if len(res.Headers) != len(expectedRes.Headers) {
		t.Errorf("Header count mismatch: expected %d, got %d", len(expectedRes.Headers), len(res.Headers))
	}

	for key, value := range expectedRes.Headers {
		if res.Headers[key] != value {
			t.Errorf("Header %q: expected %q, got %q", key, value, res.Headers[key])
		}
	}

	if len(res.Query) != len(expectedRes.Query) {
		t.Errorf("Query count mismatch: expected %d, got %d", len(expectedRes.Query), len(res.Query))
	}

	for key, value := range expectedRes.Query {

		if res.Query[key] != value {
			t.Errorf("Query %q: expected %q, got %q", key, value, res.Query[key])
		}
	}

	if res.Body != expectedRes.Body {
		t.Errorf("Wrong body received, was expecting %q. Got %q", expectedRes.Body, res.Body)
	}
}
func TestParsingVaildGetRequest(t *testing.T) {
	testRequest := "GET /api/v1/users?id=123 HTTP/1.1\r\nHost: example.com\r\n User-Agent: Mozilla/5.0\r\n Accept: application/json\r\n"
	stringReader := strings.NewReader(testRequest)
	expectedRes := Request{
		Method:  "GET",
		Route:   "/api/v1/users",
		Version: "HTTP/1.1",
		Headers: map[string]string{
			"Host":       "example.com",
			"User-Agent": "Mozilla/5.0",
			"Accept":     "application/json",
		},
		Query: map[string]string{
			"id": "123",
		},
		Body: "",
	}

	res, err := ParseRequest(stringReader)
	if err != io.EOF && !strings.Contains(err.Error(), "connection closed") {
		t.Fatalf("Error: %s", err)
	}

	CompareRequests(res, expectedRes, true, t)

}
func TestParsingInVaildGetRequest(t *testing.T) {
	badInput := "GET /api/v1/users\r\nHost: example.com\r\n\r\n"

	_, err := ParseRequest(strings.NewReader(badInput))

	if err != nil {
		t.Logf("Caught expected error: %v\n", err)
	} else {
		t.Errorf("Expected an error for invalid request line, but got none")
	}

}

// API NOT IMPLEMENTED YET
// func TestParsingVaildPostRequest(t *testing.T) {

// }
