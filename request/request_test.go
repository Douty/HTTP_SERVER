package request

import (
	"io"
	"strings"
	"testing"
)

func CompareRequests(res Request, vaildAnswer Request, t *testing.T) {
	t.Helper()

	if res.Method != vaildAnswer.Method {

		t.Errorf("Wrong method, was expecting %q. Got %q", vaildAnswer.Method, res.Method)
	}

	if res.Route != vaildAnswer.Route {
		t.Errorf("Wrong Route, was expecting %q. Got %q", vaildAnswer.Route, res.Route)
	}

	if res.Version != vaildAnswer.Version {

		t.Errorf("Wrong Version, was expecting %s. Got %s", vaildAnswer.Version, res.Version)
	}

	if len(res.Headers) != len(vaildAnswer.Headers) {
		t.Errorf("Header count mismatch: expected %d, got %d", len(vaildAnswer.Headers), len(res.Headers))
	}

	for key, value := range vaildAnswer.Headers {
		if res.Headers[key] != value {
			t.Errorf("Header %q: expected %q, got %q", key, value, res.Headers[key])
		}
	}

	if len(res.Query) != len(vaildAnswer.Query) {
		t.Errorf("Query count mismatch: expected %q, got %q", len(vaildAnswer.Query), len(res.Query))
	}

	for key, value := range vaildAnswer.Query {

		if res.Query[key] != value {
			t.Errorf("Query %q: expected %q, got %q", key, value, res.Query[key])
		}
	}

	if res.Body != vaildAnswer.Body {
		t.Errorf("Wrong body received, was expecting %q. Got %q", vaildAnswer.Body, res.Body)
	}
}
func TestParsingGetRequest(t *testing.T) {
	testRequest := "GET /api/v1/users?id=123 HTTP/1.1\r\nHost: example.com\r\n User-Agent: Mozilla/5.0\r\n Accept: application/json\r\n"
	stringReader := strings.NewReader(testRequest)
	vaildAnswer := Request{
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

	CompareRequests(res, vaildAnswer, t)

}
