package handler

import (
	"fmt"
	"io"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("========== WEBHOOK ==========")
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL.Path)

	fmt.Println("----- Headers -----")
	for k, v := range r.Header {
		fmt.Println(k, ":", v)
	}

	fmt.Println("----- Body -----")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ReadBody Error:", err)
	} else {
		fmt.Println(string(body))
	}

	fmt.Println("========== END ==========")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
