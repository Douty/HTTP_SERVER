package main

import (
	"bufio"
	"fmt"
	"httpserver/request"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func readHTMLFile(filepath string) string {
	var HTML string
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		HTML += line
	}
	return HTML
}

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
// func handleHttpPostResponse() string {
// 	httpResponse := "HTTP/1.1 201 Created\r\n"
// 	httpResponse += "date: " + time.Now().Format(time.RFC1123) + "\r\n"
// 	httpResponse += "Server: " + "Custom HTTP server\r\n"
// 	httpResponse += "content-type: " + "text/html\r\n"

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

	fmt.Println(request)

}

func main() {
	socket, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 80")
	for {
		connections, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(connections)
	}
}
