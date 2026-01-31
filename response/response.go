package response

import (
	"httpserver/request"
	"httpserver/router"
	"httpserver/status"
	"strconv"
)

func GenerateResponse(req request.Request) ([]byte, error) {
	var httpResponse string
	data, err := router.Router(req)

	if err != nil {
		httpResponse += "HTTP/1.1 " + status.ToString(err.StatusCode) + " " + err.Message + "\r\n"
	} else {
		httpResponse += "HTTP/1.1 " + status.ToString(status.OK) + " OK\r\n"
	}

	httpResponse += "Content-Type: text/html\r\n"
	httpResponse += "Content-Length: " + strconv.Itoa(len(data)) + "\r\n"
	httpResponse += "\r\n"
	httpResponse += string(data)

	return []byte(httpResponse), nil
}
