package main

import (
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

// func generateHttpGetResponse() string {
// 	htmlFile := readHTMLFile("index.html")
// 	httpResponse := "HTTP/1.1 200 OK\r\n"
// 	httpResponse += "date: " + time.Now().Format(time.RFC1123) + "\r\n"
// 	httpResponse += "Server: " + "Custom HTTP server\r\n"
// 	httpResponse += "content-type: " + "text/html\r\n"
// 	httpResponse += "content-length: " + strconv.Itoa(len(htmlFile)) + "\r\n"
// 	httpResponse += "\r\n"
// 	httpResponse += htmlFile

// 	return httpResponse
// }

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	request, err := request.ParseRequest(conn)
	if err != nil {
		if err != io.EOF && !strings.Contains(err.Error(), "connection closed") {
			fmt.Println("Error:", err)
		}
		return
	}

	res, err := response.GenerateResponse(request)
	if err != nil {
		fmt.Print("Error has occured in main", err)
	}
	conn.Write(res)

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
