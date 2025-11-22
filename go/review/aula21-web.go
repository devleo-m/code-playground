package main

import (
	"fmt"
	"net/http"
)

/*
 * ============================================================================
 * AULA 21: WEB DEVELOPMENT
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - net/http server
 * - Handlers
 * - Routing
 */

func main() {
	fmt.Println("=== AULA 21: WEB DEVELOPMENT ===\n")
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, Mundo!")
	})
	
	fmt.Println("Servidor rodando em http://localhost:8080")
	fmt.Println("Execute: go run aula21-web.go")
	// http.ListenAndServe(":8080", nil)
}
