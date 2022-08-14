package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/hioan-dev/go-auth/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)

	fmt.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
