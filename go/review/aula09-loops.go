package main

import "fmt"

/*
 * ============================================================================
 * AULA 09: LOOPS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - for clássico: for init; cond; post {}
 * - for while-style: for cond {}
 * - for infinito: for {}
 * - for range: Iterar sobre coleções (PREFERIDO)
 * - break: Sair do loop
 * - continue: Pular iteração
 * - Labels: Controlar loops aninhados
 */

func main() {
	fmt.Println("=== AULA 09: LOOPS ===\n")
	
	// ===== FOR CLÁSSICO =====
	fmt.Println("1. FOR CLÁSSICO:")
	for i := 0; i < 5; i++ {
		fmt.Printf("   Iteração %d\n", i)
	}
	
	// ===== FOR WHILE-STYLE =====
	fmt.Println("\n2. FOR WHILE-STYLE:")
	contador := 0
	for contador < 5 {
		fmt.Printf("   Contador: %d\n", contador)
		contador++
	}
	
	// ===== FOR INFINITO =====
	fmt.Println("\n3. FOR INFINITO (com break):")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("   Loop infinito: %d\n", i)
		i++
	}
	
	// ===== FOR RANGE COM SLICES =====
	fmt.Println("\n4. FOR RANGE COM SLICES:")
	numeros := []int{10, 20, 30, 40, 50}
	
	// Com índice e valor
	fmt.Println("   Com índice e valor:")
	for indice, valor := range numeros {
		fmt.Printf("      Índice %d: valor %d\n", indice, valor)
	}
	
	// Apenas valores
	fmt.Println("   Apenas valores:")
	for _, valor := range numeros {
		fmt.Printf("      Valor: %d\n", valor)
	}
	
	// Apenas índices
	fmt.Println("   Apenas índices:")
	for indice := range numeros {
		fmt.Printf("      Índice: %d\n", indice)
	}
	
	// ===== FOR RANGE COM MAPS =====
	fmt.Println("\n5. FOR RANGE COM MAPS:")
	cores := map[string]string{
		"vermelho": "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
	}
	
	for chave, valor := range cores {
		fmt.Printf("   Cor: %s = %s\n", chave, valor)
	}
	
	// ===== FOR RANGE COM STRINGS (RETORNA RUNES!) =====
	fmt.Println("\n6. FOR RANGE COM STRINGS (retorna RUNES):")
	texto := "Olá, 世界! 🚀"
	for indice, rune := range texto {
		fmt.Printf("   Posição %d: %c (Unicode: %d)\n", indice, rune, rune)
		if indice >= 10 { // Limitar output
			break
		}
	}
	
	// ===== BREAK =====
	fmt.Println("\n7. BREAK:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("   Encontrei 5! Saindo...")
			break
		}
		fmt.Printf("   Valor: %d\n", i)
	}
	
	// ===== CONTINUE =====
	fmt.Println("\n8. CONTINUE:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Pula números pares
		}
		fmt.Printf("   Número ímpar: %d\n", i)
	}
	
	// ===== LABELS =====
	fmt.Println("\n9. LABELS (loops aninhados):")
LoopExterno:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break LoopExterno // Sai do loop externo!
			}
			fmt.Printf("   i=%d, j=%d\n", i, j)
		}
	}
	
	// ===== LOOP ANINHADO =====
	fmt.Println("\n10. LOOP ANINHADO (tabela de multiplicação):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("   %d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
	
	// ===== MODIFICAR SLICE DURANTE ITERAÇÃO =====
	fmt.Println("\n11. MODIFICAR SLICE DURANTE ITERAÇÃO:")
	numeros2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Antes: %v\n", numeros2)
	
	// Modificar elementos é seguro
	for i := range numeros2 {
		numeros2[i] *= 2
	}
	fmt.Printf("   Depois (elementos modificados): %v\n", numeros2)
}



