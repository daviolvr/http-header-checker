package report

import (
	"encoding/json"
	"os"
)

// Struct pra salvar o resultado da verificação
type HeaderResult struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

// Salva o resultado como JSON em um arquivo
func SaveMultipleResultsToFile(results []HeaderResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") // deixa o JSON legível
	return encoder.Encode(results)
}
