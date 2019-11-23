package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		msg, _ := ioutil.ReadAll(request.Body)
		fmt.Fprintf(writer, "Hello, %s \n", string(msg))
	})

	fmt.Println("server is starting...")

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}

}
