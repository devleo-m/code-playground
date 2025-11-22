package main

import (
	"fmt"
	"html"
)

/*
 * ============================================================================
 * AULA 30: SECURITY
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Input validation
 * - SQL injection prevention
 * - XSS prevention
 * - Secrets management
 * - crypto package
 */

func main() {
	fmt.Println("=== AULA 30: SECURITY ===\n")
	
	// Exemplo: Prevenir XSS
	userInput := "<script>alert('xss')</script>"
	safe := html.EscapeString(userInput)
	fmt.Printf("Input: %s\n", userInput)
	fmt.Printf("Safe: %s\n", safe)
	
	fmt.Println("\nBoas práticas:")
	fmt.Println("- Sempre valide input")
	fmt.Println("- Use prepared statements para SQL")
	fmt.Println("- Escape output HTML")
	fmt.Println("- Nunca commite secrets")
}
