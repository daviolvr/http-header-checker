package report

import (
	"encoding/csv"
	"encoding/json"
	"http-header-checker/checker"
	"os"
)

// Salva o resultado como JSON em um arquivo
func SaveMultipleResultsToFile(results []checker.Result, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") // deixa o JSON legível
	return encoder.Encode(results)
}

// Salva um relatório CSV com as colunas: URL, Score
func SaveCSV(results []checker.Result, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Cabeçalho
	err = writer.Write([]string{"URL", "Score"})
	if err != nil {
		return err
	}

	// Cada linha do resultado
	for _, r := range results {
		err := writer.Write([]string{r.URL, r.Score})
		if err != nil {
			return err
		}
	}

	return nil
}
