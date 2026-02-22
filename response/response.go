package response

import (
	"httpserver/request"
	"httpserver/router"
	"httpserver/status"

	"strconv"
	"time"
)

func GenerateResponse(req request.Request) ([]byte, error) {
	var httpResponse string
	data, err := router.Router(req)

	if err != nil {
		httpResponse += "HTTP/1.1 " + status.ToString(err.StatusCode) + " " + err.Message + "\r\n"
		httpResponse += "date: " + time.Now().Format(time.RFC1123) + "\r\n"
		httpResponse += "Server: " + "Custom HTTP server\r\n"
		httpResponse += "Content-Type: text/html\r\n"
		errorContent := "<html><body><h1>" + strconv.Itoa(int(err.StatusCode)) + " " + err.Message + "</h1></body></html>"
		httpResponse += "Content-Length: " + strconv.Itoa(len(errorContent)) + "\r\n"
		httpResponse += "\r\n"
		httpResponse += errorContent
	} else {
		httpResponse += "HTTP/1.1 " + status.ToString(status.OK) + " OK\r\n"
		httpResponse += "date: " + time.Now().Format(time.RFC1123) + "\r\n"
		httpResponse += "Server: " + "Custom HTTP server\r\n"
		httpResponse += "Content-Type: " + data.ContentType + "\r\n"
		httpResponse += "Content-Length: " + strconv.Itoa(len(data.Content)) + "\r\n"
		httpResponse += "\r\n"
		httpResponse += string(data.Content)
	}

	return []byte(httpResponse), nil
}
