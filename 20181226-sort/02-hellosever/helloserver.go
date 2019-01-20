package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request){
		//fmt.Fprintln(writer, "hello world")
		fmt.Fprintf(writer, "<p>hello %s</p>",request.FormValue("name"))
	})
	http.ListenAndServe(":8888", nil)
}
