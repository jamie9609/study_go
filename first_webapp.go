package main

import (
	"fmt"
	"net/http"
)

func handler(write http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer, "hello world, %s!", request.URL.Path[1:])
}

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}