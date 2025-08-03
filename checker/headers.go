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

// Função que recebe os headers da resposta e imprime os de segurança
func CheckSecurityHeaders(headers http.Header) map[string]string {
	results := make(map[string]string)

	for _, header := range securityHeaders {
		value := headers.Get(header)
		if value == "" {
			results[header] = "MISSING"
		} else {
			results[header] = value
		}
	}

	return results
}
