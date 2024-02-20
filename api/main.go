package main

import (
	"net/http"

	"github.com/pedro/dev/goDeployVercel/api/handler"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":8080", nil)