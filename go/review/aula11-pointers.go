package main

import "fmt"

/*
 * ============================================================================
 * AULA 11: POINTERS E MEMORY MANAGEMENT
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Pointers armazenam endereços de memória
 * - &variavel: Obtém endereço
 * - *ponteiro: Acessa valor apontado
 * - Zero value: nil
 * - Slices/maps já são reference types
 * - Go gerencia memória automaticamente (GC)
 */

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	fmt.Println("=== AULA 11: POINTERS E MEMORY ===\n")
	
	// ===== POINTERS BÁSICOS =====
	fmt.Println("1. POINTERS BÁSICOS:")
	var x int = 42
	var p *int = &x
	
	fmt.Printf("   Valor de x: %d\n", x)
	fmt.Printf("   Endereço de x: %p\n", &x)
	fmt.Printf("   Valor de p (endereço): %p\n", p)
	fmt.Printf("   Valor apontado por p: %d\n", *p)
	
	// Modificar através do pointer
	*p = 99
	fmt.Printf("   x após *p = 99: %d\n", x)
	
	// ===== ZERO VALUE =====
	fmt.Println("\n2. ZERO VALUE:")
	var p2 *int
	fmt.Printf("   p2 == nil: %t\n", p2 == nil)
	fmt.Printf("   p2: %p\n", p2)
	// *p2 = 10  // PANIC! Não pode usar nil pointer
	
	// ===== PASSAGEM POR VALOR VS REFERÊNCIA =====
	fmt.Println("\n3. PASSAGEM POR VALOR VS REFERÊNCIA:")
	num := 10
	fmt.Printf("   num antes: %d\n", num)
	
	naoModifica(num)
	fmt.Printf("   num depois (por valor): %d\n", num)
	
	modifica(&num)
	fmt.Printf("   num depois (por referência): %d\n", num)
	
	// ===== POINTERS COM STRUCTS =====
	fmt.Println("\n4. POINTERS COM STRUCTS:")
	pessoa := Pessoa{Nome: "João", Idade: 30}
	fmt.Printf("   pessoa antes: %+v\n", pessoa)
	
	modificaStruct(&pessoa)
	fmt.Printf("   pessoa depois: %+v\n", pessoa)
	
	// Acesso direto (Go faz conversão automática)
	pessoaPtr := &Pessoa{Nome: "Maria", Idade: 25}
	fmt.Printf("   pessoaPtr.Nome: %s (acesso direto)\n", pessoaPtr.Nome)
	pessoaPtr.Idade = 26
	fmt.Printf("   pessoaPtr.Idade: %d\n", pessoaPtr.Idade)
	
	// ===== SLICES SÃO REFERENCE TYPES =====
	fmt.Println("\n5. SLICES SÃO REFERENCE TYPES:")
	slice := []int{1, 2, 3}
	fmt.Printf("   slice antes: %v\n", slice)
	
	modificaSlice(slice)
	fmt.Printf("   slice depois: %v (elementos modificados!)\n", slice)
	
	// Mas reatribuir não afeta original
	reatribuiSlice(slice)
	fmt.Printf("   slice após reatribuir: %v (não mudou!)\n", slice)
	
	// Para reatribuir, precisa de pointer
	slicePtr := &[]int{1, 2, 3}
	reatribuiSlicePtr(slicePtr)
	fmt.Printf("   slicePtr após reatribuir: %v (mudou!)\n", *slicePtr)
	
	// ===== MAPS SÃO REFERENCE TYPES =====
	fmt.Println("\n6. MAPS SÃO REFERENCE TYPES:")
	m := map[string]int{"a": 1, "b": 2}
	fmt.Printf("   m antes: %v\n", m)
	
	modificaMap(m)
	fmt.Printf("   m depois: %v (foi modificado!)\n", m)
	
	// ===== NEW() =====
	fmt.Println("\n7. NEW():")
	p3 := new(int)  // Aloca memória e retorna pointer
	*p3 = 42
	fmt.Printf("   *p3: %d\n", *p3)
	
	pessoa2 := new(Pessoa)
	pessoa2.Nome = "Pedro"
	pessoa2.Idade = 28
	fmt.Printf("   pessoa2: %+v\n", pessoa2)
	
	// ===== MEMORY MANAGEMENT =====
	fmt.Println("\n8. MEMORY MANAGEMENT:")
	fmt.Println("   - Stack: Variáveis locais, rápido")
	fmt.Println("   - Heap: Dados que vivem além da função")
	fmt.Println("   - Escape Analysis: Compilador decide onde alocar")
	fmt.Println("   - Garbage Collector: Automático, tri-color mark-and-sweep")
	fmt.Println("   - Você NÃO precisa gerenciar memória manualmente!")
}

func naoModifica(valor int) {
	valor = 20 // Modifica apenas a cópia
}

func modifica(ptr *int) {
	*ptr = 20 // Modifica o original
}

func modificaStruct(p *Pessoa) {
	p.Idade = 99 // Modifica o original
}

func modificaSlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Modifica o array subjacente
	}
}

func reatribuiSlice(s []int) {
	s = []int{10, 20, 30} // Reatribui apenas a cópia local
}

func reatribuiSlicePtr(s *[]int) {
	*s = []int{10, 20, 30} // Reatribui o slice original
}

func modificaMap(m map[string]int) {
	m["c"] = 3   // Adiciona novo elemento
	m["a"] = 999 // Modifica elemento existente
}





