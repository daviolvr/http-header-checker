package main

import (
	"bufio"
	"flag"
	"fmt"
	"http-header-checker/checker"
	"http-header-checker/report"
	"os"
	"strings"
)

func main() {
	// Define as flags
	urlFlag := flag.String("url", "", "URL única para verificar os headers")
	fileFlag := flag.String("file", "", "Arquivo com lista de URLs")
	outputFlag := flag.String("output", "relatorio.json", "Nome do arquivo de saída")

	flag.Parse()

	// Verificação (precisa passar pelo menos uma das flags)
	if *urlFlag == "" && *fileFlag == "" {
		fmt.Println("Você precisa passar -url ou -file.")
		flag.Usage()
		return
	}

	var results []checker.Result

	// Se for uma URL única
	if *urlFlag != "" {
		result := processURL(*urlFlag)
		results = append(results, result)
	}

	// Se for um arquivo com várias URLs
	if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Println("Erro ao abrir o arquivo:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			url := strings.TrimSpace(scanner.Text())
			if url == "" {
				continue
			}
			result := processURL(url)
			results = append(results, result)
		}
	}

	// Salva o relatório em JSON
	err := report.SaveMultipleResultsToFile(results, *outputFlag)
	if err != nil {
		fmt.Println("Erro ao salvar relatório:", err)
		return
	}

	fmt.Println("Relatório salvo em:", *outputFlag)
}

func processURL(url string) checker.Result {
	return checker.CheckURL(url)
}
