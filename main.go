package main

import (
	"fmt"
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
		fmt.Println("hit")
		_, _ = writer.Write([]byte("Hello World! " + port + " " + path))
	})
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
	fmt.Println("Done ...")
}
