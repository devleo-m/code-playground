package main

import "fmt"

// ============================================
// 1. POINTERS BÁSICOS
// ============================================

func exemploPointerBasico() {
	fmt.Println("=== Exemplo 1: Pointer Básico ===")
	
	var x int = 42
	var p *int = &x  // p armazena o endereço de x
	
	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Endereço de x: %p\n", &x)
	fmt.Printf("Valor de p (endereço): %p\n", p)
	fmt.Printf("Valor apontado por p: %d\n", *p)
	
	// Modificando através do pointer
	*p = 100
	fmt.Printf("Novo valor de x: %d\n", x)
}

func exemploPointerZeroValue() {
	fmt.Println("\n=== Exemplo 2: Zero Value de Pointer ===")
	
	var p *int
	fmt.Printf("Pointer nil: %p\n", p)
	
	if p == nil {
		fmt.Println("Pointer é nil!")
	}
	
	// Tentar usar *p aqui causaria panic!
	// *p = 10  // ERRO: panic: runtime error: invalid memory address
}

func exemploPassagemPorValor() {
	fmt.Println("\n=== Exemplo 3: Passagem por Valor ===")
	
	x := 10
	fmt.Printf("Antes da função: x = %d\n", x)
	naoModifica(x)
	fmt.Printf("Depois da função: x = %d\n", x)
}

func naoModifica(valor int) {
	valor = 20  // Modifica apenas a cópia local
	fmt.Printf("Dentro da função: valor = %d\n", valor)
}

func exemploPassagemPorReferencia() {
	fmt.Println("\n=== Exemplo 4: Passagem por Referência ===")
	
	x := 10
	fmt.Printf("Antes da função: x = %d\n", x)
	modifica(&x)  // Passa o endereço
	fmt.Printf("Depois da função: x = %d\n", x)
}

func modifica(ptr *int) {
	*ptr = 20  // Modifica o valor original
	fmt.Printf("Dentro da função: *ptr = %d\n", *ptr)
}

// ============================================
// 2. POINTERS COM STRUCTS
// ============================================

type Pessoa struct {
	Nome  string
	Idade int
}

func exemploPointerComStruct() {
	fmt.Println("\n=== Exemplo 5: Pointer com Struct ===")
	
	p := Pessoa{Nome: "João", Idade: 30}
	fmt.Printf("Antes: %+v\n", p)
	
	// Passando struct por valor (cópia)
	modificaStructPorValor(p)
	fmt.Printf("Depois (por valor): %+v\n", p)
	
	// Passando struct por referência (pointer)
	modificaStructPorReferencia(&p)
	fmt.Printf("Depois (por referência): %+v\n", p)
}

func modificaStructPorValor(p Pessoa) {
	p.Idade = 99  // Modifica apenas a cópia
}

func modificaStructPorReferencia(p *Pessoa) {
	p.Idade = 99  // Modifica o original
	// Go permite usar p.Idade ao invés de (*p).Idade
}

func exemploPointerStructShorthand() {
	fmt.Println("\n=== Exemplo 6: Shorthand de Pointer Struct ===")
	
	p := &Pessoa{Nome: "Maria", Idade: 25}
	
	// Ambas as formas são equivalentes:
	fmt.Printf("(*p).Nome: %s\n", (*p).Nome)
	fmt.Printf("p.Nome: %s\n", p.Nome)  // Go faz o dereference automaticamente
	
	// Modificando
	p.Idade = 26
	fmt.Printf("Nova idade: %d\n", p.Idade)
}

func exemploMethodReceiver() {
	fmt.Println("\n=== Exemplo 7: Method Receiver com Pointer ===")
	
	p := Pessoa{Nome: "Pedro", Idade: 20}
	
	// Método com receiver por valor (não modifica)
	p.AniversarioPorValor()
	fmt.Printf("Após aniversário (por valor): %d\n", p.Idade)
	
	// Método com receiver por pointer (modifica)
	p.AniversarioPorReferencia()
	fmt.Printf("Após aniversário (por referência): %d\n", p.Idade)
}

// Método com receiver por valor
func (p Pessoa) AniversarioPorValor() {
	p.Idade++  // Não modifica o original
}

// Método com receiver por pointer
func (p *Pessoa) AniversarioPorReferencia() {
	p.Idade++  // Modifica o original
}

// ============================================
// 3. POINTERS COM SLICES E MAPS
// ============================================

func exemploSliceComPointer() {
	fmt.Println("\n=== Exemplo 8: Slice (Reference Type) ===")
	
	slice := []int{1, 2, 3}
	fmt.Printf("Antes: %v\n", slice)
	
	modificaSlice(slice)
	fmt.Printf("Depois: %v\n", slice)  // Modificado!
	
	// Mas reatribuir não afeta o original
	slice = []int{10, 20, 30}
	fmt.Printf("Após reatribuição: %v\n", slice)
}

func modificaSlice(s []int) {
	s[0] = 999  // Modifica o slice original
	fmt.Printf("Dentro da função: %v\n", s)
}

func exemploMapComPointer() {
	fmt.Println("\n=== Exemplo 9: Map (Reference Type) ===")
	
	m := map[string]int{"a": 1, "b": 2}
	fmt.Printf("Antes: %v\n", m)
	
	modificaMap(m)
	fmt.Printf("Depois: %v\n", m)  // Modificado!
}

func modificaMap(m map[string]int) {
	m["c"] = 3  // Modifica o map original
	m["a"] = 999
}

func exemploSlicePointerExplicito() {
	fmt.Println("\n=== Exemplo 10: Pointer Explícito para Slice ===")
	
	slice := &[]int{1, 2, 3}
	fmt.Printf("Antes: %v\n", *slice)
	
	reatribuiSlice(slice)
	fmt.Printf("Depois: %v\n", *slice)  // Agora reatribuiu!
}

func reatribuiSlice(s *[]int) {
	*s = []int{10, 20, 30}  // Reatribui o slice original
}

// ============================================
// 4. MEMORY MANAGEMENT
// ============================================

func exemploStackVsHeap() {
	fmt.Println("\n=== Exemplo 11: Stack vs Heap (Conceitual) ===")
	
	// Variáveis locais pequenas geralmente vão para a stack
	x := 10
	y := 20
	resultado := x + y
	fmt.Printf("Resultado (stack): %d\n", resultado)
	
	// Quando você retorna um pointer, Go pode mover para heap
	ptr := criaIntNoHeap()
	fmt.Printf("Valor do heap: %d\n", *ptr)
}

func criaIntNoHeap() *int {
	valor := 42  // Escape analysis pode mover para heap
	return &valor  // Retornar pointer força escape para heap
}

// ============================================
// MAIN
// ============================================

func main() {
	exemploPointerBasico()
	exemploPointerZeroValue()
	exemploPassagemPorValor()
	exemploPassagemPorReferencia()
	exemploPointerComStruct()
	exemploPointerStructShorthand()
	exemploMethodReceiver()
	exemploSliceComPointer()
	exemploMapComPointer()
	exemploSlicePointerExplicito()
	exemploStackVsHeap()
}

