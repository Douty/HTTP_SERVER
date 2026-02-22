package main

import (
	"errors"
	"fmt"
	"httpserver/pool"
	"httpserver/request"
	"httpserver/response"
	"httpserver/router"
	"io"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	requestReader := pool.GetReader(conn)

	defer pool.PutReader(requestReader)

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		request, err := request.ParseRequest(requestReader)
		if err != nil {
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				return
			}

			if err != io.EOF {
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
