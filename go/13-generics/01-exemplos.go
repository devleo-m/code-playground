package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// ============================================
// 1. POR QUE GENERICS? - O PROBLEMA
// ============================================

// ANTES: Sem generics, precisávamos criar funções separadas para cada tipo
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MaxString(a, b string) string {
	if a > b {
		return a
	}
	return b
}

// ============================================
// 2. GENERIC FUNCTIONS - SOLUÇÃO
// ============================================

// Função genérica com constraint Ordered (permite comparações com >, <, etc.)
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ============================================
// 3. TYPE PARAMETERS E CONSTRAINTS
// ============================================

// Constraint básico: any (qualquer tipo)
func PrintAny[T any](value T) {
	fmt.Println(value)
}

// Constraint: comparable (tipos que podem ser comparados com == e !=)
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Constraint customizada usando interface
type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Soma[T Numeric](a, b T) T {
	return a + b
}

// ============================================
// 4. GENERIC TYPES
// ============================================

// Container genérico
type Container[T any] struct {
	value T
}

func (c *Container[T]) Set(value T) {
	c.value = value
}

func (c *Container[T]) Get() T {
	return c.value
}

// Stack genérico
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
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

// ============================================
// 5. TYPE INFERENCE
// ============================================

// O compilador infere o tipo automaticamente
func ExampleTypeInference() {
	// Type inference: o compilador sabe que é int
	maxInt := Max(10, 20)
	fmt.Println(maxInt)

	// Type inference: o compilador sabe que é string
	maxString := Max("apple", "banana")
	fmt.Println(maxString)

	// Type inference: o compilador sabe que é float64
	maxFloat := Max(3.14, 2.71)
	fmt.Println(maxFloat)

	// Você também pode especificar explicitamente
	maxIntExplicit := Max[int](10, 20)
	fmt.Println(maxIntExplicit)
}

// ============================================
// 6. GENERIC INTERFACES
// ============================================

// Interface genérica
type Adder[T any] interface {
	Add(T) T
}

type Number int

func (n Number) Add(other Number) Number {
	return n + other
}

// Função que usa interface genérica
func ProcessAdder[T Adder[T]](a, b T) T {
	return a.Add(b)
}

// ============================================
// 7. EXEMPLOS PRÁTICOS
// ============================================

// Função para encontrar elemento em slice
func Find[T comparable](slice []T, value T) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

// Função para mapear slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Função para filtrar slice
func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// ============================================
// 8. EXEMPLOS DE USO
// ============================================

func main() {
	fmt.Println("=== Exemplos de Generics ===\n")

	// 1. Usando função genérica
	fmt.Println("1. Função Max genérica:")
	fmt.Printf("Max(10, 20) = %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71) = %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\") = %s\n\n", Max("apple", "banana"))

	// 2. Usando Container genérico
	fmt.Println("2. Container genérico:")
	container := Container[string]{}
	container.Set("Hello, Generics!")
	fmt.Printf("Container value: %s\n\n", container.Get())

	// 3. Usando Stack genérico
	fmt.Println("3. Stack genérico:")
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		fmt.Printf("Popped: %d\n", val)
	}
	fmt.Println()

	// 4. Usando funções utilitárias genéricas
	fmt.Println("4. Funções utilitárias:")
	numbers := []int{1, 2, 3, 4, 5}
	
	// Map: dobrar todos os números
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Dobrado: %v\n", doubled)
	
	// Filter: apenas números pares
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Pares: %v\n", evens)
	
	// Find: encontrar índice
	idx, found := Find(numbers, 3)
	if found {
		fmt.Printf("Encontrado 3 no índice: %d\n", idx)
	}
}

