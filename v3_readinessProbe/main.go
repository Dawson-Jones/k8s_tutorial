package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v2] Hello, Kubernetes!\n")
}

func main() {
	started := time.Now()
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 15 {
			writer.WriteHeader(500)
			writer.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			writer.WriteHeader(200)
			writer.Write([]byte("ok"))
		}
	})
	http.HandleFunc("/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(500) // 有问题的版本
	})
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
