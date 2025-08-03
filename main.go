package main

import (
	"fmt"
	"net/http"
	"os"

	"http-header-checker/checker"
	"http-header-checker/report"
)

func main() {
	// Verifica se o usuário passou a url como argumento
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <URL>")
		return
	}

	url := os.Args[1] // Pega a url

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Analisando:", url)
	results := checker.CheckSecurityHeaders(resp.Header)

	// Monta o relatório
	reportData := report.HeaderResult{
		URL:     url,
		Headers: results,
	}

	err = report.SaveResultToFile(reportData, "report.json")
	if err != nil {
		fmt.Println("Erro ao salvar relatório:", err)
		return
	}

	fmt.Println("Relatório salvo como 'report.json'")
}
