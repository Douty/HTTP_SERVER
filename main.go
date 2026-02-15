package main

import (
	"bufio"
	"fmt"
	"httpserver/request"
	"httpserver/response"
	"httpserver/router"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		requestReader := bufio.NewReader(conn)
		request, err := request.ParseRequest(requestReader)
		if err != nil {
			if err != io.EOF && !strings.Contains(err.Error(), "connection closed") {
				fmt.Println("Error:", err)
			}
			return
		}

		res, err := response.GenerateResponse(request)
		if err != nil {
			fmt.Print("Error has occured generating a response", err)
		}
		conn.Write(res)
	}

}

func main() {
	socket, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 80")

	if err := router.GenerateContentMap(); err != nil {
		log.Fatal("Failed to load pages:", err)
	}
	fmt.Println("All content loaded!")

	for {
		connections, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(connections)
	}
}
