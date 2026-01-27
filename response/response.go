package response

import (
	"bufio"
	"httpserver/request"
	"httpserver/status"
	"net"
	"os"
	"strconv"
)

func readHTMLFile(filepath string) (string, error) {
	var HTML string
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		HTML += line
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return HTML, nil
}

func GenerateResponse(conn net.Conn, req request.Request) ([]byte, error) {
	httpRoute := req.Route
	var response string

	switch httpRoute {
	case "/":
		if req.Method != request.GET {
			bodyMessage := "Method Not Allowed"
			response += "HTTP/1.1 " + strconv.Itoa(int(status.NOT_ALLOWED)) + " Method Not Allowed\r\n"
			response += "Allow: " + string(request.GET) + "\r\n"
			response += "Content-Length: " + strconv.Itoa(len(bodyMessage)) + "\r\n"
			response += "\r\n"
			response += bodyMessage

			return []byte(response), nil
		}

	default:

	}

	return []byte("LOOL"), nil
}
