package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	Message string `json:"message"`
}

func requestHandler (w http.ResponseWriter, r *http.Request) {
	// message := API{"Hello, world!"}

	response := &API{Message : "Hello, world!"}

	output, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Something went wrong.")
	}

	fmt.Fprint(w, string(output))
}

func main() {
	http.HandleFunc("/api", requestHandler)
	http.ListenAndServe(":9090", nil)
}
