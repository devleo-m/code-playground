package main

import "fmt"

/*
 * ============================================================================
 * AULA 35: ESCAPE ANALYSIS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - go build -gcflags="-m"
 * - Stack vs Heap allocation
 * - Otimizações do compilador
 */

func main() {
	fmt.Println("=== AULA 35: ESCAPE ANALYSIS ===\n")
	fmt.Println("Ver escape analysis: go build -gcflags='-m'")
	fmt.Println("Compilador decide onde alocar")
	fmt.Println("Retornar pointer força escape para heap")
}


