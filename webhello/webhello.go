package main

import (
	"encoding/json"
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

type API struct {
	Message string `json:"message"`
}

func requestHandler (w http.ResponseWriter, r *http.Request) {

	urlParams := r.URL.Query()
	name := urlParams.Get(":name")
	helloMessage := "Hello, " + name + "!"

	response := &API{helloMessage}


	output, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Something went wrong.")
	}

	fmt.Fprint(w, string(output))
}

func main() {

	mux := routes.New()
	mux.Get("/api/:name", requestHandler)
	http.Handle("/", mux)
	http.ListenAndServe(":9091", nil)

}
