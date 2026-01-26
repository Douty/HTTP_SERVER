package response

import (
	"bufio"
	"httpserver/request"
	"log"
	"net"
	"os"
)

func readHTMLFile(filepath string) (syntax,error){
	var HTML string
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil{
		return "", err
	}
	for scanner.Scan() {
		line := scanner.Text()
		HTML += line
	}
	return HTML
}

func GenerateResponse(conn net.Conn, request request.Request) error {
	http_route := request.Route

	switch http_route {
	case "/":
		pagedata := readHTMLFile("./pages/index.html")
		conn.Write([[]byte(pagedata)], error) 
		
	}

}
