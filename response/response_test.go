package response

import (
	"fmt"
	"httpserver/request"
	"httpserver/router"
	"strconv"
	"strings"
	"testing"
)

func parseHTTPResponse(response []byte) (statusLine string, headers map[string]string, body string, err error) {
	responseStr := string(response)

	parts := strings.Split(responseStr, "\r\n\r\n")
	if len(parts) != 2 {
		return "", nil, "", fmt.Errorf("invalid HTTP response format")
	}

	headerSection := parts[0]
	body = parts[1]

	lines := strings.Split(headerSection, "\r\n")
	if len(lines) < 1 {
		return "", nil, "", fmt.Errorf("no status line found")
	}

	statusLine = lines[0]
	headers = make(map[string]string)

	for _, line := range lines[1:] {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			headers[parts[0]] = parts[1]
		}
	}

	return statusLine, headers, body, nil
}

func TestVaildResponse(t *testing.T) {
	err := router.GenerateContentMapFromPath("../pages")
	if err != nil {
		t.Fatalf("Failed to load pages: %v", err)
	}

	req := request.Request{Route: "/"}
	response, err := GenerateResponse(req)

	if err != nil {
		t.Fatalf("GenerateResponse returned error: %v", err)
	}

	statusLine, headers, body, err := parseHTTPResponse(response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if statusLine != "HTTP/1.1 200 OK" {
		t.Errorf("Expected status 'HTTP/1.1 200 OK', got '%s'", statusLine)
	}

	if headers["Content-Type"] != "text/html" {
		t.Errorf("Expected Content-Type 'text/html', got '%s'", headers["Content-Type"])
	}

	expectedLength := strconv.Itoa(len(body))
	if headers["Content-Length"] != expectedLength {
		t.Errorf("Content-Length mismatch: header says %s, body is %s bytes",
			headers["Content-Length"], expectedLength)
	}

}

func TestInvaildResponse(t *testing.T) {
	err := router.GenerateContentMapFromPath("../pages")
	if err != nil {
		t.Fatalf("Failed to load pages: %v", err)
	}

	req := request.Request{Route: "/invaild"}
	response, err := GenerateResponse(req)

	if err != nil {
		t.Fatalf("GenerateResponse returned error: %v", err)
	}
	statusLine, headers, body, err := parseHTTPResponse(response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if statusLine != "HTTP/1.1 404 Not Found" {
		t.Errorf("Expected status 'HTTP/1.1 404 Not Found', got '%s'", statusLine)
	}

	if headers["Content-Type"] != "text/html" {
		t.Errorf("Expected Content-Type 'text/html', got '%s'", headers["Content-Type"])
	}

	expectedLength := strconv.Itoa(len(body))
	if headers["Content-Length"] != expectedLength {
		t.Errorf("Content-Length mismatch: header says %s, body is %s bytes",
			headers["Content-Length"], expectedLength)
	}
}
