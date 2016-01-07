package main

import (
	"encoding/xml"
	"net/http"
	"fmt"
)

type User struct {
	Name string `xml:"name"`
	Email string `xml:"email"`
	ID int `xml:"int"`
}


func userRouter (resp http.ResponseWriter, req *http.Request) {

	ourUser := User{};
	ourUser.Name = "Dexter Paul Bradshaw"
	ourUser.Email = "bradshaw.dexter@gmail.com"
	ourUser.ID = 100

	output,_ := xml.Marshal(&ourUser)
	fmt.Fprintln(resp, string(output))
}

func main() {

	fmt.Println("Starting JSON Server")
	http.HandleFunc("/user", userRouter)
	http.ListenAndServe(":8090", nil)
}
