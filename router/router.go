package router

import (
	"fmt"
	"httpserver/request"
	"os"
	"path/filepath"
	"strings"
)

var pages map[string][]byte

func Router(req request.Request) ([]byte, error) {
	urlRequested := req.Route

	if strings.HasPrefix(urlRequested, "/api/") {
		handleAPI()
	}

	if content, exists := pages[urlRequested]; exists {
		return content, nil
	}

	if notFound, exists := pages["/not_found"]; exists {
		return notFound, fmt.Errorf("page not found: %s", urlRequested)
	}

	return nil, fmt.Errorf("page not found: %s", urlRequested)
}

func handleAPI() {

}
func GenerateContentMap() error {
	pages = make(map[string][]byte)
	err := filepath.WalkDir("./pages", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".html") {

			cleanPath, err := filepath.Rel("./pages", path)
			if err != nil {
				return err
			}

			if strings.Contains(cleanPath, "..") {
				return nil
			}

			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			urlPath := "/" + cleanPath
			urlPath = strings.TrimSuffix(urlPath, "/index.html")
			urlPath = strings.TrimSuffix(urlPath, ".html")

			if urlPath == "/home" {
				urlPath = "/"
			}
			pages[urlPath] = file

		}

		return nil
	})
	return err

}
