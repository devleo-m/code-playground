//go:build !windows
// +build !windows

package main

import "fmt"

/*
 * ============================================================================
 * AULA 38: BUILD CONSTRAINTS
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - // +build tags
 * - //go:build (Go 1.17+)
 * - Conditional compilation
 */

func main() {
	fmt.Println("=== AULA 38: BUILD CONSTRAINTS ===\n")
	fmt.Println("Build tags: // +build ou //go:build")
	fmt.Println("Compilação condicional por OS/arquitetura")
	fmt.Println("Exemplo: // +build linux darwin")
}




