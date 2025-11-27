package main

import (
	"fmt"
)

/*
 * ============================================================================
 * AULA 14: GENERICS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Introduzido no Go 1.18
 * - Type parameters: [T TipoConstraint]
 * - Constraints: any, comparable, interfaces customizadas
 * - Type inference: Go infere tipos automaticamente
 * - Permite código reutilizável com type safety
 */

// Função genérica com constraint any
func Print[T any](value T) {
	fmt.Printf("   Valor: %v (tipo: %T)\n", value, value)
}

// Função genérica com constraint comparable
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Função genérica para encontrar máximo
func Max[T interface {
	~int | ~float64 | ~string
}](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Tipo genérico: Stack
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	fmt.Println("=== AULA 14: GENERICS ===\n")
	
	// ===== FUNÇÕES GENÉRICAS =====
	fmt.Println("1. FUNÇÕES GENÉRICAS:")
	Print(42)
	Print("Olá")
	Print(3.14)
	
	// ===== COMPARABLE =====
	fmt.Println("\n2. COMPARABLE:")
	fmt.Printf("   Equal(10, 10): %t\n", Equal(10, 10))
	fmt.Printf("   Equal('a', 'b'): %t\n", Equal("a", "b"))
	fmt.Printf("   Equal(3.14, 3.14): %t\n", Equal(3.14, 3.14))
	
	// ===== MAX COM CONSTRAINT =====
	fmt.Println("\n3. MAX COM CONSTRAINT:")
	fmt.Printf("   Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("   Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("   Max('apple', 'banana'): %s\n", Max("apple", "banana"))
	
	// ===== TIPOS GENÉRICOS =====
	fmt.Println("\n4. TIPOS GENÉRICOS (Stack):")
	stack := Stack[int]{}
	
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Println("   Após push(1,2,3):")
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Printf("      Pop: %d\n", item)
	}
	
	// Stack de strings
	stackStr := Stack[string]{}
	stackStr.Push("primeiro")
	stackStr.Push("segundo")
	
	fmt.Println("\n   Stack de strings:")
	for !stackStr.IsEmpty() {
		item, _ := stackStr.Pop()
		fmt.Printf("      Pop: %s\n", item)
	}
	
	// ===== TYPE INFERENCE =====
	fmt.Println("\n5. TYPE INFERENCE:")
	// Go infere automaticamente
	resultado1 := Max(10, 20)        // T é int
	resultado2 := Max(3.14, 2.71)   // T é float64
	resultado3 := Max("a", "b")     // T é string
	
	fmt.Printf("   Max(10, 20): %d\n", resultado1)
	fmt.Printf("   Max(3.14, 2.71): %.2f\n", resultado2)
	fmt.Printf("   Max('a', 'b'): %s\n", resultado3)
	
	// Explícito (opcional)
	resultado4 := Max[int](10, 20)
	fmt.Printf("   Max[int](10, 20): %d\n", resultado4)
}



