package main

import (
	"fmt"
	"testing"
)

/*
 * ============================================================================
 * AULA 18: TESTING E BENCHMARKING
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - go test: Executa testes
 * - Arquivos _test.go
 * - Testes: func TestNome(t *testing.T)
 * - Benchmarks: func BenchmarkNome(b *testing.B)
 * - Tabela de testes
 * - Coverage: go test -cover
 */

// Função a ser testada
func Soma(a, b int) int {
	return a + b
}

// Exemplo de teste (execute: go test)
func TestSoma(t *testing.T) {
	resultado := Soma(2, 3)
	esperado := 5
	
	if resultado != esperado {
		t.Errorf("Soma(2, 3) = %d; esperado %d", resultado, esperado)
	}
}

// Exemplo de benchmark (execute: go test -bench=.)
func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(2, 3)
	}
}

func main() {
	fmt.Println("=== AULA 18: TESTING ===\n")
	fmt.Println("Execute os testes com: go test")
	fmt.Println("Execute benchmarks com: go test -bench=.")
	fmt.Println("Veja coverage com: go test -cover")
	
	// Exemplo de uso
	resultado := Soma(5, 3)
	fmt.Printf("Soma(5, 3) = %d\n", resultado)
}





