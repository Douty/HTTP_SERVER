package handler

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
	CONNECT Method = "CONNECT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	TRACE   Method = "TRACE"
)

type Request struct {
	Method   Method
	route    string
	version  string
	headers  map[string]string
	query    map[string]string
	Body     string
	formData map[string]string
}

func ParseRequest(conn net.Conn) (Request, error) {
	var request Request

	request.headers = make(map[string]string)
	request.query = make(map[string]string)

	broswerRequest := bufio.NewReader(conn)

	//collects the the method, route and version sent
	requestLine, err := broswerRequest.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return request, fmt.Errorf("connection closed before request line: %w", err)
		}
		return request, fmt.Errorf("Failed to properly read request line: %w", err)
	}
	requestLine = strings.TrimSpace(requestLine)

	if requestLine == "" {
		return request, fmt.Errorf("Empty request line received")
	}

	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return request, fmt.Errorf("Invalid request recieved")
	}
	request.Method = Method(parts[0])
	request.version = parts[2]

	// collecting the route and queries (for get requests)

	fullroute := parts[1]

	routepart := strings.SplitN(fullroute, "?", 2)
	request.route = routepart[0]

	if len(routepart) == 2 {
		if request.Method == GET {
			queries := routepart[1]
			fields := strings.Split(queries, "&")
			for _, field := range fields {
				pair := strings.SplitN(field, "=", 2)
				if len(pair) == 2 {
					key := pair[0]
					value := pair[1]
					request.query[key] = value
				} else if len(pair) == 1 {
					request.query[pair[0]] = ""
				}

			}

		}
	}

	//Collecting Header information
	for {
		headerLine, err := broswerRequest.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return request, fmt.Errorf("connection closed before request line: %w", err)
			}
			return request, fmt.Errorf("Failed to properly read Headers %w", err)
		}
		headerLine = strings.TrimSpace(headerLine)
		if headerLine == "" {
			break
		}

		fields := strings.SplitN(headerLine, ":", 2)
		if len(fields) == 2 {
			key := fields[0]
			value := fields[1]

			request.headers[key] = value
		}

	}

	// Collecting body information
	bodyLength := 0
	if value, present := request.headers["Content-Length"]; present {
		trimmedValue := strings.TrimSpace(value)
		bodyLength, _ = strconv.Atoi(trimmedValue)
	}

	if bodyLength > 0 {
		bodyBytes := make([]byte, bodyLength)
		_, err := io.ReadFull(broswerRequest, bodyBytes)
		if err != nil {
			return request, fmt.Errorf("Failed to read request body: %w", err)
		}

		request.Body = string(bodyBytes)

	}
	return request, nil
}
