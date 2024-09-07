package main

import (
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Info</h1>")
}
