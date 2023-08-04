package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("[v3] Hello, Kubernetes! From host: %s\n", host))
}

func main() {
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("ok"))
	})
	http.HandleFunc("/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	})
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
