package main

import "fmt"

/*
 * ============================================================================
 * AULA 34: MEMORY MANAGEMENT
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Stack vs Heap
 * - GC tuning
 * - Memory profiling
 * - Escape analysis
 */

func main() {
	fmt.Println("=== AULA 34: MEMORY MANAGEMENT ===\n")
	fmt.Println("Stack: Variáveis locais, rápido")
	fmt.Println("Heap: Dados que vivem além da função")
	fmt.Println("GC: Garbage Collector automático")
	fmt.Println("Tuning: GODEBUG=gctrace=1")
}



