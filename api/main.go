package main

import (
	"net/http"

	"github.com/PedroGuilhermeSilv/golang-vercel"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":8080", nil)
}