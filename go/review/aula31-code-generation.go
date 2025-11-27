package main

import "fmt"

/*
 * ============================================================================
 * AULA 31: CODE GENERATION
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - go generate
 * - Build tags
 * - Conditional compilation
 */

//go:generate echo "Gerando código..."

func main() {
	fmt.Println("=== AULA 31: CODE GENERATION ===\n")
	fmt.Println("go generate: Executa comandos de geração")
	fmt.Println("Build tags: // +build ou //go:build")
	fmt.Println("Conditional compilation: Compilar código condicionalmente")
}



