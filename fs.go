package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", http.StripPrefix("", fs))

	err := http.ListenAndServe(":5790", nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Starting server on port :5790")
	}
}
