package checker

import (
	"fmt"
	"net/http"
)

// Lista de headers de segurança a serem checadas
var securityHeaders = []string{
	"Content-Security-Policy",
	"Strict-Transport-Security",
	"X-Content-Type-Options",
	"X-Frame-Options",
	"Referrer-Policy",
	"Permissions-Policy",
	"X-XSS-Protection",
}

var headerDescriptions = map[string]string{
	"Content-Security-Policy":   "Controla quais recursos podem ser carregados, protegendo contra XSS",
	"Strict-Transport-Security": "Força uso de HTTPS, impedindo downgrade para HTTP",
	"X-Content-Type-Options":    "Impede que o browser interprete tipos de conteúdo incorretamente",
	"X-Frame-Options":           "Evita que a página seja carregada dentro de frames (clickjacking)",
	"Referrer-Policy":           "Controla o que é enviado no cabeçalho Referer",
	"Permissions-Policy":        "Restringe acesso a APIs sensíveis como câmera e localização",
	"X-XSS-Protection":          "Ativa (ou desativa) proteção contra XSS no navegador",
}

// Resultado combinado
type Result struct {
	URL     string `json:"url"`
	Headers map[string]string
	Score   string `json:"score"`
}

// Função que recebe os headers da resposta e imprime os de segurança
func CheckURL(url string) Result {
	resp, err := http.Get(url)
	if err != nil {
		return Result{
			URL:     url,
			Headers: map[string]string{"Erro": err.Error()},
			Score:   "0%",
		}
	}
	defer resp.Body.Close()

	headers := resp.Header
	present := 0
	total := len(securityHeaders)
	resultHeaders := make(map[string]string)

	for _, header := range securityHeaders {
		value := headers.Get(header)
		if value == "" {
			resultHeaders[header] = "MISSING - " + headerDescriptions[header]
		} else {
			resultHeaders[header] = value + " ✓ (" + headerDescriptions[header] + ")."
			present++
		}
	}

	score := fmt.Sprintf("%d/%d (%.0f%%)", present, total, (float64(present)/float64(total))*100)

	return Result{
		URL:     url,
		Headers: resultHeaders,
		Score:   score,
	}
}
