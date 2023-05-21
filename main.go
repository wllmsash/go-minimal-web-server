package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

const port uint16 = 8080
const webServerFilesPath string = "/var/local/www"

func main() {
	// Attribution: https://stackoverflow.com/a/57281420
	http.Handle("/", http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		var requestPath string
		if request.URL.Path == "/" {
			requestPath = "/index"
		} else {
			requestPath = request.URL.Path
		}

		var relativeFilePath string
		if cleanRequestPath := filepath.Clean(requestPath); filepath.Ext(cleanRequestPath) == "" {
			relativeFilePath = fmt.Sprintf("%s.html", cleanRequestPath)
		} else {
			relativeFilePath = cleanRequestPath
		}

		http.ServeFile(responseWriter, request, filepath.Join(webServerFilesPath, relativeFilePath))
	}))

	log.Printf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), nil))
}
