package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"Hello from server\n")
	})

	fmt.Println("before the server")

	http.ListenAndServe(":8001",nil)

	fmt.Println("after the server")
}
