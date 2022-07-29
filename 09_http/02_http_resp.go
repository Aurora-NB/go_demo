package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", handleGo)
	var a http.HandlerFunc = handleIndex
	http.ListenAndServe(":8888", a)
}

func handleGo(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello go\n")
}
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello index\n")
}
