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

// Função que recebe os headers da resposta e imprime os de segurança
func CheckSecurityHeaders(headers http.Header) {
	fmt.Println("Verificando headers de segurança:")

	for _, header := range securityHeaders {
		value := headers.Get(header)
		if value == "" {
			fmt.Printf("[!] Faltando: %s\n", header)
		} else {
			fmt.Printf("[OK] %s: %s\n", header, value)
		}
	}
}
