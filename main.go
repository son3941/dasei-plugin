package main

import (
	"fmt"
	"net/http"

	"github.com/son3941/dasei-plugin/handler"
)

func main() {
	handler.Hello()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dasei-plugin is running")
	})

	fmt.Println("Server started :8080")
	http.ListenAndServe(":8080", nil)
}
