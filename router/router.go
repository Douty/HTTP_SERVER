package router

import (
	"fmt"
	"httpserver/request"
	"httpserver/status"
	"os"
	"path/filepath"
	"strings"
)

type HTTPError struct {
	StatusCode status.Status
	Message    string
}
type Asset struct {
	Content     []byte
	ContentType string
}
type Context struct {
	Method request.Method
	Route  string
	Params map[string]string
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("%s, Error code: %d", err.Message, err.StatusCode)
}

var pages map[string]Asset
var apiRoute = map[string]func(Context) (Asset, *HTTPError){
	"/api/getusers": APIGetAllUsers,
}

func Router(req request.Request) (Asset, *HTTPError) {
	urlRequested := req.Route

	if req.Method != request.GET && !strings.HasPrefix(req.Route, "/api") {
		return Asset{}, &HTTPError{Message: "Method Not Allowed", StatusCode: status.NOT_ALLOWED}
	}

	if strings.HasPrefix(urlRequested, "/api/") {
		ctx := Context{Method: req.Method, Route: urlRequested, Params: req.Query}

		if fn, exists := apiRoute[urlRequested]; exists {

			return fn(ctx)
		}

		return Asset{}, &HTTPError{Message: "API Route Not Found", StatusCode: status.NOT_FOUND}
	}

	if content, exists := pages[urlRequested]; exists {
		return content, nil
	}

	if notFound, exists := pages["/not_found"]; exists {
		return notFound, &HTTPError{Message: "Not Found", StatusCode: status.NOT_FOUND}
	}

	return Asset{}, &HTTPError{Message: "Internal server Error", StatusCode: status.INTERNAL_SERVER_ERROR}
}

func handleAPI(route string) {

}

var pagesDir = "./pages"

// On startup, generate a hashmap with the
func GenerateContentMap() error {

	return GenerateContentMapFromPath(pagesDir)
}

func GenerateContentMapFromPath(folderpath string) error {
	pages = make(map[string]Asset)

	err := filepath.WalkDir(folderpath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && (strings.HasSuffix(d.Name(), ".html") ||
			strings.HasSuffix(d.Name(), ".css") ||
			strings.HasSuffix(d.Name(), ".js")) {

			cleanPath, err := filepath.Rel(folderpath, path)
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

			fileData := Asset{
				Content: file,
			}

			urlPath := "/" + filepath.ToSlash(cleanPath)

			if strings.HasSuffix(d.Name(), ".html") {
				urlPath = strings.TrimSuffix(urlPath, ".html")
				fileData.ContentType = "text/html"

				if urlPath == "/home" {
					urlPath = "/"
				}
			} else if strings.HasSuffix(d.Name(), ".css") {
				fileData.ContentType = "text/css"
			} else if strings.HasSuffix(d.Name(), ".js") {
				fileData.ContentType = "text/javascript"
			}

			pages[urlPath] = fileData
		}

		return nil
	})

	return err
}
