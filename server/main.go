package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	conn.Read(request)
	fmt.Println(string(request))
	htmlFile := readHTMLFile("index.html")
	conn.Write([]byte(htmlFile))

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
