package main

import "fmt"

/*
 * ============================================================================
 * AULA 12: METHODS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Methods são funções com receiver
 * - Value receiver: Recebe cópia (não modifica)
 * - Pointer receiver: Recebe ponteiro (pode modificar)
 * - Go converte automaticamente entre valor e ponteiro
 */

type Retangulo struct {
	Largura float64
	Altura  float64
}

// Value receiver (não modifica)
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Value receiver (retorna novo valor)
func (r Retangulo) Escalar(fator float64) Retangulo {
	return Retangulo{
		Largura: r.Largura * fator,
		Altura:  r.Altura * fator,
	}
}

type Contador struct {
	valor int
}

// Pointer receiver (modifica)
func (c *Contador) Incrementar() {
	c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}

func main() {
	fmt.Println("=== AULA 12: METHODS ===\n")
	
	// ===== VALUE RECEIVER =====
	fmt.Println("1. VALUE RECEIVER:")
	ret := Retangulo{Largura: 10, Altura: 5}
	fmt.Printf("   ret.Area() = %.2f\n", ret.Area())
	
	ret2 := ret.Escalar(2)
	fmt.Printf("   ret original: %+v\n", ret)
	fmt.Printf("   ret2 escalado: %+v\n", ret2)
	
	// ===== POINTER RECEIVER =====
	fmt.Println("\n2. POINTER RECEIVER:")
	contador := Contador{valor: 0}
	fmt.Printf("   Valor antes: %d\n", contador.Valor())
	
	contador.Incrementar()
	fmt.Printf("   Valor depois: %d\n", contador.Valor())
	
	// ===== CONVERSÃO AUTOMÁTICA =====
	fmt.Println("\n3. CONVERSÃO AUTOMÁTICA:")
	pessoa := Pessoa{Nome: "João", Idade: 30}
	pessoaPtr := &Pessoa{Nome: "Maria", Idade: 25}
	
	// Pode chamar value receiver em pointer
	pessoa.Apresentar()
	pessoaPtr.Apresentar() // Go faz (*pessoaPtr).Apresentar()
	
	// Pode chamar pointer receiver em valor
	pessoa.FazerAniversario() // Go faz (&pessoa).FazerAniversario()
	fmt.Printf("   pessoa.Idade: %d\n", pessoa.Idade)
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() {
	fmt.Printf("   Olá, eu sou %s\n", p.Nome)
}

func (p *Pessoa) FazerAniversario() {
	p.Idade++
}


