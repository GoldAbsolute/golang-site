package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, "Heh")
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
