package report

import (
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
