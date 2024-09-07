package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const portNum = ":3001"

var visited int = 0

func Home(w http.ResponseWriter, r *http.Request) {
	visited++
	var str string = "Visited " + strconv.Itoa(visited) + " times"
	fmt.Fprintf(w, str)
	fmt.Println("sencond: ", visited)
}
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Info</h1>")
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/visited", Home)
	http.HandleFunc("/info", Info)

	fmt.Println("Starting port on ", portNum)
	fmt.Println("visited in main second: ", visited)
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		fmt.Println("Something aint right MF")
	}
}
