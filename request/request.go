package request

import (
	"fmt"
	"httpserver/pool"
	"io"
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
	Method  Method
	Route   string
	Version string
	Headers map[string]string
	Query   map[string]string
	Body    string
}

func ParseRequest(r io.Reader) (Request, error) {
	var request Request

	request.Headers = make(map[string]string)
	request.Query = make(map[string]string)

	broswerRequest := pool.GetReader(r)
	defer pool.PutReader(broswerRequest)

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
	request.Version = parts[2]

	// collecting the route and queries (for get requests)

	fullroute := parts[1]

	routepart := strings.SplitN(fullroute, "?", 2)
	request.Route = routepart[0]

	if len(routepart) == 2 {
		if request.Method == GET {
			queries := routepart[1]
			fields := strings.Split(queries, "&")
			for _, field := range fields {
				pair := strings.SplitN(field, "=", 2)
				if len(pair) == 2 {
					key := pair[0]
					value := pair[1]
					request.Query[key] = value
				} else if len(pair) == 1 {
					request.Query[pair[0]] = ""
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
			key := strings.TrimSpace(fields[0])
			value := strings.TrimSpace(fields[1])

			request.Headers[key] = value
		}

	}

	// Collecting body information
	bodyLength := 0
	if value, present := request.Headers["Content-Length"]; present {
		trimmedValue := strings.TrimSpace(value)
		bodyLength, _ = strconv.Atoi(trimmedValue)
	}

	if bodyLength > 0 {

		bodyBuffer := pool.ReadBufferPool.Get().([]byte)
		defer pool.ReadBufferPool.Put(bodyBuffer)

		if bodyLength > len(bodyBuffer) {
			bodyBuffer = make([]byte, bodyLength)
			defer func() {}()
		}

		_, err := io.ReadFull(broswerRequest, bodyBuffer[:bodyLength])
		if err != nil {
			return request, fmt.Errorf("Failed to read request body: %w", err)
		}

		request.Body = string(bodyBuffer[:bodyLength])

	}
	return request, nil
}
