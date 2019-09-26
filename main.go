package main

import (
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	path := "/"
	if len(os.Args) >= 2 {
		port = ":" + os.Args[1]
	}
	if len(os.Args) >= 3 {
		path = os.Args[2]
	}

	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})
	_ = http.ListenAndServe(port, nil)
}
