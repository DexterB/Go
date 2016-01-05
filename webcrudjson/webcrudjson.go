package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ID int `json:"int"`
}


func userRouter (resp http.ResponseWriter, req *http.Request) {

	ourUser := User{};
	ourUser.Name = "Dexter Paul Bradshaw"
	ourUser.Email = "bradshaw.dexter@gmail.com"
	ourUser.ID = 100

	output,_ := json.Marshal(&ourUser)
	fmt.Fprintln(resp, string(output))
}

func main() {

	fmt.Println("Starting JSON Server")
	http.HandleFunc("/user", userRouter)
	http.ListenAndServe(":8090", nil)
}
