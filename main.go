package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	//dbURL := os.Getenv("DB_URL")
	dbPassword := os.Getenv("DB_PASSWORD")
	io.WriteString(w, fmt.Sprintf("[v3] Hello, Kubernetes! From host: %s, "+
		//"Get Database connect URL: %s\n",
		"Get Database connect Password: %s\n",
		host,
		//dbURL
		dbPassword,
	))
}

func main() {
	//started := time.Now()
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		//duration := time.Since(started)
		//if duration.Seconds() > 15 {
		//	writer.WriteHeader(500)
		//	writer.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		//} else {
		writer.WriteHeader(200)
		writer.Write([]byte("ok"))
		//}
	})
	http.HandleFunc("/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200) // 无问题的版本
		//writer.WriteHeader(500) // 有问题的版本
	})
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
