package main

import "fmt"

/*
 * ============================================================================
 * AULA 39: COMPILER FLAGS
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - -gcflags: Flags do compilador
 * - -ldflags: Flags do linker
 * - Otimizações
 */

func main() {
	fmt.Println("=== AULA 39: COMPILER FLAGS ===\n")
	fmt.Println("-gcflags: go build -gcflags='-m' (escape analysis)")
	fmt.Println("-ldflags: go build -ldflags='-s -w' (otimizações)")
	fmt.Println("Otimizações: -N (desabilita), -l (inline)")
}
