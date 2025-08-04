package checker

import (
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
	"Content-Security-Policy":   "Controla quais recursos podem ser carregados, protegendo contra XSS.",
	"Strict-Transport-Security": "Força uso de HTTPS, impedindo downgrade para HTTP.",
	"X-Content-Type-Options":    "Impede que o browser interprete tipos de conteúdo incorretamente.",
	"X-Frame-Options":           "Evita que a página seja carregada dentro de frames (clickjacking).",
	"Referrer-Policy":           "Controla o que é enviado no cabeçalho Referer.",
	"Permissions-Policy":        "Restringe acesso a APIs sensíveis como câmera e localização.",
	"X-XSS-Protection":          "Ativa (ou desativa) proteção contra XSS no navegador.",
}

// Função que recebe os headers da resposta e imprime os de segurança
func CheckSecurityHeaders(headers http.Header) map[string]string {
	results := make(map[string]string)

	for _, header := range securityHeaders {
		value := headers.Get(header)
		if value == "" {
			results[header] = "MISSING - " + headerDescriptions[header]
		} else {
			results[header] = value + " ✓ (" + headerDescriptions[header] + ")"
		}
	}

	return results
}
