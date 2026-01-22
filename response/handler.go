package handler

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Request struct {
	method  string
	route   string
	version string
	headers map[string]string
	query   map[string]string
}

func ParseRequest(conn net.Conn) Request {
	var request Request
	broswerRequest := bufio.NewReader(conn)
	for i := 0; i < broswerRequest.Size(); i++ {

		data, err := broswerRequest.ReadString('\n')
		data = strings.TrimSpace("\r\n")
		if err != nil {
			log.Fatal(err)
		}
		line := strings.Split(data, " ")

		if i == 0 {
			request.method = line[0]
			request.route = line[1]
			request.version = line[2]
		} else {
			request.headers[line[0]] = line[1]
		}

	}
	return request
}
