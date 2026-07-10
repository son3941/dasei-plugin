package handler

import (
	"fmt"
	"io"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("========== WEBHOOK ==========")
	fmt.Println("Method:", r.Method)

	body, _ := io.ReadAll(r.Body)

	fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
