package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/son3941/dasei-plugin/handler"
)

func main() {

	// トップページ確認用
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dasei-plugin is running")
	})

	// mixi2 Webhook
	http.HandleFunc("/events", handler.Webhook)

	// RenderのPORT取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server started :" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
