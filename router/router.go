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
	//if browser requests /home
	if urlRequested == "/home" || urlRequested == "/home/index" {
		urlRequested = "/"
	}

	if content, exists := pages[urlRequested]; exists {
		fmt.Printf("DEBUG: Found route '%s'\n", urlRequested)
		return content, nil
	}

	if strings.HasSuffix(urlRequested, "/") && urlRequested != "/" {
		trimmed := strings.TrimSuffix(urlRequested, "/")
		if content, exists := pages[trimmed]; exists {
			return content, nil
		}
	}

	if !strings.HasSuffix(urlRequested, "/") {
		withSlash := urlRequested + "/"
		if content, exists := pages[withSlash]; exists {
			return content, nil
		}
	}

	if notFound, exists := pages["/not_found/404"]; exists {
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

		if d.IsDir() {
			return nil
		}
		fileExt := filepath.Ext(d.Name())

		if fileExt != ".html" && fileExt != ".css" && fileExt != ".js" {
			return nil
		}

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

		fileData := Asset{Content: file}

		urlPath := "/" + filepath.ToSlash(cleanPath)

		switch fileExt {
		case ".html":
			urlPath = strings.TrimSuffix(urlPath, ".html")
			fileData.ContentType = "text/html"

			fmt.Printf("DEBUG: Full urlPath %s\n", urlPath)
			if filepath.Base(urlPath) == "index" {
				dirPath := filepath.ToSlash(filepath.Dir(urlPath))
				switch dirPath {
				case "/home", "/", ".":
					pages["/"] = fileData
				default:
					pages[dirPath] = fileData
				}
			} else {
				pages[urlPath] = fileData
			}
		case ".css":
			fileData.ContentType = "text/css"
			pages[urlPath] = fileData
		case ".js":
			fileData.ContentType = "text/javascript"
			pages[urlPath] = fileData
		}

		fmt.Printf("DEBUG: Registered route: '%s' (ContentType: %s)\n", urlPath, fileData.ContentType)

		return nil
	})

	return err
}
