package main

import "fmt"

/*
 * ============================================================================
 * AULA 04: ARRAYS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Tamanho FIXO (definido na declaração)
 * - Tipo inclui tamanho: [5]int ≠ [10]int
 * - Value type (copia quando atribuído)
 * - Indexação começa em 0
 * - len(array) retorna tamanho
 * - Raramente usado (prefira slices)
 */

func main() {
	fmt.Println("=== AULA 04: ARRAYS ===\n")
	
	// ===== DECLARAÇÃO BÁSICA =====
	fmt.Println("1. DECLARAÇÃO BÁSICA:")
	var numeros [5]int
	var nomes [3]string
	
	fmt.Printf("   numeros (valores zero): %v\n", numeros)
	fmt.Printf("   nomes (valores zero): %v\n", nomes)
	
	// ===== INICIALIZAÇÃO =====
	fmt.Println("\n2. INICIALIZAÇÃO:")
	idades := [5]int{18, 25, 30, 22, 45}
	notas := [4]float64{7.5, 8.0, 9.2, 6.8}
	dias := [...]string{"Segunda", "Terça", "Quarta"} // Tamanho automático
	
	fmt.Printf("   idades: %v\n", idades)
	fmt.Printf("   notas: %v\n", notas)
	fmt.Printf("   dias: %v\n", dias)
	
	// Inicialização parcial
	numeros2 := [5]int{1, 2} // Resto é zero
	fmt.Printf("   numeros2 (parcial): %v\n", numeros2)
	
	// Inicialização com índices específicos
	valores := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Printf("   valores (índices específicos): %v\n", valores)
	
	// ===== ACESSO E MODIFICAÇÃO =====
	fmt.Println("\n3. ACESSO E MODIFICAÇÃO:")
	frutas := [4]string{"Maçã", "Banana", "Laranja", "Uva"}
	
	fmt.Printf("   frutas[0]: %s\n", frutas[0])
	fmt.Printf("   frutas[3]: %s\n", frutas[3])
	
	frutas[1] = "Morango"
	fmt.Printf("   Após modificar: %v\n", frutas)
	fmt.Printf("   Tamanho: %d\n", len(frutas))
	
	// ===== ITERAÇÃO =====
	fmt.Println("\n4. ITERAÇÃO:")
	numeros3 := [5]int{10, 20, 30, 40, 50}
	
	// Loop tradicional
	fmt.Println("   Loop tradicional:")
	for i := 0; i < len(numeros3); i++ {
		fmt.Printf("      Posição %d: %d\n", i, numeros3[i])
	}
	
	// Range (preferido)
	fmt.Println("   Range (preferido):")
	for indice, valor := range numeros3 {
		fmt.Printf("      Posição %d: %d\n", indice, valor)
	}
	
	// Apenas valores
	fmt.Println("   Apenas valores:")
	for _, valor := range numeros3 {
		fmt.Printf("      Valor: %d\n", valor)
	}
	
	// ===== ARRAYS SÃO VALUE TYPES =====
	fmt.Println("\n5. ARRAYS SÃO VALUE TYPES (CÓPIA):")
	original := [3]int{1, 2, 3}
	fmt.Printf("   Original antes: %v\n", original)
	
	modificarArray(original)
	fmt.Printf("   Original depois: %v (NÃO mudou!)\n", original)
	
	// ===== ARRAYS MULTIDIMENSIONAIS =====
	fmt.Println("\n6. ARRAYS MULTIDIMENSIONAIS:")
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	fmt.Printf("   matriz[1][2]: %d\n", matriz[1][2])
	fmt.Println("   Matriz completa:")
	for i := 0; i < len(matriz); i++ {
		fmt.Printf("      ")
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
	
	// ===== LIMITAÇÕES =====
	fmt.Println("\n7. LIMITAÇÕES:")
	fmt.Println("   - Tamanho fixo (não pode crescer)")
	fmt.Println("   - [5]int ≠ [10]int (tipos diferentes)")
	fmt.Println("   - Cópia custosa para arrays grandes")
	fmt.Println("   - Raramente usado (prefira slices)")
}

// Função que recebe array por valor (cópia)
func modificarArray(arr [3]int) {
	arr[0] = 999
	fmt.Printf("   Dentro da função: %v\n", arr)
}





