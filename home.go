package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var visited int = 0

func Home(w http.ResponseWriter, r *http.Request) {
	visited++
	var str string = "Homepage was visited " + strconv.Itoa(visited) + " times"
	fmt.Fprintf(w, str)
	fmt.Println("sencond: ", visited)
}
