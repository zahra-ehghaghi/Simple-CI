package main

import (
	"fmt"
	"net/http"
	"time"
)

const timeout = 3 * time.Second

func main() {
	http.HandleFunc("/hello", hello)
	server := &http.Server{
		Addr:              ":8090",
		ReadHeaderTimeout: timeout,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from zahra\n")

}
