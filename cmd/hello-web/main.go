package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

// this means all the available IP addresses on port 8080
var serverAddress = ":8080"

func main() {
	// define our handlers
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/user/{user}/", logs(helloPath))

	// stat the webserver
	slog.Info("server running", "address", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		slog.Error("ListenAndServe: ", "error", err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func helloPath(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	if user == "coffee" {
		// unfortunately we are permanently a teapot, not coffee for anybody
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(w, "I'm a teapot!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", user)
}

// We used http.HandleFunc to register a handler function to an http route of
// out server. That function signature is:
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)).
// we can have a taste of how we can create some middleware by having a function
// that takes as parameter and returns a func(ResponseWriter, *Request).
// For example here is a middleware that logs the requests.
func logs(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.URL)
		next(w, r)
	}
}
