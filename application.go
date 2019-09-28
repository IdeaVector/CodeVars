package main

import (
	"github.com/CodeWars/route"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", route.Handler) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
