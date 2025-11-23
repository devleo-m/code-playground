package main

import "fmt"

/*
 * ============================================================================
 * AULA 06: MAPS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Estrutura chave-valor (dicionário)
 * - Chaves únicas
 * - Busca O(1) em média
 * - Desordenado (ordem aleatória)
 * - Reference type
 * - Chaves devem ser comparáveis
 */

func main() {
	fmt.Println("=== AULA 06: MAPS ===\n")
	
	// ===== DECLARAÇÃO =====
	fmt.Println("1. DECLARAÇÃO:")
	// Nil map (não pode escrever!)
	var nilMap map[string]int
	fmt.Printf("   nilMap == nil: %t\n", nilMap == nil)
	// nilMap["chave"] = 10  // PANIC! Não pode escrever em nil map
	
	// Empty map (pronto para usar)
	emptyMap := make(map[string]int)
	fmt.Printf("   emptyMap == nil: %t\n", emptyMap == nil)
	emptyMap["chave"] = 10  // OK!
	
	// Literal
	cores := map[string]string{
		"vermelho": "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
	}
	fmt.Printf("   cores: %v\n", cores)
	
	// ===== OPERAÇÕES BÁSICAS =====
	fmt.Println("\n2. OPERAÇÕES BÁSICAS:")
	idades := make(map[string]int)
	
	// Adicionar
	idades["Alice"] = 25
	idades["Bob"] = 30
	idades["Carlos"] = 28
	fmt.Printf("   Após adicionar: %v\n", idades)
	
	// Buscar
	idadeAlice := idades["Alice"]
	fmt.Printf("   Idade de Alice: %d\n", idadeAlice)
	
	// Buscar inexistente (retorna zero)
	idadeDesconhecido := idades["Desconhecido"]
	fmt.Printf("   Idade de Desconhecido: %d (zero value)\n", idadeDesconhecido)
	
	// Verificar existência (padrão idiomático)
	idade, existe := idades["Bob"]
	if existe {
		fmt.Printf("   Bob tem %d anos\n", idade)
	} else {
		fmt.Println("   Bob não encontrado")
	}
	
	idade2, existe2 := idades["Desconhecido"]
	if existe2 {
		fmt.Printf("   Desconhecido tem %d anos\n", idade2)
	} else {
		fmt.Println("   Desconhecido não encontrado")
	}
	
	// Atualizar
	idades["Alice"] = 26
	fmt.Printf("   Após atualizar Alice: %v\n", idades)
	
	// ===== DELETE =====
	fmt.Println("\n3. DELETE:")
	frutas := map[string]int{
		"maçã":    10,
		"banana":  5,
		"laranja": 8,
		"uva":     12,
	}
	fmt.Printf("   Antes: %v\n", frutas)
	
	delete(frutas, "banana")
	fmt.Printf("   Após deletar 'banana': %v\n", frutas)
	
	delete(frutas, "melancia") // Não causa erro se não existir
	fmt.Printf("   Após deletar 'melancia' (não existe): %v\n", frutas)
	
	// ===== ITERAÇÃO =====
	fmt.Println("\n4. ITERAÇÃO (ordem aleatória!):")
	notas := map[string]float64{
		"Matemática": 8.5,
		"Português":  7.0,
		"História":   9.0,
		"Física":     6.5,
	}
	
	fmt.Println("   Chave e valor:")
	for disciplina, nota := range notas {
		fmt.Printf("      %s: %.1f\n", disciplina, nota)
	}
	
	fmt.Println("   Apenas chaves:")
	for disciplina := range notas {
		fmt.Printf("      %s\n", disciplina)
	}
	
	// ===== TIPOS DE CHAVES =====
	fmt.Println("\n5. TIPOS DE CHAVES PERMITIDAS:")
	// String (mais comum)
	mapa1 := map[string]int{"a": 1, "b": 2}
	fmt.Printf("   String keys: %v\n", mapa1)
	
	// Int
	mapa2 := map[int]string{1: "um", 2: "dois"}
	fmt.Printf("   Int keys: %v\n", mapa2)
	
	// Struct (se todos campos comparáveis)
	type Coordenada struct {
		X, Y int
	}
	mapa3 := map[Coordenada]string{
		{0, 0}: "origem",
		{1, 1}: "diagonal",
	}
	fmt.Printf("   Struct keys: %v\n", mapa3)
	
	// ===== MAPS SÃO REFERENCE TYPES =====
	fmt.Println("\n6. MAPS SÃO REFERENCE TYPES:")
	estoque := map[string]int{
		"maçã":   10,
		"banana": 5,
	}
	fmt.Printf("   Antes: %v\n", estoque)
	
	adicionarItem(estoque, "laranja", 8)
	fmt.Printf("   Depois: %v (foi modificado!)\n", estoque)
	
	// ===== PADRÕES COMUNS =====
	fmt.Println("\n7. PADRÕES COMUNS:")
	
	// Contador de frequência
	texto := "go é legal go é rápido go é moderno"
	palavras := []string{"go", "é", "legal", "go", "é", "rápido", "go", "é", "moderno"}
	contagem := make(map[string]int)
	
	for _, palavra := range palavras {
		contagem[palavra]++ // Incrementa (0 se não existir)
	}
	
	fmt.Println("   Contador de frequência:")
	for palavra, count := range contagem {
		fmt.Printf("      %s: %d vezes\n", palavra, count)
	}
	
	// Simular Set
	set := make(map[string]bool)
	set["maçã"] = true
	set["banana"] = true
	set["maçã"] = true // Duplicata ignorada
	
	fmt.Printf("   Set (map[string]bool): %v\n", set)
	fmt.Printf("   Contém 'maçã'? %t\n", set["maçã"])
	fmt.Printf("   Contém 'uva'? %t\n", set["uva"])
}

func adicionarItem(m map[string]int, chave string, valor int) {
	m[chave] = valor // Modifica o map original
}


