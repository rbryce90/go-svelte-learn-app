package main

import (
	"fmt"
	"net/http"
)

const portNum = ":3001"

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	fmt.Println("Starting port on ", portNum)
	fmt.Println("visited in main second: ", visited)
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		fmt.Println("Something aint right MF")
	}
}
