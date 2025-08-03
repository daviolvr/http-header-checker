package main

import (
	"fmt"
	"net/http"

	"http-header-checker/checker"
)

func main() {
	url := "http://127.0.0.1:8000/api/v1/auth/me/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Analisando:", url)
	checker.CheckSecurityHeaders(resp.Header)
}
