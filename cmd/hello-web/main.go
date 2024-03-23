package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", helloServer)
	http.HandleFunc("GET /user/{user}/", helloPath)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("ListenAndServe: ", "MSG", err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
}

func helloPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.PathValue("user"))
}
