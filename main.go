package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	redigo "github.com/gomodule/redigo/redis"
)

type Client struct {
	Pool *redigo.Pool
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/", pong)
	http.HandleFunc("/hello", hello)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	p := fmt.Sprintf(":%v", port)
	fmt.Println("Heruku asign Port ", p)
	http.ListenAndServe(p, nil)
}
