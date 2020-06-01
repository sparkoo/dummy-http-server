package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

type Profile struct {
	Name string
	Cookies []*http.Cookie
}

func main() {
	port := ":8080"
	serverPath := "/"
	if len(os.Args) >= 2 {
		port = ":" + os.Args[1]
	}
	if len(os.Args) >= 3 {
		serverPath = os.Args[2]
	}
	name := os.Getenv("WHOAMI_NAME")

	http.HandleFunc(serverPath, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hit")
		_, _ = writer.Write([]byte("Hello World! " + port + " " + serverPath + " [" + name + "]"))
	})
	http.HandleFunc("/cookie", func(writer http.ResponseWriter, r *http.Request) {

		profile := Profile{Name: name, Cookies: r.Cookies()}
		fmt.Printf("%+v\n", r.Cookies())

		fp := path.Join("templates", "cookie.html")
		http.SetCookie(writer, &http.Cookie{Name: "test-with-path", Value: "server " + name, Path: r.URL.Path})
		http.SetCookie(writer, &http.Cookie{Name: "test-without-path", Value: "server " + name})
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(writer, profile); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening ...")
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
	fmt.Println("Done ...")
}
