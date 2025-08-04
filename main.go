package main

import (
	"bufio"
	"fmt"
	"http-header-checker/checker"
	"http-header-checker/report"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Verifica se o usuário passou a url como argumento
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <URL ou arquivo .txt>")
		return
	}

	arg := os.Args[1] // Pega a url

	if strings.HasSuffix(arg, ".txt") {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Println("Erro ao abrir arquivo:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() {
			url := scanner.Text()
			if url == "" {
				continue
			}
			count++
			processURL(url, fmt.Sprintf("report_%d.json", count))
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Erro ao ler arquivo:", err)
		}

		return
	}

	processURL(arg, "report.json")
}

func processURL(url string, output string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Analisando:", url)
	results := checker.CheckSecurityHeaders(resp.Header)

	reportData := report.HeaderResult{
		URL:     url,
		Headers: results,
	}

	err = report.SaveResultToFile(reportData, output)
	if err != nil {
		fmt.Println("Erro ao salvar relatório:", err)
	} else {
		fmt.Printf("Relatório salvo: %s\n", output)
	}
}
