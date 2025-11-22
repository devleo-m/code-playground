package main

import (
	"fmt"
	"testing"
)

/*
 * ============================================================================
 * AULA 32: PERFORMANCE
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - pprof: Profiling
 * - Benchmarking
 * - Memory profiling
 * - CPU profiling
 */

func BenchmarkExemplo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i * 2
	}
}

func main() {
	fmt.Println("=== AULA 32: PERFORMANCE ===\n")
	fmt.Println("pprof: go tool pprof")
	fmt.Println("Benchmarking: go test -bench=.")
	fmt.Println("Memory profiling: -memprofile")
	fmt.Println("CPU profiling: -cpuprofile")
}

