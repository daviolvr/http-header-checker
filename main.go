package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://example.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Headers da resposta:")
	for key, value := range resp.Header {
		fmt.Printf("%s: %s\n", key, value)
	}
}
