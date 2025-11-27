package main

import "fmt"

/*
 * ============================================================================
 * AULA 05: SLICES
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Tamanho DINÂMICO (pode crescer)
 * - Reference type (compartilha array subjacente)
 * - Componentes: ponteiro, length, capacity
 * - append(): Adiciona elementos (sempre atribua resultado!)
 * - copy(): Copia elementos
 * - Slicing: slice[inicio:fim] cria sub-slices
 * - Sub-slices compartilham memória!
 */

func main() {
	fmt.Println("=== AULA 05: SLICES ===\n")
	
	// ===== DECLARAÇÃO =====
	fmt.Println("1. DECLARAÇÃO:")
	// Com make
	numeros1 := make([]int, 5)       // length=5, cap=5
	numeros2 := make([]int, 5, 10)   // length=5, cap=10
	
	fmt.Printf("   make([]int, 5): len=%d, cap=%d, %v\n", 
		len(numeros1), cap(numeros1), numeros1)
	fmt.Printf("   make([]int, 5, 10): len=%d, cap=%d, %v\n", 
		len(numeros2), cap(numeros2), numeros2)
	
	// Literal
	frutas := []string{"Maçã", "Banana", "Laranja"}
	fmt.Printf("   Literal: %v\n", frutas)
	
	// Nil vs vazio
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("   nil slice: len=%d, nil? %t\n", len(nilSlice), nilSlice == nil)
	fmt.Printf("   empty slice: len=%d, nil? %t\n", len(emptySlice), emptySlice == nil)
	
	// ===== LENGTH VS CAPACITY =====
	fmt.Println("\n2. LENGTH VS CAPACITY:")
	slice := make([]int, 3, 5)
	fmt.Printf("   len=%d (elementos existentes)\n", len(slice))
	fmt.Printf("   cap=%d (capacidade máxima)\n", cap(slice))
	fmt.Printf("   Array subjacente: [%d %d %d _ _]\n", slice[0], slice[1], slice[2])
	
	// ===== APPEND =====
	fmt.Println("\n3. APPEND (sempre atribua resultado!):")
	numeros := []int{1, 2, 3}
	fmt.Printf("   Antes: %v\n", numeros)
	
	numeros = append(numeros, 4)
	fmt.Printf("   Após append(4): %v\n", numeros)
	
	numeros = append(numeros, 5, 6, 7)
	fmt.Printf("   Após append(5,6,7): %v\n", numeros)
	
	// Append de outro slice
	outros := []int{8, 9, 10}
	numeros = append(numeros, outros...)
	fmt.Printf("   Após append(outros...): %v\n", numeros)
	
	// ===== SLICING =====
	fmt.Println("\n4. SLICING (cria sub-slices):")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	parte1 := original[2:5]    // [2 3 4]
	parte2 := original[:4]     // [0 1 2 3]
	parte3 := original[6:]     // [6 7 8 9]
	parte4 := original[:]      // Cópia completa
	
	fmt.Printf("   original[2:5]: %v\n", parte1)
	fmt.Printf("   original[:4]: %v\n", parte2)
	fmt.Printf("   original[6:]: %v\n", parte3)
	fmt.Printf("   original[:]: %v\n", parte4)
	
	// ===== COMPARTILHAMENTO DE MEMÓRIA =====
	fmt.Println("\n5. COMPARTILHAMENTO DE MEMÓRIA (CUIDADO!):")
	original2 := []int{1, 2, 3, 4, 5}
	parte := original2[1:4]  // [2 3 4]
	
	fmt.Printf("   Antes: original=%v, parte=%v\n", original2, parte)
	
	parte[0] = 999  // Modifica o array subjacente!
	
	fmt.Printf("   Depois: original=%v (FOI MODIFICADO!)\n", original2)
	fmt.Printf("           parte=%v\n", parte)
	
	// ===== COPY =====
	fmt.Println("\n6. COPY (cria cópia independente):")
	orig := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(orig))
	
	copy(copia, orig)
	copia[0] = 999
	
	fmt.Printf("   original: %v (não mudou)\n", orig)
	fmt.Printf("   cópia: %v\n", copia)
	
	// ===== REMOVENDO ELEMENTOS =====
	fmt.Println("\n7. REMOVENDO ELEMENTOS:")
	nums := []int{1, 2, 3, 4, 5}
	
	// Remover primeiro
	nums = nums[1:]
	fmt.Printf("   Remover primeiro: %v\n", nums)
	
	// Remover último
	nums = nums[:len(nums)-1]
	fmt.Printf("   Remover último: %v\n", nums)
	
	// Remover do meio (índice 1)
	nums = append(nums[:1], nums[2:]...)
	fmt.Printf("   Remover índice 1: %v\n", nums)
	
	// ===== PRÉ-ALOCAÇÃO (PERFORMANCE) =====
	fmt.Println("\n8. PRÉ-ALOCAÇÃO (IMPORTANTE PARA PERFORMANCE):")
	// RUIM: Múltiplas realocações
	ruim := []int{}
	fmt.Printf("   Sem pré-alocação: len=%d, cap=%d\n", len(ruim), cap(ruim))
	
	// BOM: Uma única alocação
	bom := make([]int, 0, 1000)
	fmt.Printf("   Com pré-alocação: len=%d, cap=%d\n", len(bom), cap(bom))
	fmt.Println("   Pré-alocar evita múltiplas realocações (muito mais rápido!)")
}



